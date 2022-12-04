package main

import (
	"github.com/gin-gonic/gin"

	"go-crud/models"
	"go-crud/controllers"
)


func main() {
	router := gin.Default()

	models.ConnectDatabase()
	
	router.GET("/api/v1/users", controllers.FindUsers)
	router.POST("/api/v1/users", controllers.CreateUser)
	router.GET("/api/v1/users/adult", controllers.FindAdults)
	router.GET("/api/v1/users/:id", controllers.FindUser)
	router.DELETE("/api/v1/users/:id", controllers.DeleteUser)

	router.Run(":3000")
}