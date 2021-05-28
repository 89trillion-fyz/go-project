package main

import (
	"go-project/demo3/controller"
)

func init() {

}
func main() {
	r := controller.Routers()
	r.Run(":8083")
}
