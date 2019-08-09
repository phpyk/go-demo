package main

import "fmt"

func main() {
	fmt.Println(ChangeVar())
	first,last,nick := GetName()
	fmt.Println(first,last,nick)
	fmt.Println(GetConst())
	fmt.Println(Sunday,Monday,Tuesday,Wednesday,Thursday,Friday,Saturday,numberOfDays)
}
//交换变量，多重赋值
func ChangeVar() (i,j,k int) {
	var v1 int = 11
	var v2 = 12
	v3 := 13
	v1,v2,v3 = v3,v2,v1
	return v1,v2,v3
}

//匿名函数，多重返回值
func GetName () (firstName,lastName,nickName string){
	return "yue" ,"kai", "kk"
}

//常量声明
func GetConst() (p float64,s int64,w float64){
	const Pi = 3.1415926
	const (
		Size int64 = 1024
		Weight = 112.82
	)
	return Pi,Size,Weight
}
//枚举多个常量
const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	numberOfDays
)

