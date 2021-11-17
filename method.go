package main

import (
	"fmt"
)

type account struct {
	balance int
}

type myInt int

func withdrawFunc(a *account, amount int) {
	a.balance -= amount
}

func (a *account) withdrawMethod(amount int) { // a *account: reciever
	a.balance -= amount
}

func (m myInt) Add(b int) myInt{ // 패키지 내 별칭 타입 method
	rst := int(m) + b
	return myInt(rst)
}

func addFunc(m myInt, b int) myInt{
	rst := int(m) + b
	return myInt(rst)
}
 
func main() {
	a := &account{100} // a's type: *account
	 
	withdrawFunc(a, 30)
	fmt.Println(a.balance) // 70
	
	a.withdrawMethod(30)
	fmt.Println(a.balance) // 40
	
	var b myInt
	b = b.Add(10)
	fmt.Println(b) // 10
}
