package main

import (
	//"golangSocketPhp"
	"golangSocketPhp"
	"golangSocketPhp/demo"
)

func main() {

	webSocket := golangsocketphp.SocketServer{Network: "tcp", Address: "localhost:8181"}
	webSocket.Objects = []interface{}{&demo.User{}}
	webSocket.Register()
}
