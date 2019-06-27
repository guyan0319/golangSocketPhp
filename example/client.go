package main

import (
	"encoding/json"
	"fmt"
	"golangSocketPhp"
	"os"
)

func main() {
	//客户端主动连接服务器
	socket := golangsocketphp.SocketClientTcp{Network: "tcp", Address: "127.0.0.1:8181"}
	conn, err := socket.Connect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		return
	}
	defer conn.Close() //关闭
	type user struct {
		Controller string                 `json:"controller"`
		Action     string                 `json:"action"`
		Params     map[string]interface{} `json:"params"`
	}
	param := make(map[string]interface{})
	param["id"] = 1
	param["name"] = "golang"
	u := user{Controller: "demo.User", Action: "Test", Params: param}
	data, err := json.Marshal(u)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		return
	}
	conn.Write([]byte(string(data)))

	buf := make([]byte, 2048) //缓冲区
	n, err := conn.Read(buf)  //读取数据
	if err != nil {
		fmt.Println(err)
		return
	}
	result := buf[:n]
	fmt.Printf("receive data string[%d]:%s\n", n, string(result))

}
