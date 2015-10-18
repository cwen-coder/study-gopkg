package main

import (
	"fmt"
	"strings"
)

// TrimSuffix 删除 s 尾部的 suffix 字符串
// 如果 s 不是以 suffix 结尾，则返回原始 s
//func TrimSuffix(s, suffix string) string

func main() {
	s := "Hello 世界!!!!!"
	ts := strings.TrimSuffix(s, "!!!!")
	fmt.Printf("%q\n", ts) // " 世界"
}
