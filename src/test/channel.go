package main

import (
	"fmt"
	"sync"
)

//channel中读取/写入都会堵塞等待
//goroutime/channel是消息传递并发模型的一直实现
func Count(ch chan int, num int) {
	fmt.Println("countimg")
	ch <- num //写入时没有读取方，会堵塞
}

/*type Vector []float64

func (v Vector) DoSome(begin, end int, u Vector, c chan int)  {
	for i := begin; i < end; i++ {
	}
}*/
var str string
var once sync.Once

func setup() {
	str = "hello world"
	fmt.Println(str)
}

func printStr(c chan int) {
	once.Do(setup) //全局唯一性操作，setup方法只调用一次，在调用once.Do()方法时其他goroutine会阻塞
	c <- 1
}

func main() {
	/*chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count(chs[i], i)
	}

	for _, ch := range chs{
		value := <- ch
		fmt.Println(value)
	}*/
	/*
		ch := make(chan int, 1) //设置缓存区为1的channel

		for {
			select {        //select机制，case表达式只能是对channel的操作，只要有个操作通过程序就会往下执行
			case ch <- 0:
			case ch <- 1:
			}
			i := <-ch
			fmt.Println(i)
		}*/

	/*ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		go Count(ch, i)        //在缓冲区中写入时，没有读取方也能写入，直到缓冲区填满
	}*/

	/*for value := range ch{
		fmt.Println(value)
	}*/

	/*timeout := make(chan bool, 1)  //超时等待
	go func() {
		time.Sleep(1e9) //等待1秒钟
		timeout <- true
	}()
	ch := make(chan bool)
	select {
		case <-ch:
		case <-timeout:
	}
	close(timeout)   //关闭channel
	_, ok := <-timeout
	fmt.Println(ok)
	fmt.Println("timeout")*/
	/*fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.Gosched)//让出CPU时间片*/
	c := make(chan int, 2)
	go printStr(c)
	go printStr(c)
	for i := 0; i < 2; i++ {
		<-c
	}
}
