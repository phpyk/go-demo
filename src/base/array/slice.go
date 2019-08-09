package main

import "fmt"

func main() {

	//定义数组
	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10,}
	/**
	基于数组创建一个数组切片
	 */

	//创建数组切片  myArray[first:last]
	var mySlice []int = myArray[9:10]
	//基于myArray的所有元素创建数组切片
	mySlice = myArray[:]
	//基于myArray的前五个元素创建切片
	mySlice = myArray[:5]
	//基于myArray从第五个元素开始的剩余元素创建切片
	mySlice = myArray[5:]

	/*
	直接创建
	 */
	 //创建一个有5个元素的切片，元素初始值为0
	mySlice1 := make([]int, 5)
	//直接创建并初始化元素值的切片
	mySlice2 := []int{1,2,3,4,5}



	for i, v := range mySlice {
		fmt.Printf("mySlice[%d]=%d \n", i, v)
	}
	for i, v := range mySlice1 {
		fmt.Printf("mySlice1[%d]=%d \n", i, v)
	}
	for i, v := range mySlice2 {
		fmt.Printf("mySlice2[%d]=%d \n", i, v)
	}
}

