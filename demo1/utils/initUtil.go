package utils

import (
	"demo1/global"
	"encoding/json"
	"fmt"
	flag "github.com/spf13/pflag"
	"gopkg.in/ini.v1"
	io "io/ioutil"
	"strings"
)

func wordSepNormalizeFunc(f *flag.FlagSet, name string) flag.NormalizedName {
	from := []string{"-", "_"}
	to := "."
	for _, sep := range from {
		name = strings.Replace(name, sep, to, -1)
	}
	return flag.NormalizedName(name)
}
func loadConfig(filename string) {
	data, err := io.ReadFile(filename) //read config file
	if err != nil {
		fmt.Println("read json file error")
	}
	datajson := []byte(data)
	err = json.Unmarshal(datajson, &global.G_JSONDATA)
	if err != nil {
		fmt.Println("unmarshal json file error")
	}
	fmt.Println("global.G_JSONDATA", global.G_JSONDATA)
}
func InitData() {
	flag.StringVar(&global.G_JSONPATH, "jsonPath", "config.army.model.json", "Input Your Json Path")
	flag.CommandLine.SetNormalizeFunc(wordSepNormalizeFunc)
	// 把用户传递的命令行参数解析为对应变量的值
	flag.Parse()
	fmt.Println("jsonPathName", global.G_JSONPATH)
	err := ini.MapTo(&global.G_CONFIG, "app.ini")
	if err != nil {
		fmt.Println("Failed to parse config file: %s", err)
	}
	// 设置标准化参数名称的函数
	fmt.Println("globalConfig", global.G_CONFIG)
	//加载json数据
	loadConfig(global.G_JSONPATH)
}
