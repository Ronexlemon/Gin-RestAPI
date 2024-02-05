package controller

import (
	
	"ginrestapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(context *gin.Context){
	var input models.AuthenticationInput
	if err:= context.ShouldBindJSON(&input); err!= nil{
		context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	user := models.User{
		Username: input.Username,
		Password: input.Password,
	}
	savedUser, err:= user.Save()

	if err!= nil{
		context.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}
	context.JSON(http.StatusCreated,gin.H{"user":savedUser})
	
}