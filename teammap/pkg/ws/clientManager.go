package ws

// The socket server main function starts the ClientManager's Start() method as a goroutine.
// Clients send requests to the ClientManager using the register, unregister and broadcast channels.
//
// The ClientManager registers clients by adding the client pointer as a key in the clients map. The map value is always true.
//
// The unregister code is a little more complicated. In addition to deleting the client pointer from the clients map,
// the ClientManager closes the clients' send channel to signal the client that no more messages will be sent to the client.
//
// The ClientManager handles messages by looping over the registered clients and sending the message to the client's send channel.
// If the client's send buffer is full, then the ClientManager assumes that the client is dead or stuck.
// In this case, the ClientManager unregisters the client and closes the websocket.

import (
	"sync"
	"teammap/pkg/errorlog"
	"teammap/pkg/myerr"
	"teammap/pkg/util"
	"teammap/protocol/protobuf/request"

	"google.golang.org/protobuf/proto"
)

const (
	PUSH_EVENT_KICKOFF   = "kickoff"
	PUSH_EVENT_BE_ATTACK = "be.attack"
	PUSH_BD_UP_FINISH    = "bd.finish"
	PUSH_EVENT_ACHIV     = "achiv.finish"
	PUSH_EVENT_TASK      = "task.progress"
)

type EventHandler func(ctx *Client, req *request.General) []byte
type BytesEventHandler func(ctx *Client, msg *[]byte) []byte

type Router struct {
	EventMap      map[string]EventHandler
	BytesEventMap map[string]BytesEventHandler
}

// 包内全局遍历
var wsManager *ClientManager

// ClientManager is a websocket manager，维护当前连接中的client信息
type ClientManager struct {
	Clients     map[*Client]bool // client -> 是否连接
	ClientsLock sync.RWMutex     // 读写锁

	Users    map[string]*Client // 登录的用户 // userId -> Client
	UserLock sync.RWMutex       // 读写锁

	WsRouter Router // event路由

	Register   chan *Client // 建立连接
	Unregister chan *Client // 断开连接
	Broadcast  chan []byte  // 广播 channel
}

// 获取当前连接数量
func (manager *ClientManager) GetClientsLen() (clientsLen int) {
	clientsLen = len(manager.Clients)

	return
}

// 添加客户端
func (manager *ClientManager) AddClients(client *Client) {
	manager.ClientsLock.Lock()
	defer manager.ClientsLock.Unlock()

	manager.Clients[client] = true
}

// 删除客户端
func (manager *ClientManager) DelClients(client *Client) {
	manager.ClientsLock.Lock()
	defer manager.ClientsLock.Unlock()

	if _, ok := manager.Clients[client]; ok {
		delete(manager.Clients, client)
	}
}

func (manager *ClientManager) InClient(client *Client) (ok bool) {
	manager.ClientsLock.RLock()
	defer manager.ClientsLock.RUnlock()

	// 连接存在，在添加
	_, ok = manager.Clients[client]

	return
}

// GetClients
func (manager *ClientManager) GetClients() (clients map[*Client]bool) {

	clients = make(map[*Client]bool)

	manager.ClientsRange(func(client *Client, value bool) (result bool) {
		clients[client] = value

		return true
	})

	return
}

// 遍历
func (manager *ClientManager) ClientsRange(f func(client *Client, value bool) (result bool)) {

	manager.ClientsLock.RLock()
	defer manager.ClientsLock.RUnlock()

	for key, value := range manager.Clients {
		result := f(key, value)
		if result == false {
			return
		}
	}

	return
}

// 获取用户的连接
func (manager *ClientManager) GetUserClient(userId string) (client *Client) {
	manager.UserLock.RLock()
	defer manager.UserLock.RUnlock()

	userKey := userId
	if value, ok := manager.Users[userKey]; ok {
		client = value
	}

	return
}

// 获取 当前连接用户数量
func (manager *ClientManager) GetUsersLen() (userLen int) {
	userLen = len(manager.Users)

	return
}

