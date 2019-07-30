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

	list := make(List,3)
	list[0] = 1
	list[1] = "hello"
	list[2] = People{"yuekai",30,"17505818455"}

	for index,element := range list {
		if value,ok := element.(int);ok {
			fmt.Printf("list[%d] is an int and its value is %d\n",index,value)
		}else if value,ok := element.(string);ok {
			fmt.Printf("list[%d] is a string and its value is %s\n",index,value)
		}else if value,ok := element.(People);ok {
			fmt.Printf("list[%d] is a People and its value is %s\n",index,value)
		}else {
			fmt.Printf("list[%d] is a different type\n",index)
		}
	}

	for index, element := range list{
		switch value,ok := element.(type) {
			case int:
				fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
			case string:
				fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
			case People:
				fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
			default:
				fmt.Printf("list[%d] is of a different type", index)
		}
	}
}

type Element interface {}
type People struct {
	name string
	age int
	phone string
}
type List[] Element;
// 通过这个方法 Human 实现了 fmt.Stringer
func (h People) String() string {
	return "❰"+h.name+" - "+strconv.Itoa(h.age)+" years - ☎️ " +h.phone+"❱"
}


