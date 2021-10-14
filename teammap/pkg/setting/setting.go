package setting

import (
	"fmt"
	"log"
	"time"

	"github.com/go-ini/ini"
)

type App struct {
	RunMode   string
	JwtSecret string
}

var AppSetting = &App{}

type Platform struct {
	AppId            int
	TeamApisDomain   string
	MailApisDomain   string
	OnlineUserDomain string
	MailApiDomain    string
}

var ModuleDomains map[string]map[string]string

var AndroidSetting = &Platform{}
var IOSSetting = &Platform{}

type Server struct {
	HttpPort          int
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	WsPort            int
	WsPongWait        time.Duration
	WsReadBufferSize  int
	WsWriteBufferSize int
	WsWriteWait       time.Duration
	WsMaxMessageSize  int64
	IntranetIp        string
	MainRedisMode     string
	TeamRedisMode     string
}

var ServerSetting = &Server{}

type MySql struct {
	User        string
	Password    string
	Host        string
	Port        int
	DbName      string
	Timeout     string
	MaxConn     int
	MaxIdleConn int
}

var DatabaseSetting = &MySql{}

type Redis struct {
	Host         string
	Password     string
	MaxIdle      int
	MaxActive    int
	IdleTimeout  time.Duration
	PoolSize     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	DialTimeout  time.Duration
	DB           int
}

type RedisClu struct {
	Slot1        string
	Slot1Start   int
	Slot1End     int
	Slot1Slave   string
	Slot2        string
	Slot2Start   int
	Slot2End     int
	Slot2Slave   string
	Slot3        string
	Slot3Start   int
	Slot3End     int
	Slot3Slave   string
	MaxRedirects int

	ReadOnly       bool
	RouteByLatency bool
	RouteRandomly  bool

	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	DialTimeout  time.Duration

	PoolSize           int
	MinIdleConns       int
	PoolTimeout        time.Duration //等待获取连接的最大时长
	IdleTimeout        time.Duration //闲置超时
	IdleCheckFrequency time.Duration //闲置连接检查的周期

	MaxRetries int
	/*MinRetryBackoff time.Duration
	MaxRetryBackoff time.Duration*/

}

type Mongo struct {
	ApplyURI        string
	Hosts           string
	MaxConnIdleTime uint
	MaxPoolSize     uint64
	MinPoolSize     uint64
}

type Config struct {
	BadWordsURI string
	IapApiURI   string
}

var MainRedisSetting = &Redis{}
var MainRedisCluSetting = &RedisClu{}
var TeamRedisSetting = &Redis{}
var TeamRedisCluSetting = &RedisClu{}
var OldRedisSetting = &Redis{}
var MongoSetting = &Mongo{}
var ConfigSetting = &Config{}

// Setup initialize the configuration instance
func Setup(mode string) {
	var err error

	appSettingPath := "conf/sys/app.ini"

	var appCfg *ini.File
	appCfg, err = ini.Load(appSettingPath)
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse %s : %v", appSettingPath, err)
	}

	mapTo(appCfg, "app", AppSetting)
	if len(mode) > 0 {
		AppSetting.RunMode = mode
	}

	serverSettingPath := fmt.Sprintf("conf/sys/%s.ini", AppSetting.RunMode)

	var serverCfg *ini.File
	serverCfg, err = ini.Load(serverSettingPath)
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse %s: %v", serverSettingPath, err)
	}

	mapTo(serverCfg, "server", ServerSetting)
	mapTo(serverCfg, "mysql", DatabaseSetting)
	mapTo(serverCfg, "main-redis", MainRedisSetting)
	mapTo(serverCfg, "main-redisCluster", MainRedisCluSetting)
	mapTo(serverCfg, "team-redis", TeamRedisSetting)
	mapTo(serverCfg, "team-redisCluster", TeamRedisCluSetting)
	mapTo(serverCfg, "mongo", MongoSetting)
	mapTo(serverCfg, "android", AndroidSetting)
	mapTo(serverCfg, "ios", IOSSetting)
	mapTo(serverCfg, "v1-redis", OldRedisSetting)
	mapTo(serverCfg, "config", ConfigSetting)

	ModuleDomains = map[string]map[string]string{
		"android": {
			"team":    AndroidSetting.TeamApisDomain,
			"mail":    AndroidSetting.MailApisDomain,
			"user":    AndroidSetting.OnlineUserDomain,
			"mailApi": AndroidSetting.MailApiDomain,
		},
		"ios": {
			"team":    IOSSetting.TeamApisDomain,
			"mail":    IOSSetting.MailApisDomain,
			"user":    IOSSetting.OnlineUserDomain,
			"mailApi": IOSSetting.MailApiDomain,
		},
	}

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
	ServerSetting.WsPongWait = ServerSetting.WsPongWait * time.Second
	ServerSetting.WsWriteWait = ServerSetting.WsWriteWait * time.Second

	MainRedisSetting.IdleTimeout = MainRedisSetting.IdleTimeout * time.Second
	MainRedisSetting.ReadTimeout = MainRedisSetting.ReadTimeout * time.Second
	MainRedisSetting.WriteTimeout = MainRedisSetting.WriteTimeout * time.Second
	MainRedisSetting.DialTimeout = MainRedisSetting.DialTimeout * time.Second

	TeamRedisSetting.IdleTimeout = TeamRedisSetting.IdleTimeout * time.Second
	TeamRedisSetting.ReadTimeout = TeamRedisSetting.ReadTimeout * time.Second
	TeamRedisSetting.WriteTimeout = TeamRedisSetting.WriteTimeout * time.Second
	TeamRedisSetting.DialTimeout = TeamRedisSetting.DialTimeout * time.Second

	MainRedisCluSetting.ReadTimeout = MainRedisCluSetting.ReadTimeout * time.Second
	MainRedisCluSetting.WriteTimeout = MainRedisCluSetting.WriteTimeout * time.Second
	MainRedisCluSetting.DialTimeout = MainRedisCluSetting.DialTimeout * time.Second
	MainRedisCluSetting.PoolTimeout = MainRedisCluSetting.PoolTimeout * time.Second
	MainRedisCluSetting.IdleTimeout = MainRedisCluSetting.IdleTimeout * time.Second
	MainRedisCluSetting.IdleCheckFrequency = MainRedisCluSetting.IdleCheckFrequency * time.Second

	OldRedisSetting.IdleTimeout = OldRedisSetting.IdleTimeout * time.Second
	OldRedisSetting.ReadTimeout = OldRedisSetting.ReadTimeout * time.Second
	OldRedisSetting.WriteTimeout = OldRedisSetting.WriteTimeout * time.Second
	OldRedisSetting.DialTimeout = OldRedisSetting.DialTimeout * time.Second
}

// mapTo map section
func mapTo(cfg *ini.File, section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
