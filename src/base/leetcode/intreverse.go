package main

import (
	"fmt"
	"math"
)

func reverse(num int) int {
	ret := 0
	for num != 0 {
		pop := num % 10
		num /= 10
		ret = ret*10 + pop
		//if ret > 1<<31-1 || ret < -1<<31 {
		if ret > math.MaxInt32 || ret < math.MinInt32 {
			return 0
		}
	}
	return ret
}

func main() {
	num := 1234560000
	revNum := reverse(num)
	fmt.Println("revNum:",revNum)
}