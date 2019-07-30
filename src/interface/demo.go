package main

import (
	"fmt"
	"reflect"
)

type Human struct {
	name string
	age int
	phone string
}
type Student struct {
	Human
	school string
	loan float64
}
type Employee struct {
	Human
	company string
	money float64
}
func (h Human) SayHi()  {
	fmt.Printf("Hi,I am %s,I am a %s you can call me %s \n",h.name,reflect.TypeOf(h) ,h.phone)
}
func (h Human) Sing(lyrics string) {
	fmt.Println("La la la ,la la la,la la ...",lyrics)
}
func (h Human) Guzzle(beerStein string)  {
	fmt.Println("Guzzle Guzzle Guzzle...",beerStein)
}
func (e Employee) SayHi()  {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone) // 此句可以分成多行
}
func (s Student) BorrowMoney(amount float64) {
	s.loan += amount
}
func (e Employee) SpendSalary(amount float64)  {
	e.money -= amount
}
type Men interface {
	SayHi()
	Sing(lyrics string)
}
//type YoungChap interface {
//	SayHi()
//	Sing(song string)
//	BorrowMoney(amount float64)
//}
//type ElderlyGent interface {
//	SayHi()
//	Sing(song string)
//	SpendSalary(amount float64)
//}

func main() {
	mike := Student{Human{"麦克", 25, "18811112222"}, "清华大学", 0.00,}
	paul := Student{Human{"保罗", 32, "19988881111",}, "北京大学", 100.00,}
	sam  := Employee{Human{"山姆", 28, "17788889999",}, "阿里巴巴", 35000,}
	tom  := Employee{Human{"汤姆", 30, "15500006666",}, "腾讯体育", 25000,}

	var i Men
	i = mike
	fmt.Println("我是麦克，我是一个学生：")
	i.SayHi()
	i.Sing("五环之歌")

	i = tom
	fmt.Println("我是汤姆，我是一个员工")
	i.SayHi()
	i.Sing("咱们工人有力量")

	list := make([]Men,3)
	list[0],list[1],list[2] = paul,sam,tom
	for _,m := range list {
		m.SayHi()
	}

}
