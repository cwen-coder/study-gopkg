package main

import (
	"fmt"
	"strings"
)

func main() {
	a := strings.FieldsFunc("astaxie", splitfunc)
	fmt.Println(a) //输出：[asta ie]
	//如果把下面的函数t字符改成，那么将返回空slice
}

func splitfunc(a rune) bool {
	if a > 't' {
		return true
	}
	return false
}
