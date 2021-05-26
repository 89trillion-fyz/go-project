package router

import (
	"go-project/demo3/api"
	"go-project/demo3/global"
	"go-project/demo3/initialize"

	"github.com/gin-gonic/gin"
)

func InitCdkRouter(Router *gin.RouterGroup) {
	cdkeyRouter := Router.Group("cdkey")
	cdkeyRouter.Use(func(c *gin.Context) {
		if global.GB_REDIS == nil {
			initialize.Redis()
		}
	})
	{
		cdkeyRouter.POST("createCdkey", api.CreateCdkey)
		cdkeyRouter.GET("getCdkeyDetails", api.GetCdkeyDetails)
		cdkeyRouter.GET("verifyCdkey", api.VerifyCdkey)
	}
}
