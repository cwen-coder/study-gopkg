/*
func IsPrint(r rune) bool

参数列表
r 表示rune字符

返回值：
返回bool true表示可以打印，false表示不可以打印。

功能说明：
判断rune字符在golang中是否被定义为可打印，与unicode.IsPrint相同。可打印范围包括字符、数字、标点、符号以及ASCII中的空格。
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	c := strconv.IsPrint('\u263a')
	fmt.Println(c)

	bel := strconv.IsPrint('\007')
	fmt.Println(bel)

	fmt.Println(strconv.IsPrint(' '))
	fmt.Println(strconv.IsPrint('\t'))
	fmt.Println(strconv.IsPrint('\n'))
	fmt.Println(strconv.IsPrint('\r'))
}
