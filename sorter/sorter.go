package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

var inFile *string = flag.String("in", "unsorted.dat", "values fro sort")
var outFile *string = flag.String("out", "sorted.dat", "")

func main() {
	/*	flag.Parse() //解析命令行参数
		var values []int
		if inFile != nil {
			values, _ = readValues(*inFile)
		}

		if outFile != nil {
			writeValues(values, *outFile)
		}
		fmt.Fprint()*/
	fmt.Println("hello")
}

//从文件中读取需要排序的数组
func readValues(inFile string) (values []int, err error) {
	file, err := os.Open(inFile)
	if err != nil {
		fmt.Println("fail to open file", inFile)
		return
	}
	defer file.Close() //是否文件句柄

	br := bufio.NewReader(file)
	values = make([]int, 0)
	for {
		line, isPrefix, err1 := br.ReadLine() //line为当前行的字节数组，isPrefix当前行的长度超过给的buff大小
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}
		if isPrefix {
			fmt.Println("line too long")
			return
		}

		lineStr := string(line)              //将字节数组转换成字符串
		value, err1 := strconv.Atoi(lineStr) //将字符串转换为int
		if err1 != nil {
			err = err1
			return
		}
		values = append(values, value)
	}
	return values, err
}

//将数组写入文件中
func writeValues(values []int, outFile string) error {
	file, err := os.Create(outFile)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, value := range values {
		file.WriteString(strconv.Itoa(value) + "\n")
	}
	return nil
}
