package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := '9'
	fmt.Println(a)
	fmt.Println(a - 0)
	println(reflect.TypeOf(a - '0'))
	fmt.Println(a - '0')
}
