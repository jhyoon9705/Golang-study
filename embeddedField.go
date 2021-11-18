package main

import (
	"fmt"
)

type User struct {
	Name string
	Age int
}

func (u User) String() string {
	return fmt.Sprintf("%s, %d", u.Name, u.Age)
}

type VIPUser struct {
	User		// embedded field
	VIPLevel int
}

func (v VIPUser) vipLevel() int {
	return v.VIPLevel
}

func main() {
	vip := VIPUser{ User{"Kim", 30}, 2}
	fmt.Println(vip.String()) // Kim, 30
	fmt.Println(vip.vipLevel()) // 2
}
