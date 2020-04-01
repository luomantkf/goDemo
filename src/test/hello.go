package main /*每一个go程序都需要有一个main包，文件名与包名无直接关系，但一个文件夹下只能有一个包名，否则会报错*/

import "fmt"

//main方法是程序的入口(如何有init()与先执行)
func main() {
	fmt.Println("Hello, World!") //go语言中一行代表一个语句
	fmt.Println("hello2")
}
func init() {
	fmt.Println("init")
}
