package main

import (
	"go-clean-architecture/internal/config"
	"go-clean-architecture/internal/handler"
	"go-clean-architecture/internal/repository"
	"go-clean-architecture/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDB()

	r := gin.Default()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	r.GET("/users", userHandler.GetAll)
	r.POST("/users", userHandler.Create)

	r.Run(":8080")
}
