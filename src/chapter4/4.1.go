package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090",nil)
	if err != nil {
		log.Fatal("ListenAndServe error:",err)
	}
}

func login(w http.ResponseWriter, r *http.Request) {

	//获取请求的方法 r.Method
	fmt.Println(r.Method)
	if r.Method == "GET" {
		t,_ := template.ParseFiles("login.gtpl")
		log.Println(t.Execute(w,nil))
	} else {

		// 解析 url 传递的参数，对于 POST 则解析响应包的主体（request body）
		err := r.ParseForm()
		if err != nil {
			log.Fatal("ParseForm error:", err)
		}

		//从form拿到的数据，是string切片类型
		userName := r.Form["username"]
		password := r.Form["password"]

		fmt.Printf("userName: %v \n",userName)
		fmt.Printf("password: %v \n",password)
		if userName[0] == "yuekai" && password[0] == "123456" {
			//输出到客户端 fmt.Fprintf
			_,err := fmt.Fprintf(w,"登录成功")
			if err != nil {
				log.Fatal("login error:",err)
			}
		}
	}
}