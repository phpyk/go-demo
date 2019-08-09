package main

import "fmt"

func fibonacci(n int, c chan int) {
	x,y := 1,1
	for i := 0; i < n; i++ {
		c <- x
		x,y = y,x+y
	}
	close(c)
}

//在多个channel之间选择
func fibonacci2(c,quit chan int)  {
	x,y := 1,1
	for {
		select {
		case c <- x:
			x,y = y,x+y
		case <-quit:	//当quit channel中可以取到值的时候退出
			return
		}
	}
}

func main() {
	//指定channel的缓冲区大小为20
	fmt.Println("fibonacci:")
	c := make(chan int , 20)
	go fibonacci(cap(c), c)
	for i := range c{
		fmt.Print(i," ")
	}

	fmt.Println()
	fmt.Println("fibonacci2:")
	c2 := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i<15 ; i++ {
			fmt.Print(<-c2," ")
		}
		//发送0到quit channel
		quit <- 0
	}()

	fibonacci2(c2,quit)
}