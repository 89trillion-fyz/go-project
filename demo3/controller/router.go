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
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	chTranslations "github.com/go-playground/validator/v10/translations/zh"
	"go.uber.org/zap"
	"gopkg.in/ini.v1"
)

// loca 通常取决于 http 请求头的 'Accept-Language'
func transInit(local string) (err error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		zhT := zh.New() //chinese
		enT := en.New() //english
		uni := ut.New(enT, zhT, enT)

		var o bool
		trans, o := uni.GetTranslator(local)
		if !o {
			return fmt.Errorf("uni.GetTranslator(%s) failed", local)
		}
		//register translate
		// 注册翻译器
		switch local {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		case "zh":
			err = chTranslations.RegisterDefaultTranslations(v, trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		}
		global.GB_TRANS = trans
		return
	}
	return
}

// 初始化总路由
func init() {
	global.GB_GONFIG = new(config.Config)
	initialize.Zap()
	if err := transInit("zh"); err != nil {
		global.GB_LOG.Error("transInit error", zap.Any("error", err))
		return
	}
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
