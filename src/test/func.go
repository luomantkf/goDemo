package main

import "fmt"

func add(a, b int) (c int, d string) { //函数支持多返回值, c、d为对返回值的命令，可有可无
	return 1, "c"
}

func Add(a int) int { //函数名的大小代表函数的调用权限，函数名小写只能包内调用
	return 1
}

func func1(args ...int) { //不定参数，只能出现在参数中且只能为最后一个参数相当于一个数组切片，并可以进行类型的传递

}

func prinf(args ...interface{}) { //任意类型的不定参数
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "int")
		case string:
			fmt.Println(arg, "string")
		}
	}
}

func main() {
	Add(1)
	prinf(1, "2")

	f := func(a int) int { //匿名函数赋值给变量
		return a
	}
	fmt.Println(f(1))

	func(a int) {
		fmt.Println(a)
	}(1) //匿名函数直接执行,花括号后跟传参表示直接调用函数

	i := 1
	func1 := func() func() {
		j := 1
		return func() {
			fmt.Println(i, j)
		}
	}() //使用匿名函数，变量j指向的闭包函数中，只有函数内部才可以获取到变量j,保证了变量j的安全性，闭包函数的作用可以存储变量传递到其他函数中
	func1()
	i++
	func1()
}
