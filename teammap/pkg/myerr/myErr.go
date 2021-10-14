package myerr

import "runtime/debug"

// Api handle status
type MyErr struct {
	Code    int
	Message string
	Level   Level
	Stack   []byte
}

type Level uint32

const (
	ErrorLevel Level = iota
	WarnLevel
	InfoLevel
	DebugLevel
	SuccessLevel
)

// 附加stack信息
func (e *MyErr) WithStack() *MyErr {
	err := *e
	err.Stack = debug.Stack()
	return &err
}

// get error information
func (e *MyErr) Error() string {
	return e.Message
}

func (e *MyErr) Detail() string {
	msg := e.Message
	if len(e.Stack) > 0 {
		msg += string(e.Stack)
	}
	return msg
}

var (
	SUCCESS = &MyErr{Code: 200, Message: "OK", Level: SuccessLevel}
	FAILED  = &MyErr{Code: 500, Message: "Fail", Level: InfoLevel}

	// 缺少必选参数
	LACK_OF_HEADER  = &MyErr{Code: 1001, Message: "lack of header", Level: InfoLevel}
	LACK_OF_PARAMS  = &MyErr{Code: 1002, Message: "lack of params", Level: InfoLevel}
	LACK_OF_SUBTYPE = &MyErr{Code: 1003, Message: "lack of subtype", Level: InfoLevel}
	LACK_OF_FLAG    = &MyErr{Code: 1004, Message: "lack of flag", Level: InfoLevel}
	LACK_OF_VALUES  = &MyErr{Code: 1005, Message: "lack of values", Level: InfoLevel}
	LACK_OF_EXT     = &MyErr{Code: 1006, Message: "lack of ext", Level: InfoLevel}

	// 参数值错误
	INVALID_HEADER   = &MyErr{Code: 1101, Message: "invalid header", Level: InfoLevel}
	INVALID_EVENT    = &MyErr{Code: 1102, Message: "not support event", Level: InfoLevel} // 不支持的事件
	INVALID_PARAMS   = &MyErr{Code: 1103, Message: "invalid params", Level: InfoLevel}
	INVALID_SUBTYPE  = &MyErr{Code: 1104, Message: "invalid subtype", Level: InfoLevel}
	INVALID_TROOP_ID = &MyErr{Code: 1105, Message: "invalid troop id", Level: InfoLevel}
	INVALID_SKU      = &MyErr{Code: 1106, Message: "invalid sku", Level: InfoLevel}
	INVALID_GOODS_ID = &MyErr{Code: 1107, Message: "invalid goods id", Level: InfoLevel}
	INVALID_POS      = &MyErr{Code: 1108, Message: "invalid pos", Level: InfoLevel}
	INVALID_PLATFORM = &MyErr{Code: 1109, Message: "invalid platform", Level: InfoLevel}
	INVALID_LEVEL    = &MyErr{Code: 1110, Message: "invalid level", Level: InfoLevel}
	INVALID_VALUES   = &MyErr{Code: 1111, Message: "invalid values", Level: InfoLevel}
	INVALID_FLAG     = &MyErr{Code: 1112, Message: "invalid flag", Level: InfoLevel}
	INVALID_SEASON   = &MyErr{Code: 1113, Message: "invalid season", Level: InfoLevel}
	INVALID_AMOUNT   = &MyErr{Code: 1114, Message: "invalid amount", Level: InfoLevel}
	INVALID_USERID   = &MyErr{Code: 1115, Message: "invalid userid", Level: InfoLevel}

	PROTO_UNMARSHAL_ERROR = &MyErr{Code: 3007, Message: "invalid protobuf format", Level: InfoLevel} // protobuf 解析错误
)
