/*
func AppendQuoteRune(dst []byte, r rune) []byte

参数列表
dst 原列表
r 需要append到列表的字符

返回值:
[]byte 返回列表

功能说明：
将符文s转换为单引号引起来的字符串，并将结果追加到dst的尾部，返回追加后的[]byte。其中的特殊字符将被转换为转义字符

*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	newlist := strconv.AppendQuoteRune(make([]byte, 0), 'a')
	fmt.Println(string(newlist))
	newlist = strconv.AppendQuoteRune(make([]byte, 0), '\'')
	fmt.Println(string(newlist))
	newlist = strconv.AppendQuoteRune(make([]byte, 0), '中')
	fmt.Println(string(newlist))
}
