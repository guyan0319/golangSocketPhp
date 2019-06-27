package golangsocketphp

import (
	"net"
)

type SocketEvent struct {
}

//发送数据
func (s *SocketEvent) OnMessage(conn *net.Conn, message string) {
	(*conn).Write([]byte(message))
}
