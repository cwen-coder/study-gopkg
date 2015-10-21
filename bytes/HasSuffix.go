package main

import (
	"bytes"
	"fmt"
)

/*参数列表

s 要检查的字节切片
suffix 后缀字节切片

返回值
bool 是否有指定后缀

功能说明
HasSuffix检查s是否以suffix为后缀*/

func main() {
	s := []byte("1234567")
	fmt.Println(bytes.HasSuffix(s, []byte("567")))
	fmt.Println(bytes.HasSuffix(s, []byte("456")))
	fmt.Println(bytes.HasSuffix(s, []byte("678")))
}
