package ws

// 参考 https://github.com/link1st/gowebsocket

import (
	"net"
	"runtime/debug"
	"strings"
	"teammap/pkg/errorlog"
	httpPkg "teammap/pkg/http"
	"teammap/pkg/setting"
	"teammap/pkg/util"
	"teammap/protocol/protobuf"
	"time"

	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

const (
	HeartbeatTimeout      = 1 * 60                             // 用户心跳超时时间
	OfflineTimeout        = 3 * 60                             // 掉线超时， 掉线3min内重连不需要重写发login
	UseClosedNetworkError = "use of closed network connection" // 同一个用户设备，游戏中切飞行模式，服务端检测不到，socket还在，再连接建立一个新的socket，旧的报错
)

// Client is a websocket client
type Client struct {
	UserId    string `json:"id"` // userId
	IsOffline bool   // 是否离线

	Addr          string          // 客户端地址
	Socket        *websocket.Conn // 用户连接，每个client实例化一个链接
	Send          chan []byte     // 向客户端发送数据
	ExitCh        chan bool       // 信号chan，收到数据就退出 写数据 协程
	FirstTime     uint64          // 首次连接事件
	HeartbeatTime uint64          // 用户上次心跳时间
	HttpHeader    *httpPkg.HttpHeader
}

// 新Client
func NewClient(userId string, addr string, socket *websocket.Conn, firstTime uint64, header *httpPkg.HttpHeader) *Client {
	client := wsManager.GetUserClient(userId)
	if client != nil {
		oldSocket := client.Socket

		client.Socket = socket

		err := oldSocket.Close()
		if err != nil {
			errorlog.Info("close old socket err ", userId, addr, err.Error())
		}
	} else {
		client = &Client{
			UserId:        userId,
			Addr:          addr,
			Socket:        socket,
			Send:          make(chan []byte, 100),
			ExitCh:        make(chan bool),
			FirstTime:     firstTime,
			HeartbeatTime: firstTime,
		}
	}

	return client
}

var (
	// Time allowed to write a message to the peer.
	writeWait time.Duration

	// Time allowed to read the next pong message from the peer.
	pongWait time.Duration

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod time.Duration

	// Maximum message size allowed from peer.
	maxMessageSize int64
)

// 初始化配置
func initClientReadWriteSetting() {
	// Time allowed to write a message to the peer.
	writeWait = setting.ServerSetting.WsWriteWait

	// Time allowed to read the next pong message from the peer.
	pongWait = setting.ServerSetting.WsPongWait

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = setting.ServerSetting.WsMaxMessageSize
}

// 用户心跳
func (c *Client) Heartbeat(currentTime uint64) {
	c.HeartbeatTime = currentTime

	return
}

// 心跳超时
func (c *Client) IsHeartbeatTimeout(currentTime uint64) (timeout bool) {
	if c.HeartbeatTime+HeartbeatTimeout <= currentTime {
		timeout = true
	}

	return
}

// 用户下线
func (c *Client) Offline() {
	c.ExitCh <- true // 关闭write 协程

	removeFunc := func() {

	}

	// 等待心跳超时后 删除redis中在线状态
	time.AfterFunc(OfflineTimeout*time.Second, removeFunc)
}

// 发数据
func (c *Client) SendMsg(msg []byte) {
	if c == nil {
		return
	}

	defer func() {
		if r := recover(); r != nil {
			errorlog.Info("SendMsg stop:", c.UserId, c.Addr, r, string(debug.Stack()))
		}
	}()

	c.Send <- msg
}

// 从链接中进行读取
func (c *Client) Read() {
	// errorlog.Info("start read() ", c.UserId)

	// 设置读取信息的最大值
	// c.Socket.SetReadLimit(maxMessageSize)

	pingHandler := func(message string) error {
		_ = c.Socket.SetReadDeadline(time.Now().Add(pongWait)) // 设置 read 超时

		// 回复pong消息
		err := c.Socket.WriteControl(websocket.PongMessage, []byte(message), time.Now().Add(time.Second))
		if err == websocket.ErrCloseSent {
			return nil
		} else if e, ok := err.(net.Error); ok && e.Temporary() {
			return nil
		}
		return err
	}
	c.Socket.SetPingHandler(pingHandler)

	for {
		if c.IsOffline {
			// fmt.Println(" isoffline  break read() loop")
			wsManager.OnDisconnectAfterTimeout(c)
			break
		}
		// 阻塞，读取连接中的消息
		_, message, err := c.Socket.ReadMessage()
		if err != nil {
			// 如果另外一个连接已经启动，就不需要set offline
			if !strings.Contains(err.Error(), UseClosedNetworkError) {
				wsManager.OnDisconnectAfterTimeout(c)
			}
			break
		}

		_ = c.Socket.SetReadDeadline(time.Now().Add(pongWait)) // 设置read 超时
		c.Heartbeat(uint64(time.Now().Unix()))                 // 收到消息，相当于心跳一次

		go processData(c, wsManager, message)
	}

	// errorlog.Debug("exit client read() loop")
}

// 从chan中获取信息进行写入
func (c *Client) Write() {

	// errorlog.Debug("start write() ", c.UserId)

	defer func() {
		// errorlog.Debug("stop write()-3 ", c.UserId)
		if r := recover(); r != nil {
			errorlog.Debug("write stop", c.UserId, c.Addr, string(debug.Stack()), r)
		}
	}()

	defer func() {
		// errorlog.Debug("stop write()-2 ", c.UserId)
		wsManager.OnDisconnectAfterTimeout(c)
	}()

	// ticker := time.NewTicker(pingPeriod)
	//
	// defer func() {
	//	ticker.Stop()
	//	c.Socket.Close()
	// }()

FOR:
	for {
		select {
		case message, ok := <-c.Send:
			// 接收到要send的message
			if !ok {
				// closed the channel
				// c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			_ = c.Socket.SetWriteDeadline(time.Time{}) // clear any previous deadline
			_ = c.Socket.WriteMessage(websocket.BinaryMessage, message)

		case _, ok := <-c.ExitCh:
			if ok {
				break FOR
			}

			// case <-ticker.C:
			//	// 接收到定时器的心跳ping-pong
			//	c.Socket.SetWriteDeadline(time.Now().Add(writeWait))
			//	if err := c.Socket.WriteMessage(websocket.PingMessage, nil); err != nil {
			//		return
			//	}
		}
	}

	// errorlog.Debug("stop write()-1 ", c.UserId)
}

// 处理收到的消息
func processData(c *Client, manager *ClientManager, message []byte) {

	defer util.RecoverPanic()

	// protobuf 解析
	req := &protobuf.Msg{}
	err := proto.Unmarshal(message, req)
	if err != nil {
		errorlog.Info("解析proto出错", c.UserId, c.Addr)
		return
	}

	switch req.Type {
	case protobuf.MsgType_EVENT:
		// 处理消息
		event := req.GetEvent()
		// errorlog.Debug("handle event:", event)
		data := manager.HandleEvent(c, event, req.GetData())

		// 回复消息
		response := &protobuf.Msg{
			Seq:   req.Seq,
			Type:  req.Type,
			Event: event,
			Data:  data,
		}
		out, err := proto.Marshal(response)
		if err != nil {
			errorlog.Debug("make response err on handle ", event, c.UserId, c.Addr)
		}
		c.Send <- out

	case protobuf.MsgType_RAWEVENT:
		// 处理消息
		event := req.GetEvent()
		// errorlog.Debug("handle event:", event)
		data := manager.HandleBytesEvent(c, event, req.GetData())
		// 回复消息
		response := &protobuf.Msg{
			Seq:   req.Seq,
			Type:  req.Type,
			Event: event,
			Data:  data,
		}
		out, err := proto.Marshal(response)
		if err != nil {
			errorlog.Debug("make response err on handle ", event, c.UserId, c.Addr)
		}
		c.Send <- out

	case protobuf.MsgType_PUSH: // 不应该收到广播消息
		break
	}
}
