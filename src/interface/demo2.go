package main

import (
	"fmt"
	"strconv"
)

func main() {
	//空interface,可用来存储任意类型
	var a interface{}

	var i int = 5
	s := "hello world"
	a = s
	a = i

	fmt.Printf("%v\n",a)
	Bob := People{"Bob", 39, "000-7777-XXX"}
	fmt.Println(Bob)
}

type People struct {
	name string
	age int
	phone string
}
// 通过这个方法 Human 实现了 fmt.Stringer
func (h People) String() string {
	return "❰"+h.name+" - "+strconv.Itoa(h.age)+" years -  ✆ " +h.phone+"❱"
}

