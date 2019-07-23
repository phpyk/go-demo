package main

import (
	"fmt"
)

func main() {
	//无限循环
	sum := 0
	for {
		sum ++
		fmt.Println("sum:", sum)
		if sum >= 100 {
			break
		}
	}
	//for循环多重赋值,将数组倒序排列
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10,}
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i] ,a[j] = a[j],a[i]
	}
	fmt.Printf("a:%v\n",a)

	//定义循环标签
	JLoop:
	for j := 0; j < 5; j++ {
		for i := 0; i < 10; i++ {
			if i>5{
				break JLoop //此时中断的是外层循环
			}
			fmt.Println(i)
		}
	}

	//goto跳转语句
	myFunc()

}

func myFunc() {
	i := 0
	HERE:
		fmt.Println("i=",i)
	i++
	if i< 10 {
		goto HERE
	}
}
