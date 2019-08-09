package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm() //解析参数
	if err != nil {
		fmt.Println("parse form error",err)
	}
	name := r.Form["name"]
	//map[name:[beiqiaoling] age:[30] url_long:[111 222]]
	//form是个map结构，每个参数的值都是字符串切片类型 []string
	fmt.Println("form:",r.Form)	//form 内容
	fmt.Println("path:",r.URL.Path)
	fmt.Println("scheme:",r.URL.Scheme)
	fmt.Println("url_long:",r.Form["url_long"])
	fmt.Println("url_long type:",reflect.TypeOf(r.Form["url_long"]))
	for k,v := range r.Form {
		fmt.Printf("key:%v val:%v value_type:%v \n",k,strings.Join(v,""),reflect.TypeOf(v))
	}
	//strings.Join(name,",") 使用逗号分隔name，相当于php的implode函数
	n,er := fmt.Fprintf(w,"Hello %v",strings.Join(name,","))
	if er != nil {
		fmt.Println("er:",er)
	}
	fmt.Println("n:",n)
}
func main() {
	http.HandleFunc("/hello",sayHelloName)	//设置访问路由
	/**
	如果你以前是 PHP 程序员，那你也许就会问，我们的 nginx、apache 服务器不需要吗？
	Go 就是不需要这些，因为他直接就监听 tcp 端口了，做了 nginx 做的事情，
	然后 sayhelloName 这个其实就是我们写的逻辑函数了，跟 php 里面的控制层（controller）函数类似。
	 */
	err := http.ListenAndServe(":9090", nil) 	//设置监听端口
	if err != nil {
		log.Fatal("ListenAndServe:",err)
	}
}