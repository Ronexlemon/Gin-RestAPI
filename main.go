package main

import (
	"fmt"
	"ginrestapi/controller"
	"ginrestapi/database"
	"ginrestapi/models"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()

	loadDatabase()
	serverApplication()

	

}

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&models.User{})
	database.Database.AutoMigrate(&models.Book{})
}

func loadEnv() {
	err := godotenv.Load(".env.local")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func serverApplication(){
	router := gin.Default()

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register",controller.Register)
	publicRoutes.POST("/login",controller.Login)

	router.Run(":8000")
	fmt.Println("Surver running on port 8000")

}
