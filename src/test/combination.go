package main

import "fmt"

type Base struct {
	name string
}

func (base *Base) getName() string {
	return base.name
}

//非匿名组合
type Foo struct {
	base Base
	name string
}

//匿名组合
type Bar struct {
	Base
	name string
}

func main() {
	base := &Base{"base"}
	foo := &Foo{*base, "foo"}
	fmt.Println(foo.base.getName())
	fmt.Println(foo.name)
	//匿名组合类中bar.getName()相当于bar.base.getName()，跟非匿名类foo.base.getName()一样，最终执行都是执行基类的方法
	bar := &Bar{*base, "bar"}
	fmt.Println(bar.getName())
	fmt.Println(bar.name) //name名字冲突，通过bar.name指向的还是外层类，基类无法访问外层类的成员变量
}
