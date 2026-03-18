package main

import (
	"go-rest-api/config"
	"go-rest-api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {

	config.ConnectDB()

	r := gin.Default()

	r.GET("/users", controllers.GetUsers)
	r.POST("/users", controllers.CreateUsers)

	r.Run(":8080")
}