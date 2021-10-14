package wsctrl

import (
	"teammap/pkg/util"
	"teammap/pkg/ws"
	"teammap/protocol/protobuf/request"
)

func WsTest(ctx *ws.Client, req *request.General) []byte {
	return util.MakeGeneralResponse(0, "ok")
}
