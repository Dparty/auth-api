package controllers

import (
	"fmt"

	"github.com/Dparty/auth-api/models"
	"github.com/gin-gonic/gin"
)

func CreateSession(ctx *gin.Context) {
	// fmt.Println("good")
	var createSessionRequest models.CreateSessionRequest
	ctx.ShouldBindJSON(&createSessionRequest)
	fmt.Println(createSessionRequest)
	token, err := authService.CreateSession(createSessionRequest.Email, createSessionRequest.Password)
	fmt.Println(token, err)
	// if err != nil {
	// 	ctx.JSON(http.StatusUnauthorized, "")
	// 	return
	// }
	// ctx.JSON(http.StatusCreated, &models.Session{
	// 	Token: token,
	// })
}
