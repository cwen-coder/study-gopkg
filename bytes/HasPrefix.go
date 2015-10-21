package main

import (
	"bytes"
	"fmt"
)

/*参数列表
s 要检查的字节切片
prefix 前缀字节切片

返回值
bool 是否符合前缀

功能说明
HasPrefix检查s的前缀是否为prefix
*/
func main() {
	s := []byte("abcdef")
	fmt.Println(bytes.HasPrefix(s, []byte("abc")))
	fmt.Println(bytes.HasPrefix(s, []byte("bcd")))
	fmt.Println(bytes.HasPrefix(s, []byte("dabc")))
}
