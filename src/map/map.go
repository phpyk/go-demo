package main

import "fmt"

type PersonInfo struct {
	ID string
	Name string
	Address string
}
func main() {
	//map声明，personDB是map的变量名，string是key类型，PersonInfo是value类型
	var personDB map[string] PersonInfo
	//使用make函数创建map
	personDB = make(map[string]PersonInfo)
	//创建map并指定存贮能力
	personDB = make(map[string]PersonInfo,10)
	//创建并初始化map
	personDB = map[string]PersonInfo{
		"1": {"111", "tom", "Room1",},
		"2": {"222", "jerry", "Room2",},
	}
	//元素赋值
	personDB["111"] = PersonInfo{"111", "Tom", "Room203",}
	personDB["222"] = PersonInfo{"222", "Jerry", "Room105",}
	personDB["333"] = PersonInfo{"222", "Jerry", "Room105",}
	//元素删除
	delete(personDB,"222")
	//元素查找
	person,ok := personDB["111"]
	if ok {
		fmt.Println("found person:",person.Name, "with id:111")
	} else {
		fmt.Println("not found person with id:111")
	}

}
