package main

import "fmt"

func main() {
	var mySlice = make([]int, 5)
	fmt.Println("len_value", len(mySlice))
	fmt.Println("cap_value:", cap(mySlice))

	//append
	mySlice = append(mySlice,1,2,3)

	mySlice2 := []int{1,1,1,1,1}

	mySlice = append(mySlice,mySlice2...)
	for i, v := range mySlice {
		fmt.Printf("mySlice[%d]=%d \n", i, v)
	}

	fmt.Println("len_value", len(mySlice))
	fmt.Println("cap_value:", cap(mySlice))

	//基于slice创建slice
	oldSlice := mySlice
	newSlice := oldSlice[10:]
	for i,v :=range newSlice  {
		fmt.Printf("newSlice[%d]=%d \n", i, v)
	}

	//内容复制
	slice1 := []int{1,2,3,4,5}
	slice2 := []int{6,7,8}
	copy(slice1,slice2)
	copy(slice2,slice1)
	for i,v :=range slice1  {
		fmt.Printf("slice1[%d]=%d \n", i, v)
	}
	for i,v :=range slice2  {
		fmt.Printf("slice2[%d]=%d \n", i, v)
	}
}
