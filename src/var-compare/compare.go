package main

import "fmt"
import "math"
import "reflect"

func main() {
	compare()
	change()
	floatEqual()
	complexComp()
}

//比较
func compare(){
	i,j := 1,2
	if i == j {
		fmt.Println("result: i and j are equal")
	}else {
		fmt.Println("result: i and j are not equal")
	}
}
func change () {
	var fvalue1 float32
	fvalue1 = 11.22		//float32
	fvalue2 := 11.22 	//float64

	fmt.Println(fvalue1 == float32(fvalue2))
}

func floatEqual (){
	var fvalue1 float64
	fvalue1 = 11.22		//float32
	fvalue2 := 11.22 	//float64

	p := 0.000001		//自定义精度
	isEqual := math.Dim(fvalue1,fvalue2) < p;
	fmt.Println(isEqual)
}

func complexComp(){
	var v1 complex64;
	v1 = 12 + 3i;		//complex128
	v2 := 12+3i;		//complex128
	v3 := complex(12,3);//complex128
	fmt.Println(reflect.TypeOf(v1),reflect.TypeOf(v2),reflect.TypeOf(v2),reflect.TypeOf(v3))
	fmt.Println(complex128(v1) == v2)
	fmt.Println(v2==v3)
	fmt.Println(real(v1))	//获取复数实部
	fmt.Println(imag(v1))	//获取复数虚部
}