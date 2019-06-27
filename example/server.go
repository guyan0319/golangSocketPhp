package main

import (
	"golangSocketPhp"
	"golangSocketPhp/demo"
)

func main() {
	webSocket := golangsocketphp.SocketServer{Network: "tcp", Address: "localhost:8181"}
	//注册业务逻辑处理的结构体
	webSocket.Objects = []interface{}{&demo.User{}}
	webSocket.Register()

}
