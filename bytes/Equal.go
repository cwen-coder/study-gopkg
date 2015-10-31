package main

import (
	"bytes"
	"fmt"
)

/*
参数列表
a, b 要比较的字节切片
返回值
bool 是否相同
功能说明
比较两个字节切片是否相等，如果参数为nil，则等同于空的字节切片*/
func main() {
	fmt.Println(bytes.Equal([]byte{1, 2, 3}, []byte{1, 2, 3}))
	fmt.Println(bytes.Equal([]byte{1, 2, 3}, []byte{1, 2}))
	fmt.Println(bytes.Equal([]byte{1, 2, 3}, nil))
	fmt.Println(bytes.Equal([]byte{}, nil))
	fmt.Println(bytes.Equal(nil, nil))
}
