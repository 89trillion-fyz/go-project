package app

import (
	"github.com/gin-gonic/gin"
	"teammap/pkg/errorlog"
	"teammap/pkg/myerr"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
type ResponseWithNoData struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode int, status *myerr.MyErr, data interface{}) {
	var r interface{}
	if data == nil {
		r = ResponseWithNoData{
			Code: status.Code,
			Msg:  status.Message,
		}
	} else {
		r = Response{
			Code: status.Code,
			Msg:  status.Message,
			Data: data,
		}
	}

	g.C.JSON(httpCode, r)

	errorlog.LogMyError(status)
	return
}
