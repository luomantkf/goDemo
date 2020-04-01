package main

import (
	"fmt"
	"unsafe"
)

var ( //这种方式一般用于声明全局变量
	a int
	b bool
	f bool
)

/*const定义枚举类型*/
const (
	MAN   = 1
	WOMAN = 2
)

/*ioba相当于const表达式中的行标，依次递增*/
const (
	ioba1 = iota
	ioba2 = iota
	test  = 1
	ioba3 = iota
)

var c, d int = 1, 2

func main() {

	var str1, str2 string = "str1", "str2" //声明的变量必须使用，否则会编译报错，全局变量可以不使用
	fmt.Println(str1 + str2)
	var num1, num2 = 1, 2
	fmt.Println(num1, num2)

	/*	var value1 int32   //go语言中int 和int32类型相关，int的所占字节由使用的平台决定, 可进行int32(value2)进行类型强转
		value2 := 1
		value1 = value2
		value1 = int32(value2)*/

	var num int
	fmt.Println(num)

	var var1 = "strVar"
	fmt.Println(len(var1), unsafe.Sizeof(var1))

	//var intVar int
	intVar := 1 //省略var的方式声明变量,只能出现在方法体中，不可用:=方式进行赋值
	fmt.Println(intVar)

	fmt.Println(a, b, c, d, intVar)

	_, var2 := 1, 2 //_空白标识符表示舍弃1变量，常用于go调用方法返回多个值，舍弃不需要的值
	fmt.Println(var2)

	const LENGTH = 1 //常量
	const WIDTH = 2
	area := LENGTH * WIDTH
	fmt.Println(area)

	fmt.Println(MAN)

	fmt.Println(ioba1, ioba2, ioba3)

	if LENGTH > WIDTH {
		fmt.Println("大于")
	} else {
		fmt.Println("小于")
	}

	i, j := 1, 2
	if i == j {
		fmt.Println("i不等于j")
	}

	//遍历字符串
	str := "Hello World"
	for i := 0; i < len(str); i++ {
		fmt.Println(str[i]) //获取到的str[i]字符类型为byte
		//str[i] = 'c' 字符串在初始化后不可再被赋值
	}
	for i, ch := range str {
		fmt.Println(i, ch) //获取到的ch字符类型为rune
	}

	// b := 1 bool类型类型自动/强转转换
	b := (1 == 2)
	fmt.Println(b)

	//&获取变量内存地址，*指针变量,指针变量存的是地址值，只有指针变量前加*方可等到变量
	fmt.Println(&var1)
	fmt.Println(*(&var1))
	var ptr1 *string
	ptr1 = &var1
	ptr := &var1
	fmt.Println(ptr, ptr1)
}
