package core

import "go-project/demo3/router"

func RunServer() {
	r := router.Routers()
	r.Run(":8083")
}
