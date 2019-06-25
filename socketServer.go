package golangsocketphp

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
)

/*
 *  socketserver
 * author Guo Zhiqiang
 * datetime 2019/6/25 9:02
 */

type SocketServer struct {
	Network string
	Address string
	Objects []interface{}
	funcMap map[string]string
}

func (s *SocketServer) Register() {
	//建立socket，监听端口
	netListen, err := net.Listen(s.Network, s.Address)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		return
	}
	//填充funcMap
	s.populate()

	defer netListen.Close()
	fmt.Println("Waiting for clients")
	for {
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}
		fmt.Println(conn.RemoteAddr().String(), " tcp connect success")
		s.connHandle(conn)
	}
}
func (s *SocketServer) populate() {
	for v := range s.Objects {

		fmt.Println(v)

	}

}

func (s *SocketServer) connHandle(conn net.Conn) {
	buffer := make([]byte, 2048)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println(conn.RemoteAddr().String(), " connection error: ", err)
			return
		}
		//fmt.Println(string(buffer))
		s.read(buffer[:n])
		//fmt.Println(conn.RemoteAddr().String(), "receive data string:\n", string(buffer[:n]))
		//words := "hi, john"
		//conn.Write([]byte(words))
		//time.Sleep(time.Second * 1)
	}
}

func (s *SocketServer) read(msg []byte) {
	var data map[string]interface{}
	err := json.Unmarshal(msg, &data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}
