package main

import "fmt"

func example1(x int) int {
	if x = 1; x == 0 { //可以在条件语句前进行变量赋值
		return 0
	} else {
		return 1
	}
}

func example2(x int) {
	switch x {
	case 1:
		fmt.Println("1")
	case 2:
		fallthrough //跳过执行下一个case
	case 3:
		fmt.Println("3")
	}

	switch { //可以不设定条件语句,相当于多个if..case..语句嵌套
	case x >= 4:
		fmt.Println("大于4")
	}
}

func example3() {
	for i, j := 0, 10; i < j; i, j = i+1, j-1 { //for循环支持多重赋值，go语言中无while,do-while

	}
	sum := 0
Jloop:
	for { //无限循环
		sum++
		if sum > 100 {
			break Jloop //for同样支持continue与break来控制循环,同时支持break特定的外循环
		}
	}
}

func example4() {
HEAD: //goto语句操作不当容易造成死循环
	fmt.Println("goto")
	fmt.Println("test")
	goto HEAD
}

func main() {
	fmt.Println(example1(1))
	example2(2)
	example4()
}
