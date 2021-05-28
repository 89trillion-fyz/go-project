package router

import (
	"fmt"

	"go-project/demo3/global"
	"go-project/demo3/initialize"

	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
)

// 初始化总路由
func init() {
	err := ini.MapTo(&global.GB_GONFIG, "../app.ini")
	if err != nil {
		fmt.Println("Failed to parse config file: %s", err)
	}
	fmt.Println("init global.GB_GONFIG", global.GB_GONFIG)
	initialize.Redis()
}
func Routers() *gin.Engine {

	var Router = gin.Default()
	// 跨域
	//Router.Use(middleware.Cors()) // 如需跨域可以打开
	// 方便统一添加路由组前缀 多服务器上线使用
	PublicGroup := Router.Group("")
	PublicGroup.Use(func(c *gin.Context) {
		if global.GB_REDIS == nil {
			initialize.Redis()
		}
	})
	{
		InitCdkRouter(PublicGroup)  // 自动初始化相关
		InitUserRouter(PublicGroup) // 自动初始化相关
	}
	return Router
}
