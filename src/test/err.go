package main

import (
	"io"
	"os"
)

type FileError struct {
	Op  string
	err error
}

func (e *FileError) Error() string { //为FileError类型提供内置函数，Error()代表把FileError当作error来传递
	return e.Op + " " + e.err.Error()
}

/*func returnErr() error {
	return &FileError{"err", error()}
}*/

func CopyFile(src, dest string) (w int64, err error) { //defer关键字相当于try/catch/finally中的finally,在函数结束后无论是否抛出异常都会按照先进后出的顺序关闭资源
	srcFile, err := os.Open(src)
	if err != nil {
		return
	}
	defer srcFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return
	}
	defer destFile.Close()

	/*	defer func() {          //定义匿名函数处理
		srcFile.Close()
		destFile.Close()
	}()*/
	/*defer func(){
		if r := recover(); r != nil {  //recover()函数用于终止panic({}interface)函数提起的错误处理流程
			fmt.Println("runtime err", r)
		}
	}()*/

	return io.Copy(srcFile, destFile)
}
