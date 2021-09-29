package controller

import (
	"go-project/demo3/service"

	"github.com/gin-gonic/gin"
)

func InitCdkRouter(Router *gin.RouterGroup) {
	cdkeyRouter := Router.Group("cdkey")
	{
		cdkeyRouter.POST("createCdkey", service.CreateCdkey)
		cdkeyRouter.GET("getCdkeyDetails", service.GetCdkeyDetails)
		cdkeyRouter.GET("verifyCdkey", service.VerifyCdkey)
	}
}
