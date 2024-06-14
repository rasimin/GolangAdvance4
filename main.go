package main

import (
	"advance2/entity"
	// "advance2/handler"
	// "advance2/repository/slice"
	// "advance2/router"
	// "advance2/service"
	"context"
	"fmt"
	"log"

	// "github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	// r := gin.Default()

	// // setup service
	// var mockUserDBInSlice []entity.User
	// userRepo := slice.NewUserRepository(mockUserDBInSlice)
	// userService := service.NewUserService(userRepo)
	// userHandler := handler.NewUserHandler(userService)

	// // Routes
	// router.SetupRouter(r, userHandler)

	// "postgres: //YourUserName: YourPassword@YourHostname: 5432/YourDatabaseName"
	dsn := "postgresql://postgres:admin@localhost:5432/NewDB"
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalln(err)
	}
	err = pool.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("successfully connect to db")

	var u entity.User

	ctx := context.Background()

	err = pool.QueryRow(ctx, "select id, name from users order by id desc limit 2").Scan(&u.ID, &u.Name)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("user retrieved", u)

	_, err = pool.Exec(ctx, "insert into users (name, email, password, created_at, updated_at) values "+
		"('test','test@test.com','pass', NOW(), NOW())")
	if err != nil {
		log.Fatalln(err)
	}

	//openrow hanya untuk 1 data
	err = pool.QueryRow(ctx, "select id, name from users order by id desc limit 1").Scan(&u.ID, &u.Name)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("user retrieved2", u)

	//untuk memunculkan semuad ata table
	rows, err := pool.Query(ctx, "select id,name from users order by id desc limit 3")
	var users []entity.User
	for rows.Next() {
		var u2 entity.User
		rows.Scan(&u2.ID, &u2.Name)
		if err != nil {
			log.Fatalln(err)
		}
		users = append(users, u2)
	}
	fmt.Println("all user retrieved", users)

	// Run the server
	// log.Println("Running server on port 8080")
	// r.Run(":8080")

}
