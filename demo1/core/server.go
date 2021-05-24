package core

import (
	"demo1/global"
	"demo1/router"
)

func RunServer() {
	address := global.G_CONFIG.Server.HttpPort
	Router := router.Routers()
	Router.Run(":" + address)

}
