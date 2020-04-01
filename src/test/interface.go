package main

import "fmt"

//go语言的非侵入式接口，只需实现接口所要求的所有方法，我们就是说这个类实现了接口
//如下File类实现了Read、Write接口方法，所以说File类是IFile、IRead、IWrite的实现类
type IFile interface {
	Read(in string) string
	Write(out string) string
}

type IWrite interface {
	Write(out string) string
}

type IRead interface {
	Read(in string) string
}
type IWriteRead interface { //接口组合
	IWrite
	IRead
}
type File struct {
}

func (file *File) Write(out string) string {
	return out
}

func (file *File) Read(in string) string {
	return in
}

type Less interface {
	Less(b Integer) bool
	add2(b Integer)
}

type Less1 interface {
	Less(b Integer) bool
}
type Int1 struct {
}

func (a *Int1) Less(b Integer) bool {
	return true
}

func main() {
	//var file1 IFile = new(File)
	//var file2 IWrite = new(File)
	//var a Integer = 1    接口赋值会现在第二种方式，因为Integer add2方式使用(a *Integer) add2(b Integer)实现索引传递，第一种为值拷贝，跟用户相违背
	//var b Less = a		接口赋值中只要A是B接口类的子集，即可以将B实现类赋值给A，反之不可
	//var b Less = &a
	//var a Integer = 1   less实现方式是传的(a Integer) Less(b Integer) bool 值对象，可以赋值
	//var b Less1 = a
	var a Less1 = new(Int1)
	if b, ok := a.(Less1); ok { //接口查询，判断a是否是Less1类型
		fmt.Println(b, ok)
	}

	/*if b, ok := a.(*Int1); ok {  //类型查询

	}*/
}
