package main

import "fmt"

type person struct {
	name string
	age int
	gender string
}

func main() {
	var P person
	P.name = "岳凯"
	P.age = 31
	P.gender = "男"
	fmt.Printf("person:%+v \n",P) //%+v 打印json形式的结构体

	P2 := person{"贝巧玲",30,"女"}
	fmt.Printf("person2:%+v \n",P2) //%+v 打印json形式的结构体

	older_person,age_diff := Older(P, P2)
	fmt.Printf("更老的是：%v，老%v岁 \n",older_person.name,age_diff)
}

func Older(p1, p2 person) (person, int) {
	if p1.age > p2.age {
		return p1,(p1.age - p2.age)
	}else {
		return p2,(p2.age-p1.age)
	}
}

