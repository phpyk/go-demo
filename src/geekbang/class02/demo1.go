package main

import (
	"flag"
	"fmt"
)

var name string
func main() {
	flag.Parse()
	fmt.Println("hello",name)
}

func init() {
	flag.StringVar(&name,"name","everyone","Please input user name")
}