// 添加用户
func (manager *ClientManager) AddUsers(key string, client *Client) {
	manager.UserLock.Lock()
	defer manager.UserLock.Unlock()

	manager.Users[key] = client
}

// 删除用户
func (manager *ClientManager) DelUsers(client *Client) (result bool) {
	manager.UserLock.Lock()
	defer manager.UserLock.Unlock()

	key := client.UserId
	if value, ok := manager.Users[key]; ok {
		// 判断是否为相同的用户
		if value.Addr != client.Addr {

			return
		}
		delete(manager.Users, key)
		result = true
	}

	return
}

// 新客户端注册
func (manager *ClientManager) ClientRegister(client *Client) {
	manager.AddClients(client)
	// 添加用户
	if manager.InClient(client) {
		manager.AddUsers(client.UserId, client)
	}

	errorlog.Info("EventLogin 用户登录", client.Addr, client.UserId)
	// 连接后检查是否要推送消息
	manager.HandleEvent(client, "con.push", nil)
}

// 客户端注销
func (manager *ClientManager) ClientUnregister(client *Client) {
	errorlog.Info("清理 用户连接数据", client.UserId)

	manager.DelClients(client)

	// 删除用户连接
	deleteResult := manager.DelUsers(client)
	if deleteResult == false {
		// 不是当前连接的客户端
		return
	}

	// 清除redis登录数据
	manager.HandleEvent(client, "offline", nil)

	// 关闭 chan
	client.Socket.Close()
}

// client 断开连接
func (manager *ClientManager) OnDisconnectAfterTimeout(c *Client) {

	// 标记 已断开，并等待超时
	c.Offline()
}

// 注册消息
func (manager *ClientManager) RegisterEvent(event string, handler EventHandler) {
	manager.WsRouter.EventMap[event] = handler
}

// 注册Bytes消息
func (manager *ClientManager) RegisterBytesEvent(event string, handler BytesEventHandler) {
	manager.WsRouter.BytesEventMap[event] = handler
}

// 处理Event消息
func (manager *ClientManager) HandleEvent(ctx *Client, event string, msg []byte) []byte {
	handler, ok := manager.WsRouter.EventMap[event]
	if !ok {
		// not have handler
		return util.MakeErrorResponse(myerr.INVALID_EVENT)
	} else {
		req := request.General{}
		err := proto.Unmarshal(msg, &req)
		if err != nil {
			errorlog.Error("解析proto出错")
			return util.MakeErrorResponse(myerr.PROTO_UNMARSHAL_ERROR)
		} else {
			// 处理消息
			return handler(ctx, &req)
		}
	}
}

// 处理Event(二进制)消息
func (manager *ClientManager) HandleBytesEvent(ctx *Client, event string, msg []byte) []byte {
	handler, ok := manager.WsRouter.BytesEventMap[event]
	if !ok {
		// not have handler
		return util.MakeErrorResponse(myerr.INVALID_EVENT)
	} else {
		return handler(ctx, &msg)
	}
}

// 新建Manager
func NewClientManager() (clientManager *ClientManager) {
	wsManager = &ClientManager{
		Clients: make(map[*Client]bool),
		Users:   make(map[string]*Client),

		WsRouter: Router{
			EventMap:      make(map[string]EventHandler),
			BytesEventMap: make(map[string]BytesEventHandler),
		},

		Register:   make(chan *Client, 1000),
		Unregister: make(chan *Client, 1000),
		Broadcast:  make(chan []byte),
	}

	return wsManager
}

// 管道处理程序
func (manager *ClientManager) Start() {

	// client 读写初始化
	initClientReadWriteSetting()

	for {
		select {
		case conn := <-manager.Register:
			// 添加新客户端
			manager.ClientRegister(conn)

		case conn := <-manager.Unregister:
			manager.ClientUnregister(conn)

			// case message := <-manager.Broadcast:
			// do noting

			// for conn := range manager.GetClients() {
			//	select {
			//	case conn.Send <- message:
			//	default:
			//		close(conn.Send)
			//	}
			// }
		}
	}
}
