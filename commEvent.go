package golangsocketphp

import "net"

type CommEvent interface {
	OnMessage(*net.Conn, string)
}
