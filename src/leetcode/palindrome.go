package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	arr := []int{}
	for x != 0 {
		y := x % 10
		arr = append(arr,y)
		x /= 10
	}
	l := len(arr)
	half := l/2
	if l%2 ==1 {
		half = l/2 + 1
	}
	for i := 0; i < half; i++ {
		if arr[i] != arr[len(arr)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	res := isPalindrome(12321)
	fmt.Printf("res:%v",res)
}