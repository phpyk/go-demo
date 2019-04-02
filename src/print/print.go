package main

import (
	"fmt"
	"log"
)

const Pi = 3.14159261
const mask = 1 << 3
const left = 8 >> 2

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

const (
	yuekai = iota
)
const isWeekend bool = false

func main() {
	var fval = 110.49
	var ival = 200
	var sval = "This is a string"
	log.Println("this is a log ")
	fmt.Println("The value of fval is", fval)
	fmt.Printf("fval = %f, ival = %d, sval = %s \n", fval, ival, sval)
	fmt.Printf("fval = %v, ival = %v, sval = %v \n", fval, ival, sval)

	_, _, nickName := getName()
	fmt.Println("the nickName is", nickName)
	fmt.Println("thie pi is", Pi)
	fmt.Println("thie mask is", mask)
	fmt.Println("thie left is", left)

	fmt.Printf("Sunday = %v, Monday = %v, Tuesday = %v, yuekai = %v \n", Sunday, Monday, Tuesday,yuekai)

	fmt.Println("today is weekend?", !isWeekend)
}

func getName() (firstName, lastName, nickName string) {
	return "yue", "kai", "kk"
}

