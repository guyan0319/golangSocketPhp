package demo

import (
	"fmt"
	"net"
)

type User struct {
	Name string
}

func (a *User) Test(conn *net.Conn, params interface{}) {
	fmt.Println(params.(map[string]interface{}))
	p := params.(map[string]interface{})
	fmt.Println(p["id"], p["name"])
	(*conn).Write([]byte("hi jerry"))
}
