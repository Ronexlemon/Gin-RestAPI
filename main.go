package main

import (
	//"fmt"
	"ginrestapi/logger"
	"io"
	"log"
	"net/http"
	"os"

	// "ginrestapi/controller"
	// "ginrestapi/database"
	// "ginrestapi/middleware"
	// "ginrestapi/models"
	//"log"

	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
	// "github.com/joho/godotenv"
)

func main() {
	logRus()
	router := gin.Default()

	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint formatted information is %v %v %v %v", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	f, _ := os.Create("ginLogging.log")
	gin.ForceConsoleColor()
	gin.DefaultWriter = colorable.NewColorableStdout()

	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	router.Use(gin.LoggerWithFormatter(logger.FormatLogs))

	router.GET("/", getDefault)
	router.Run(":9090")

	//loadEnv()

	//loadDatabase()
	//serverApplication()

}

func getDefault(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{"Data": "hello"})
}

// LOGRUS
func logRus() {
	//logrus.Println("Hi I Am Logrus")
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableTimestamp: false,
		FullTimestamp: false,

	})

	f, _ := os.Create("logrus.log")
	multi := io.MultiWriter(f,os.Stdout)
	logrus.SetOutput(multi)
		logrus.Traceln("Trace")
		logrus.Debugln("Debug")
		logrus.Infoln("Info")
		logrus.Warnln("Warn")
		logrus.Errorln("Error")
	// 	logrus.Panicln("Panic")
	// 	logrus.Fatalln("Fatal")
}

// func loadDatabase() {
// 	database.Connect()
// 	database.Database.AutoMigrate(&models.User{})
// 	database.Database.AutoMigrate(&models.Book{})
// }

// func loadEnv() {
// 	err := godotenv.Load(".env.local")

// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// }

// func serverApplication(){
// 	router := gin.Default()

// 	// publicRoutes := router.Group("/auth")
// 	// publicRoutes.POST("/register",controller.Register)
// 	// publicRoutes.POST("/login",controller.Login)

// 	// protectedRoutes := router.Group("/api")
//     // protectedRoutes.Use(middleware.JWTAuthMiddleware())
//     // protectedRoutes.POST("/entry", controller.GetAllEntries)
//     // protectedRoutes.GET("/entry", controller.GetAllEntries)

// 	// router.Run(":8000")
// 	// fmt.Println("Surver running on port 8000")

// }
