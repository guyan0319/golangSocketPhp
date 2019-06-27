package golangsocketphp

import (
	"fmt"
	"net"
	"os"
	"time"
)

type SocketClientTimeout struct {
	Network string
	Address string
	Timeout time.Duration
}

//timeout
func (s *SocketClientTimeout) Connect() (net.Conn, error) {
	conn, err := net.DialTimeout(s.Network, s.Address, s.Timeout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		return nil, nil
	}
	return conn, err
}
