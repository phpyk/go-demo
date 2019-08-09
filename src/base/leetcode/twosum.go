package main

import (
	"fmt"
	"time"
)

func twoSum(nums []int,target int) []int {
	for idx0,val0 := range nums {
		for idx,val := range nums {
			if idx == idx0 {
				continue
			}
			if val0 + val == target {
				return []int{idx0,idx}
			}
		}
	}
	return nil
}

func twoSum2(nums []int, target int) []int {
	len := len(nums)
	for i := 0;i<= len-1 ;i++ {
		for j := i+1; j <= len-1; j++ {
			if nums[i] + nums[j] == target {
				return []int{i,j}
			}
		}
	}
	return nil
}


func twoSum3(nums []int, target int) []int {
	mp := make(map[int]int, len(nums))
	for k, v := range nums {
		remain := target - v
		if kk, ok := mp[remain]; ok {
			if k > kk {
				return []int{kk, k}
			}
			return []int{k, kk}
		}
		if _, ok := mp[v]; !ok {
			mp[v] = k
		}
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 22
	time1 := time.Now()
	res := twoSum(nums,target)
	t := time.Since(time1)
	fmt.Println("time:",t)
	fmt.Printf("res:%v",res)
}