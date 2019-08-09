package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := &MyMux{}
	http.ListenAndServe(":9090", mux)
}

type MyMux struct {

}

func (p *MyMux) ServeHTTP(w http.ResponseWriter,r *http.Request)  {
	if r.URL.Path == "/" {
		sayHelloToUser(w,r)
		return
	}
	http.NotFound(w, r)
	return
}

func sayHelloToUser(w http.ResponseWriter,r *http.Request) {
	_,err := fmt.Fprintf(w,"hello my route!")
	if err!= nil{
		log.Fatal("sayHelloToUser error:",err)
	}
}