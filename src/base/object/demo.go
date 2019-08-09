package main
/**
面向对象--方法
 */
import (
	"fmt"
	"math"
)

type Rectangle struct {
	width,height float64
}
//对象的方法
//定义方式：func (r ReceiverType) funcName(parameters) (results)
func (r Rectangle) area() float64 {
	return r.width * r.height
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	r1 := Rectangle{4, 5,}
	r2 := Rectangle{8, 8,}
	c1 := Circle{3,}
	c2 := Circle{6}

	fmt.Println("r1 的面积为：",r1.area())
	fmt.Println("r2 的面积为：",r2.area())
	fmt.Println("c1 的面积为：",c1.area())
	fmt.Println("c2 的面积为：",c2.area())
}

