package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasramosdev/cnotes/internal/users"
)

func Login(ctx *gin.Context) {
	var authInput users.AuthInput

	if err := ctx.ShouldBindJSON(&authInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := usersService.Login(&authInput)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"bearer_token": token,
	})

}
