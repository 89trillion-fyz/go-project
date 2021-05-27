package router

import (
	"go-project/demo3/api"

	"github.com/gin-gonic/gin"
)

func InitCdkRouter(Router *gin.RouterGroup) {
	cdkeyRouter := Router.Group("cdkey")
	{
		cdkeyRouter.POST("createCdkey", api.CreateCdkey)
		cdkeyRouter.GET("getCdkeyDetails", api.GetCdkeyDetails)
		cdkeyRouter.GET("verifyCdkey", api.VerifyCdkey)
	}
}
