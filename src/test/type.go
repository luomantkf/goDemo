package main

import "fmt"

type Integer int

//为类型Integer提供内置函数,即使是基本类型也能提供内置函数
func (a Integer) Less(b Integer) bool {
	return a > b
}

func (a Integer) add1(b Integer) { //this对象值传递
	a += b
}

func (a *Integer) add2(b Integer) { //(a *Integer)类似于java中this指针,对象引用传递, go语言中没有隐形指针this的概率
	*a += b
}

//大小写定义成员变量的访问权限，小写表示只能包内访问，但是在go语言中是包一级而不是类型，所以说包中的其他类型也能访问其内部变量
type Rect struct {
	x, Y int
}

func (reat *Rect) getX() int {
	return reat.x
}
func (reat *Rect) area() int {
	return reat.x + reat.Y
}

type Style struct {
	id   int
	name string
}

func main() {
	var a Integer = 1
	fmt.Println(a.Less(2))
	a.add1(1)
	fmt.Println(a)
	a.add2(2)
	fmt.Println(a)
	//初始方式，默认会初始化该类型的零值,go语言中没有构造函数的概念
	//style := new(Style)
	//style := &Style{}
	//style := &Style{1, "xx"}
	style := &Style{id: 1, name: "xxx"}
	fmt.Println(style.id)
}
