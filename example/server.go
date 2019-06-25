package main

import (
	//"golangSocketPhp"
	"golangSocketPhp/demo"
	"golangSocketPhp"
)

func main() {
	//var webSocket golangsocketphp.SocketServer
	//webSocket.Network="tcp"
	//webSocket.Address="localhost:8181"
	webSocket :=golangsocketphp.SocketServer{Network:"tcp",Address:"localhost:8181"}

	webSocket.Objects=[]interface{}{&demo.User{}}

	webSocket.Register()
}
