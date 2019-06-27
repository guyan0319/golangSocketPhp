package golangsocketphp

import (
	"fmt"
	"net"
	"os"
)

type SocketClient interface {
	Connect() (net.Conn, error)
}
type SocketClientTcp struct {
	Network string
	Address string
}

func (s *SocketClientTcp) Connect() (net.Conn, error) {
	tcpAddr, err := net.ResolveTCPAddr(s.Network, s.Address)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		return nil, nil
	}
	conn, err := net.DialTCP(s.Network, nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		return nil, nil
	}
	return conn, err
}
