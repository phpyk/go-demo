package main

import (
	"fmt"
	"simplemath"
)

func main() {
	ret,err := simplemath.IntAdd(99999,99999)
	fmt.Printf("result:%v, error: %v \n",ret,err)

	totalAmount := manyArgs(1,2,3,4,5,6,7)
	fmt.Println("args sum :",totalAmount)
	totalAmount2 := manyArgs(1,2,3,)
	fmt.Println("args sum2 :",totalAmount2)
}

//不定参数
func manyArgs(args ...int) (sum int) {
	for _,arg := range args {
		sum += arg
	}
	return sum
}

