package router

import (
	"go-project/demo3/api"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user")
	{
		userRouter.POST("registerUser", api.RegisterUser)
		userRouter.POST("login", api.Login)
	}

}
