package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.IndexFunc("astaxie", splitfunc)) //符合条件的是字符x，因为字符x的rune大于t，所以这个位置应该返回4
	fmt.Println(strings.IndexFunc("aaabbbb", splitfunc)) //所有的字符都不符合条件，则返回-1
}

func splitfunc(a rune) bool {
	if a > 't' {
		return true
	}
	return false
}
