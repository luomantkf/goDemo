package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		return
	}

	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body) //将获取的数据打印到控制台
}
