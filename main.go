package main

import (
	"github.com/Dparty/auth-api/controllers"
	"github.com/Dparty/common/config"
)

func main() {
	controllers.Init(":" + config.GetString("server.port"))
}
