package routers

/*
	router, 路由转发控制。 定义 http请求 -》 controller 层接口
*/

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	ctrlWs "teammap/internal/ctrl/ws"
	httpPkg "teammap/pkg/http"
	"teammap/pkg/myerr"
	"teammap/pkg/setting"
	"teammap/pkg/ws"
	"time"
)

var wsManager *ws.ClientManager

// InitRouter initialize routing information
func WsInitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	wsManager = InitWsManager()
	go wsManager.Start()

	api := r.Group("./")
	api.GET("/metrics", gin.WrapH(promhttp.Handler()))

	{
		api.GET("/ws", serveWs)
	}

	return r
}

func serveWs(c *gin.Context) {
	appG := httpPkg.Gin{C: c}
	UserId := c.GetString("userId") // 从 token里解析的userId
	var header *httpPkg.HttpHeader
	// 升级协议
	upgrader := &websocket.Upgrader{
		ReadBufferSize:  setting.ServerSetting.WsReadBufferSize,
		WriteBufferSize: setting.ServerSetting.WsWriteBufferSize,
		CheckOrigin: func(r *http.Request) bool {
			if r.Method != "GET" {
				fmt.Println("method is not GET")
				return false
			}
			if r.URL.Path != "/ws" {
				fmt.Println("path error")
				return false
			}

			return true
		},
	}

	// change the request to websocket model
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		appG.Response(http.StatusNotFound, myerr.FAILED, nil)
		return
	}

	currentTime := uint64(time.Now().Unix())
	// 新client连接进来，每新增一个连接初始化一个客户端进行维护
	client := ws.NewClient(UserId, conn.RemoteAddr().String(), conn, currentTime, header)
	// 保存连接状态

	// 使用两个协程处理客户端请求数据和向客户端发送数据，读取分离，减少收发数据堵塞的可能
	go client.Read()  // 客户端读取
	go client.Write() // 客户端写入

	log.Print("new connection from ", client.Addr)
	wsManager.Register <- client
}

// 声明event 和 对应响应函数
func InitWsManager() *ws.ClientManager {
	// init manager
	var manager = ws.NewClientManager()

	// 内部接口，不对客户端提供
	{
		// 断线时，删除连接信息
		manager.RegisterEvent("wstest", ctrlWs.WsTest)
	}
	return manager
}
