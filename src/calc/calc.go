package main

import (
	"os"
	"fmt"
	"simplemath"
	"strconv"
)

var Usage = func() {
	fmt.Println("USAGE: calc command [arguments] ...")
	fmt.Println("\nThe commands are: \n\tadd\t两个参数相加 \n\tsqrt\t开根号")
}

func main() {
	args := os.Args

	if args == nil || len(args) < 4 {
		Usage()
		return
	}

	switch args[1] {
		case "add":
			if len(args) != 4 {
				fmt.Println("USAGE: calc add <integer1> <integer2>")
				return
			}
			v1, err1 := strconv.Atoi(args[2])
			v2, err2 := strconv.Atoi(args[3])
			if err1 != nil || err2 != nil {
				fmt.Println("USAGE: calc add <integer1> <integer2>")
				return
			}
			ret := simplemath.Add(v1, v2)
			fmt.Println("Result: ", ret)
		case "sqrt":


		default:
			Usage()
	}
}
