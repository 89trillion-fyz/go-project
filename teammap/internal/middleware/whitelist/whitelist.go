package whitelist

import (
	"iw/pkg/util"
	"net/http"

	httpPkg "iw/pkg/http"
	"iw/pkg/myerr"

	"github.com/gin-gonic/gin"
)

func HandleCheckWhiteIP() gin.HandlerFunc {
	return func(c *gin.Context) {
		status := util.CheckIp(c.ClientIP(), httpPkg.WhiteList)
		err := myerr.SUCCESS
		if !status {
			err = myerr.ERROR_NO_ACCESS
			response(err, c)
			c.Abort()
			return
		}
		c.Next()
	}
}

func response(status *myerr.MyErr, c *gin.Context) {
	if status != myerr.SUCCESS {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": status.Code,
			"msg":  status.Message,
			"data": "",
		})
	}
}
