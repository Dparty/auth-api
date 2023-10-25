package controllers

import (
	"fmt"
	"net/http"

	authServices "github.com/Dparty/auth-services"
	"github.com/Dparty/common/fault"
	"github.com/Dparty/common/server"
	"github.com/Dparty/dao"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var authService authServices.AuthService
var router *gin.Engine

func Init(addr ...string) {
	var err error
	viper.SetConfigName(".env.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("databases fatal error config file: %w", err))
	}
	user := viper.GetString("database.user")
	password := viper.GetString("database.password")
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	database := viper.GetString("database.database")
	db, err := dao.NewConnection(user, password, host, port, database)
	if err != nil {
		panic(err)
	}
	authService = authServices.NewAuthService(db)
	router = gin.Default()
	router.Use(authService.Auth())
	router.Use(server.CorsMiddleware())
	var authApi AuthApi
	router.POST("/sessions", authApi.CreateSession)
	router.POST("/accounts", authApi.CreateAccount)
	router.GET("/version", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"version": "0.0.1",
		})
	})
	router.Run(addr...)
}

func getAccount(ctx *gin.Context) *authServices.Account {
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
