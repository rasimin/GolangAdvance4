package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// MockUserHandler adalah mock untuk handler.UserHandler
type MockUserHandler struct{}

func (m *MockUserHandler) CreateUser(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "user created"})
}

func (m *MockUserHandler) GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "user found"})
}

func (m *MockUserHandler) GetAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "all users"})
}

func (m *MockUserHandler) UpdateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "user updated"})
}

func (m *MockUserHandler) DeleteUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}
