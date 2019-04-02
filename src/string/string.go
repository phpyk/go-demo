package main

import "fmt"
import (
	"reflect"
	"unicode/utf8"
)

func main(){
	var stringx = "Hello"
	fmt.Println(stringx[0])
	ctype := reflect.TypeOf(stringx[0])
	fmt.Println("%s",ctype)

	var stringy = "world"

	fmt.Println(stringx +" "+ stringy + "\r")
	fmt.Println(len(stringx),len(stringy),"\r")

	//遍历字符串
	//方法一
	fmt.Println("遍历字符串方法一：")
	var str = "HelloWorld你好世界"
	for i:=0;i<len(str);i++ {
		//fmt.Println(i, str[i])
		//fmt.Printf("%q", str[i])
		fmt.Print(string(str[i]))
	}

	//方法二
	fmt.Println("遍历字符串方法二：")
	//unicode字符遍历
	var str2 = "HelloWorld你好世界"
	for i,ch := range str2 {
		fmt.Println(i,ch,string(ch))
	}
	fmt.Println(utf8.RuneCountInString(str2))
	fmt.Println(reflect.TypeOf('a'))
	fmt.Println(reflect.TypeOf("a"))

}
