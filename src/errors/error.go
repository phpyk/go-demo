package main

import (
	"errors"
	"fmt"
)

func main() {
	sum, err := getSum(0, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("sum is :", sum)
	}
}
//方式三：采用fmt.Errof 将string信息转化为error信息 并返回
func getSum3(a, b int) (sum int, err error) {
	sum = a + b
	if a == 0 || b == 0 {
		err = fmt.Errorf("%s", "a和b都不能为0哦")
		return sum, err
	} else if a < 0 || b < 0 {
		err = fmt.Errorf("%s", "a和b都不能为小于哦")
		return sum, err
	} else {
		return sum, nil
	}
}

//方式二：采用errors包的New方法 返回一个err的类型
func getSum2(a, b int) (sum int, err error) {
	sum = a + b
	if a == 0 || b == 0 {
		return sum, errors.New("a和b都不能为0哦")
	} else if a < 0 || b < 0 {
		return sum, errors.New("a和b都不能为小于哦")
	} else {
		return sum, nil
	}
}

//方式一：自定义一个valueError对象，实现Error()方法即可
type valueError struct {
	errMsg string
}
func (te *valueError) Error() string {
	return te.errMsg
}
func getSum(a, b int) (sum int, err error) {
	sum = a + b
	if a == 0 || b == 0 {
		return sum, &valueError{"a和b都不能为0哦",}
	} else if a < 0 || b < 0 {
		return sum, &valueError{"a和b都不能为小于哦"}
	} else {
		return sum, nil
	}
}
