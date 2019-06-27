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
	Network string
	Address string
	Objects []interface{}
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
	//定义一个WaitGroup
	var wg sync.WaitGroup
	//计数器设置为2000
	wg.Add(2000)
	for {
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}
		fmt.Println(conn.RemoteAddr().String(), " tcp connect success")
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
	defer wg.Done()
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fatal error: ", err)
			return
			//fmt.Println(conn.RemoteAddr().String(), " connection error: ", err)
			//return
		}
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
		routerReflect.Call(routers)
		fmt.Println(routerReflect)
		//fmt.Println(conn.RemoteAddr().String(), "receive data string:\n", string(buffer[:n]))
		//words := "hi, john"
		//words := conn.RemoteAddr().String()
		//conn.Write([]byte(words))
		//time.Sleep(time.Second * 1)
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
