package config

type Mysql struct {
	User        string
	Password    string
	Host        string
	Port        int
	DbName      string
	Timeout     string
	MaxConn     int
	MaxIdleConn int
}
