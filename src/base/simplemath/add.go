package simplemath

import "errors"

func Add(a int, b int) int {
	return a + b
}

func IntAdd(a int, b int) (ret int, err error) {
	if a < 0 || b < 0 {
		err = errors.New("只支持非负整数相加")
		return
	}
	return a + b, nil
}
