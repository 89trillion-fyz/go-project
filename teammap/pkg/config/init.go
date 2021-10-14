package config

import (
	"github.com/spf13/pflag"
)

// 配置文件路径
var (
//armyModelCfgPath = pflag.StringP("army", "", "conf/data/config.army.model.json", "iw config file path.")
)

const (
	DefaultPort = 0
	DefaultMode = ""
)

// 启动参数
var (
	port       = pflag.IntP("port", "", DefaultPort, "")
	mode       = pflag.StringP("mode", "", DefaultMode, "")
	testConfig = pflag.StringP("config", "", "", "iw test config file path")
)

func ParseFlag() (int, string, string) {
	pflag.Parse()

	// mode 只支持3种，参数检验，防止非法mode
	if *mode != "debug" && *mode != "test" && *mode != "release" {
		*mode = ""
	}
	return *port, *mode, *testConfig
}

func Setup() {
	//InitMoraleDecrease(moraleDecreasePath)

}
