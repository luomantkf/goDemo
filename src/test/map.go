package main

import "fmt"

type Person struct {
	ID   string
	NAME string
}

func main() {
	var personDB map[string]Person
	personDB = make(map[string]Person, 100) //string指定key的类型，Person指定value的类型，内置函数make()创建map,也可直接指定容器大小
	personDB["k1"] = Person{"id1", "name1"} //给map赋值
	person, ok := personDB["k1"]            //获取map中的值，ok表示是否包含该值
	if ok == true {
		fmt.Println(person)
	}
	delete(personDB, "k1") //内置函数删除map中的某个值
}
