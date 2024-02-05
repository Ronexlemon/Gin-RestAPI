package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
router:= gin.Default()
router.GET("/books",func(c *gin.Context){
	c.JSON(http.StatusFound, gin.H{"data":"hello"})
})

router.Run(":8000")

}
