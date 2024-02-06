package middleware

import (
	"ginrestapi/helper"
	"ginrestapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
    return func(context *gin.Context) {
        err := helper.ValidateJWT(context)
        if err != nil {
            context.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
            context.Abort()
            return
        }
        context.Next()
    }
}


func AddEntry(context *gin.Context) {
    var input models.Book
    if err := context.ShouldBindJSON(&input); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := helper.CurrentUser(context)

    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    input.UserID = user.ID

    savedEntry, err := input.Save()

    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    context.JSON(http.StatusCreated, gin.H{"data": savedEntry})
}