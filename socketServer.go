package golangsocketphp

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"reflect"
	"sync"
)

/*
 *  socketserver
 * author Guo Zhiqiang
 * datetime 2019/6/25 9:02
 */

type SocketServer struct {
	sync.RWMutex
	//The network must be "tcp", "tcp4", "tcp6", "unix" or "unixpacket".
	Network string
	// If the port in the address parameter is empty or "0", as in
	// "127.0.0.1:" or "[::1]:0", a port number is automatically chosen.
	Address string
	//业务逻辑struct
	Objects []interface{}
	//用来存储业务逻辑的方法
	funcMap map[string]reflect.Value
}

func (s *SocketServer) Register() {
	//建立socket，监听端口
	netListen, err := net.Listen(s.Network, s.Address)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		return
	}
	//填充funcMap
	s.Populate()
	defer netListen.Close()
	fmt.Println("Waiting for clients")
	//定义一个WaitGroup 实现并发控制
	var wg sync.WaitGroup
	//计数器设置为2000
	wg.Add(2000)
	for {
		//用来返回一个新的连接，进行后续操作
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}
		fmt.Println(conn.RemoteAddr().String(), " tcp connect success")
		//处理
		go s.ConnHandle(conn, wg)
	}
}
func (s *SocketServer) Populate() {
	s.funcMap = make(map[string]reflect.Value, 0)
	for _, v := range s.Objects {
		reflectValue := reflect.ValueOf(v)
		reflectType := reflect.TypeOf(v).Elem()
		vft := reflectValue.Type()

		//遍历路由器的方法，并将其存入控制器映射变量中
		for i := 0; i < reflectValue.NumMethod(); i++ {
			mName := vft.Method(i).Name
			s.funcMap[reflectType.String()+mName] = reflectValue.Method(i)
		}
	}
}

func (s *SocketServer) ConnHandle(conn net.Conn, wg sync.WaitGroup) {
	buffer := make([]byte, 2048)
	//减一
	defer wg.Done()
	for {
		//用于接收数据，返回接收的长度或者返回错误，是TCPConn的方法
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fatal error: ", err)
			return
			//fmt.Println(conn.RemoteAddr().String(), " connection error: ", err)
			//return
		}
		//读取数据到data
		data := s.Read(buffer[:n])
		if _, ok := data["controller"]; !ok {
			fmt.Fprintf(os.Stderr, "Fatal error: ", "The controller does not exist")
			return
		}
		if _, ok := data["action"]; !ok {
			fmt.Fprintf(os.Stderr, "Fatal error: ", "The action does not exist")
			return
		}
		if _, ok := data["params"]; !ok {
			fmt.Fprintf(os.Stderr, "Fatal error: ", "The params does not exist")
			return
		}
		routerReflect, ok := s.funcMap[data["controller"].(string)+data["action"].(string)]
		if !ok {
			fmt.Fprintf(os.Stderr, "Fatal error: ", "The func does not exist")
			return
		}
		routers := []reflect.Value{reflect.ValueOf(&conn), reflect.ValueOf(data["params"])}
		//根据controller  action  params调用相应的方法
		routerReflect.Call(routers)
		//fmt.Println(routerReflect)
	}
}

func (s *SocketServer) Read(msg []byte) map[string]interface{} {
	var data map[string]interface{}
	err := json.Unmarshal(msg, &data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		return nil
	}
	return data
}
