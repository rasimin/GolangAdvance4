package main

import (
	// "advance2/handler"
	"advance2/repository/postgres_gorm"
	"advance2/repository/postgres_pgx"
	"net"
	"time"

	// "advance2/router"
	"advance2/service"
	"context"
	"log"

	//"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	grpcHandler "advance2/handler/grpc"
	pb "advance2/proto/user_service/v1"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	//r := gin.Default()

	// setup database connection
	//dsn := "postgresql://postgres:postgres@localhost:5432/postgres"
	// setup pgx connection
	//pgxPool, err := connectDB(dsn)
	//if err != nil {
	//	log.Fatalln(err)
	//}

	// setup gorm connectoin
	dsn := "postgresql://postgres:admin@localhost:5432/NewDB"
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		log.Fatalln(err)
	}
	// setup service

	// uncomment to use postgres gorm
	userRepo := postgres_gorm.NewUserRepository(gormDB)
	userService := service.NewUserService(userRepo)
	//userHandler := ginHandler.NewUserHandler(userService)
	userHandler := grpcHandler.NewUserHandler(userService)

	// Run the grpc server
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, userHandler)
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	go func() {
		log.Println("Running grpc server in port :50051")
		_ = grpcServer.Serve(lis)
	}()
	time.Sleep(1 * time.Second)

	// Run the grpc gateway
	conn, err := grpc.NewClient(
		"0.0.0.0:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}
	gwmux := runtime.NewServeMux()
	if err = pb.RegisterUserServiceHandler(context.Background(), gwmux, conn); err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	// dengan default http server
	/*
		gwServer := &http.Server{
			Addr:    ":8080",
			Handler: gwmux,
		}
		log.Println("Running grpc gateway server in port :8080")
		_ = gwServer.ListenAndServe()
	*/

	// dengan GIN
	gwServer := gin.Default()
	gwServer.Group("v1/*{grpc_gateway}").Any("", gin.WrapH(gwmux))
	log.Println("Running grpc gateway server in port :8080")
	_ = gwServer.Run()
}

func connectDB(dbURL string) (postgres_pgx.PgxPoolIface, error) {
	return pgxpool.New(context.Background(), dbURL)
}
