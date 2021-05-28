package controller

import (
	"go-project/demo3/service"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user")
	{
		userRouter.POST("registerUser", service.RegisterUser)
		userRouter.POST("login", service.Login)
	}

}
