package main

import (
	"fmt"
	"go-project/demo3/core"
	"go-project/demo3/global"
	"go-project/demo3/initialize"

	"gopkg.in/ini.v1"
)

func init() {
	err := ini.MapTo(&global.GB_GONFIG, "app.ini")
	if err != nil {
		fmt.Println("Failed to parse config file: %s", err)
	}
	initialize.Redis()
}
func main() {
	core.RunServer()
}
