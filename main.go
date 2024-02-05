package main

import (
	"ginrestapi/database"
	"ginrestapi/models"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	loadEnv()

	loadDatabase()

	// router := gin.Default()
	// router.GET("/books", func(c *gin.Context) {
	// 	c.JSON(http.StatusFound, gin.H{"data": "hello"})
	// })

	// router.Run(":8000")

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
