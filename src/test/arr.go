package main

import "fmt"

func main() {
	//go语言中arr传递时值拷贝
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	modify(arr)
	fmt.Println(arr)

	//数组切片基于数组
	test1()

}

func modify(arr [5]int) {
	arr[0] = 10
	fmt.Println("modified arr:", arr)
}

func test1() {
	arr1 := [4]int{11, 12, 13, 14}
	arr2 := arr1[:3] //对前两个元素进行切片
	// arr3 := arr1[2:] 从第二个元素开始切片
	for index, value := range arr2 {
		fmt.Println(index, value)
	}

	arr4 := make([]int, 2)
	arr3 := make([]int, 2, 4)         //初始化含有两个元素值为0的数组，并预留4个元素空间
	fmt.Println(len(arr3), cap(arr3)) //len()获取元素总数， cap()分配空间总大小
	arr3 = append(arr3, 1, 2)
	arr3 = append(arr3, arr4...)
	for _, v := range arr3 {
		fmt.Println(v)
	}
}
