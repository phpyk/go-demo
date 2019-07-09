/**
初识go语言
 */
package main

import "fmt"

func main() {
	fmt.Println("hello world")
	fmt.Println(getName())
	fname,lname,nname := getName()
	fmt.Println("fname:"+fname+" lname:"+lname+" nname:"+nname)

	//匿名函数
	getAge := func() int {
		return 30
	}
	fmt.Println(nname+" age is",getAge())
}
//多返回值，不需要每个都赋值
func getName() (firstName,lastName,nickNmae string) {
	firstName = "yue"
	lastName = "kai"
	nickNmae = "kk"
	return
}

