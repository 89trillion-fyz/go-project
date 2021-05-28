package controller

import (
	"fmt"
	"os"
	"strings"
	"time"

	"go-project/demo3/config"
	"go-project/demo3/global"
	"go-project/demo3/initialize"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gopkg.in/ini.v1"
)

// 初始化总路由
func init() {
	global.GB_GONFIG = new(config.Config)
	initialize.Zap()
	err := ini.MapTo(&global.GB_GONFIG, "app.ini")
	if err != nil {
		path, _ := os.Getwd()
		global.GB_LOG.Info("current dir", zap.Any("dir", path))
		// 单元测试
		if strings.Contains(path, "test") {
			err := ini.MapTo(&global.GB_GONFIG, "../app.ini")
			if err != nil {
				global.GB_LOG.Error("Failed to parse config file:", zap.Any("error", err))
			}
		}
		global.GB_LOG.Error("Failed to parse config file:", zap.Any("error", err))
	}
	fmt.Println("init global.GB_GONFIG", global.GB_GONFIG)
	initialize.Redis()
	initialize.ConnectToMongoDB(global.GB_GONFIG.Mongo.ApplyURI, time.Duration(global.GB_GONFIG.Mongo.Timeout), global.GB_GONFIG.Mongo.PoolSize)

}
func Routers() *gin.Engine {

	var Router = gin.Default()
	// 跨域
	//Router.Use(middleware.Cors()) // 如需跨域可以打开
	// 方便统一添加路由组前缀
	PublicGroup := Router.Group("")
	{
		InitCdkRouter(PublicGroup)  // 自动初始化相关
		InitUserRouter(PublicGroup) // 自动初始化相关
	}
	return Router
}
