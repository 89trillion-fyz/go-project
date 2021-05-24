package main

import (
	"demo1/core"
	"demo1/utils"
	"fmt"
)

func init() {
	utils.InitData()
}
func main() {
	core.RunServer()
	//r.Run(":"+glo/balConfig.Server.HttpPort) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	fmt.Println("server is starting ....")
}
