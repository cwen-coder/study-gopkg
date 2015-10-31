package main

import (
	"fmt"
	"strings"
	"unicode"
)

//该函数把s字符串里面的每个单词转化为标题体，但是调用的是unicode.SpecialCase的ToTitle方法
func main() {
	var SC unicode.SpecialCase
	fmt.Println(strings.ToTitleSpecial(SC, "Gopher"))
	//GOPHER
}
