package main

import (
	"bytes"
	"fmt"
)

/*参数列表

s 字节切片
sep 子字节切片
返回值

int 子字节切片sep的数量
功能说明

Count计算子字节切片sep在字节切片s中非重叠实例的数量*/

func main() {
	s := []byte("12345677777777")
	fmt.Println(bytes.Count(s, []byte("123")))
	fmt.Println(bytes.Count(s, []byte("77")))
	fmt.Println(bytes.Count(s, []byte("777")))
}
