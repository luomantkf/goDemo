package main

import (
	"encoding/json"
	"fmt"
)

type Result struct {
	Code    int    `json:"code"` //结构体中使用标签, json:后不能跟空格,对json编解码的字段进行重命名
	Message string `json:"msg"`
}

func main() {
	var result Result
	result.Code = 200
	result.Message = "success"

	res, err := json.Marshal(result)
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(string(res))

	var res2 Result
	if err := json.Unmarshal(res, &res2); err == nil { //Unmarshal解码时第二个参数需要传入接受对象的地址
		fmt.Println(res2.Code)
	}
}
