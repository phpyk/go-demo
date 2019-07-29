package main

import "fmt"

type Skills []string

type Person struct {
	name string
	age int
	weight int
	gender string
	money int
}

type Student struct {
	Person	// 匿名字段，那么默认 Student 就包含了 Person 的所有字段
	school string
	Skills
	money int
}

func main() {
	s1 := Student{Person{"岳凯",30,120,"男",0},"富力十号家里蹲大学",[]string{"种地","写代码"},20000}
	fmt.Printf("detail:%+v \n",s1)
	s1.school = "岳戈庄坡里干大学"
	fmt.Printf("detail:%+v \n",s1)
	s1.name = "岳志强"
	fmt.Printf("detail:%+v \n",s1)
	s1.Person = Person{"贝巧玲", 30, 120, "女",0}
	s1.Person.age -= 1
	s1.Skills = []string{"做饭","洗衣服","耍赖皮"}
	s1.Person.money = 10000
	s1.money = 500000
	fmt.Printf("detail:%+v \n",s1)
}

