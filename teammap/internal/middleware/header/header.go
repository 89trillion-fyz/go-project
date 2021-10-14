package header

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

	httpPkg "iw/pkg/http"
	"iw/pkg/myerr"
	"iw/pkg/util"

	"github.com/gin-gonic/gin"
)

func HandleHttpHeader() gin.HandlerFunc {
	return func(c *gin.Context) {

		status := myerr.SUCCESS

		info := &httpPkg.HttpHeader{}
		basicInfo := c.Request.Header.Get("Basic-Info")
		if len(basicInfo) < 1 {
			status = myerr.LACK_OF_HEADER
			response(status, c)
			c.Abort()
			return
		}

		decode, err := base64.StdEncoding.DecodeString(basicInfo)
		if err != nil {
			status = myerr.INVALID_HEADER
			response(status, c)
			c.Abort()
			return
		}

		err = json.Unmarshal(decode, info)
		if err != nil {
			status = myerr.INVALID_HEADER
			response(status, c)
			c.Abort()
			return
		}

		info.IP = c.ClientIP()
		info.Cty = util.GetCountryCodeByIp(info.IP)
		info.Ap = util.GetAppIdByPlatform(info.Pf)

		// 合法 header 放在 gin.Context 里
		c.Set("headers", info)

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
