package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-project/demo2/core"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	Router.POST("/getResult", func(c *gin.Context) {
		//var data map[string]string
		var body map[string]string
		_ = c.ShouldBindJSON(&body)
		fmt.Println("body", body)
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("error", err)
				c.JSON(200, gin.H{
					"msg":    err,
					"result": 0,
				})
			}
		}()
		c.JSON(200, gin.H{
			"msg":    "ok",
			"result": core.GetResult(body["str"]),
		})
	})
	return Router
}
