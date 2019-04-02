package main

import "fmt"
import "math"

func main() {
	compare()
	change()
	floatEqual()
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
	isEqual := math.Fdim(fvalue1,fvalue2) < p;
	fmt.Println(isEqual)
}
