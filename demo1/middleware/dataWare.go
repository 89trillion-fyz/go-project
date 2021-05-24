package middleware

import (
	"demo1/global"
	"demo1/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GlobalDataHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if global.G_JSONDATA == nil {
			fmt.Println("init global data")
			utils.InitData()
		}
		c.Next()
	}
}
