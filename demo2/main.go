package main

import (
	"go-project/demo2/router"
)

func main() {
	r := router.Routers()
	r.Run(":8081")
}
