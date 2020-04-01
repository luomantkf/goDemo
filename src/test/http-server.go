package main

import (
	"fmt"
	"io"
	"net/http"
)

func helloHandle(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello http")
}

func main() {
	http.HandleFunc("/hello", helloHandle) //注册url映射的处理方法

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
