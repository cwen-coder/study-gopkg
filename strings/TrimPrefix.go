package main

import (
	"fmt"
	"strings"
)

// TrimPrefix 删除 s 头部的 prefix 字符串
// 如果 s 不是以 prefix 开头，则返回原始 s
//func TrimPrefix(s, prefix string) string

func main() {
	s := "Hello 世界!"
	ts := strings.TrimPrefix(s, "Hello")
	fmt.Printf("%q\n", ts) // " 世界"
}
