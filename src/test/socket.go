package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}

	con, err := net.Dial("tcp", os.Args[1]) //获取连接，在go语言中无论建立tcp/udp/icmp等协议都直接调用该方法
	checkError(err)

	_, err = con.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)
	err, result := readFully(con)
	if err != nil {
		fmt.Println(string(result))
	}
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err.Error())
		os.Exit(1)
	}
}

//读取返回的数据信息
func readFully(con net.Conn) (error, []byte) {
	defer con.Close() //defer关键字在函数调用结束后执行该语句

	result := bytes.NewBuffer(nil)
	var buf [512]byte

	for {
		len, err := con.Read(buf[0:])
		result.Write(buf[0:len])

		if err != nil {
			if err == io.EOF { //文本结束标识符
				break
			}
			return err, nil
		}
	}
	return nil, result.Bytes()
}
