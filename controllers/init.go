package controllers

import (
	"net/http"

	authServices "github.com/Dparty/auth-services"
	"github.com/Dparty/common/fault"
	"github.com/Dparty/common/server"
	"github.com/gin-gonic/gin"
)

var authService = authServices.GetAuthService()
var router *gin.Engine

func Init(addr ...string) {
	router = gin.Default()
	router.Use(authService.Auth())
	router.Use(server.CorsMiddleware())
	var authApi AuthApi
	server.MetricsMiddleware(router)
	router.GET("/me", authApi.GetMe)
	router.POST("/sessions", authApi.CreateSession)
	router.POST("/accounts", authApi.CreateAccount)
	router.GET("/version", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"version": "0.0.2",
		})
	})
	router.Run(addr...)
}

func GetAccount(ctx *gin.Context) *authServices.Account {
	accountInterface, ok := ctx.Get("account")
	if !ok {
		fault.GinHandler(ctx, fault.ErrUnauthorized)
		return nil
	}
	account, ok := accountInterface.(authServices.Account)
	if !ok {
		fault.GinHandler(ctx, fault.ErrUnauthorized)
		return nil
	}
	return &account
}
