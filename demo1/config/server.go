package config

type Server struct {
	HttpPort          string
	WsPort            int
	ReadTimeout       int
	WriteTimeout      int
	WsPongWait        int
	WsReadBufferSize  int
	WsWriteBufferSize int
	WsWriteWait       int
	WsMaxMessageSize  int
	MainRedisMode     string
	TeamRedisMode     string
}
