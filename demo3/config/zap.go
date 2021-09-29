package config

type Zap struct {
	Level         string `json:"level" ini:"level"`                 // 级别
	Format        string `json:"format" ini:"Format"`               // 输出
	Prefix        string `json:"prefix" ini:"Prefix"`               // 日志前缀
	Director      string `json:"director" ini:"Director"`           // 日志文件夹
	LinkName      string `json:"linkName" ini:"LinkName"`           // 软链接名称
	ShowLine      bool   `json:"showLine" ini:"ShowLine"`           // 显示行
	EncodeLevel   string `json:"encodeLevel" ini:"EncodeLevel"`     // 编码级
	StacktraceKey string `json:"stacktraceKey" ini:"StacktraceKey"` // 栈名
	LogInConsole  bool   `json:"logInConsole" ini:"LogInConsole"`   // 输出控制台
}
