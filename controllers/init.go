package controllers

import (
	"fmt"

	authservices "github.com/Dparty/auth-api/services"
	"github.com/Dparty/common/server"
	"github.com/Dparty/dao"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var authService authservices.AuthService
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
	authService = authservices.NewAuthService(db)
	router = gin.Default()
	// router.Use(authService.Auth())
	router.Use(server.CorsMiddleware())
	router.POST("/sessions", CreateSession)
	router.Run(addr...)
}
