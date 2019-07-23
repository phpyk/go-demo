package main

import "fmt"

func main() {
	//数组定义的几种方法
	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10,}
	var myArray2  = [...]int{1,2,3,5}
	//基于数组创建一个数组切片
	//var mySlice []int = myArray[:5]
	//数组遍历
	for index, value := range myArray {
		fmt.Printf("myArray[%d]=%d \n", index, value)
	}

	for i, v := range myArray2 {
		fmt.Printf("myArray2[%d]=%d \n", i, v)
	}
}