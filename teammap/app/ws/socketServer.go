package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/soheilhy/cmux"
	"log"
	"net"
	"net/http"
	"strings"
	"teammap/internal/routers"
	"teammap/pkg/config"
	"teammap/pkg/errorlog"
	"teammap/pkg/setting"
	"teammap/pkg/util"
	"teammap/pkg/ws"
)

func init() {
	// 解析命令行参数
	wsPort, wsMode, _ := config.ParseFlag()
	// 系统启动setting
	setting.Setup(wsMode)
	// 命令行参数 覆盖 setting
	if wsPort != config.DefaultPort {
		setting.ServerSetting.WsPort = wsPort
	}
	// 初始化 ip
	ips, err := util.IntranetIP()
	if err != nil || len(ips) < 1 {
		log.Fatalf("cannot get intranet ip")
	} else {
		setting.ServerSetting.IntranetIp = ips[0]
	}

	errorlog.Setup()

	// 加载业务配置文件
	config.Setup()
}

func main() {
	gin.SetMode(setting.AppSetting.RunMode)
	// gin.DefaultWriter = errorlog.Logger.Out

	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.WsPort)

	// 创建监听端口的 Listener
	l, err := net.Listen("tcp", endPoint)
	if err != nil {
		errorlog.Fatal(err.Error())
		// log.Fatal(err)
	}
	log.Printf("listening on port %s%s", setting.ServerSetting.IntranetIp, endPoint)

	// 根据创建成功的 Listener 初始化 cmux 实例
	m := cmux.New(l)

	// 注册不同协议的 *协议匹配器*
	httpL := m.Match(cmux.HTTP1Fast()) // 使用 cmux 内置函数标识 http1 协议
	//grpcL := m.Match(cmux.Any())       // 其他为 grpc 协议

	// 初始化不同协议的服务实例
	//go grpcserver.ServeGRPC(grpcL)
	go serveHttpAndWs(httpL)
	if err := m.Serve(); !strings.Contains(err.Error(), ws.UseClosedNetworkError) {
		panic(err)
	}
}

func serveHttpAndWs(l net.Listener) {
	// ws Router
	routersInit := routers.WsInitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout

	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	fmt.Println("[info] start http server for websocket")

	if err := server.Serve(l); err != cmux.ErrListenerClosed {
		panic(err)
	}
}
