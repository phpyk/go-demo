package main

import "fmt"

func main() {
	var myArray [10]int = [10]int{1,2,3,4,5,6,7,8,9,0}
	
	//元素遍历一
	for i := 0; i < len(myArray); i++ {
		fmt.Printf("myArray[%d]=%d \n",i ,myArray[i])
	}
	//元素遍历二
	for i, v := range myArray {
		fmt.Printf("myArray[%d]=%d \n",i ,v)
	}
}
