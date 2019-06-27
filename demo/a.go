package demo

import (
	"fmt"
	"golangSocketPhp"
	"net"
)

type User struct {
	Name string
}

//业务逻辑处理
func (a *User) Test(conn *net.Conn, params interface{}) {
	fmt.Println(params.(map[string]interface{}))
	p := params.(map[string]interface{})
	//fmt.Println(p["id"], p["name"])
	event := golangsocketphp.SocketEvent{}
	if _, ok := p["name"]; !ok {
		event.OnMessage(conn, "name not found")
		return
	}
	event.OnMessage(conn, p["name"].(string))

}
