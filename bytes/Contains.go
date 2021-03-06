package main

import (
	"bytes"
	"fmt"
)

/*参数列表

b 字节切片
subslice 子字节切片
返回值

bool 是否包含子切片
功能说明

Contains检查字节切片b是否包含子字节切片subslice*/

func main() {
	b := []byte("12345678")
	s1 := []byte("456")
	s2 := []byte("789")
	fmt.Println(bytes.Contains(b, s1))
	fmt.Println(bytes.Contains(b, s2))
}
