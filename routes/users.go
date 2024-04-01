package routes

import (
	"example/gingonic/models"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User

	if err := context.ShouldBindJSON(&user); err != nil {
        context.JSON(400, gin.H{"error": "Could not parse data", "details": err.Error()})
        return
    }

	if err := user.Save(); err != nil {
		context.JSON(500, gin.H{"message": "Could not create user. Try again later.", "details": err.Error()})
		return
	}
	context.JSON(201, gin.H{"message": "User Created", "user": user})
}