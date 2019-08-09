package main

import (
	"fmt"
	"reflect"
)

/**
method继承
 */
type Human struct {
	name string
	age int
	phone string
}
type Student struct {
	Human
	school string
}
type Employee struct {
	Human
	company string
}

func (h *Human) SayHi()  {
	fmt.Printf("Hi,I am %s,I am a %s you can call me %s \n",h.name,reflect.TypeOf(h) ,h.phone)
}
func (e *Employee) SayHi()  {
	fmt.Printf("Hi,I am %s,I am a Employee,I work in %s you can call me %s \n",e.name,e.company ,e.phone)
}
func main() {
	YueKai := Student{Human{"YueKai",25,"17505818455"},"家里蹲大学"}
	BeiQiaoling := Employee{Human{"BeiQiaoling", 26, "15088615227"}, "家里干科技有限公司",}
	YueKai.SayHi()
	BeiQiaoling.SayHi()
}