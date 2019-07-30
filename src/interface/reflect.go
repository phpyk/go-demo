package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i int16
	i = 1000
	t := reflect.TypeOf(i)
	r := reflect.Type(t).Name()
	v := reflect.ValueOf(i)
	fmt.Printf("%v,%v,%v\n",t,r,v)

	p := Person{"yuekai", 30, "man",}
	//获取p的类型
	pt := reflect.TypeOf(p)
	//获取p第一个、第二个、第三个字段名
	fn0 := pt.Field(0).Name
	fn1 := pt.Field(1).Name
	fn2 := pt.Field(2).Name
	v1,ok := pt.FieldByName("name")
	if ok {
		fmt.Println("v1:",v1)
	}
	fmt.Printf("type:%v,field:%v %v %v\n",pt,fn0,fn1,fn2)


	//获取p的值
	pv := reflect.ValueOf(p)
	//获取p的类别
	pk := pv.Kind()  //struct
	fmt.Printf("pv: %+v\n",pv)
	fmt.Printf("pk: %v\n",pk)
	fmt.Println(pk == reflect.Struct)//true

	//修改反射的值
	var x float64 = 3.4
	pp := reflect.ValueOf(&x)
	vv := pp.Elem()
	vv.SetFloat(7.1)
	fmt.Printf("yuekai:%+v\n",vv)

	pPointer := &p
	pPointer.name = "Beiqiaoling"
	pPointer.gender = "woman"
	fmt.Printf("Beiqiaoling:%+v",p)
}

type Person struct {
	name string
	age int
	gender string
}