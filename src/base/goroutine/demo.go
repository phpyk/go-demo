package main

import (
	"fmt"
	"runtime"
	"time"
)

func say(s string) {
	for i := 0; i < 100; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}
func loop() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

func main() {
	go say("world") // 开一个新的 Goroutines 执行
	say("hello") // 当前 Goroutines 执行
	time.Sleep(time.Second)	//让主线程等待一会再退出


	/**
	channel 通过操作符 <- 来接收和发送数据
	ch <- v    // 发送 v 到 channel ch.
	v := <-ch  // 从 ch 中接收数据，并赋值给v
	 */
	a := []int{1, 1, 1, 1, 1, 2, 2, 2, 2, 2,}
	c := make(chan int)
	go sum(a[:len(a)/2],c)
	go sum(a[len(a)/2:],c)
	x,y := <-c,<-c
	fmt.Printf("x:%d,y:%d,x+y:%d \n",x,y,x+y)


	d := make(chan int ,4)
	d <- 1
	d <- 2
	d <- 3
	fmt.Println(<-d)
	fmt.Println(<-d)
	fmt.Println(<-d)
}



func sum(a []int, c chan int) {
	total := 0
	for _,v := range a{
		total += v
	}
	c <- total  // send total to c
}

