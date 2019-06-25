package demo

import "fmt"

type User struct {
	Name string
}
func (a *User) test()  {
	fmt.Println("test")
}