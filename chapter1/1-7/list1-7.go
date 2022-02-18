package main

import (
	"fmt"
	"module_hello/hello"
)

func main() {
	name := hello.Input("type your name")
	fmt.Println("Hello, " + name + "!!")
}

// モジュールの参考
// https://qiita.com/tkj06/items/a5f79417935100045650