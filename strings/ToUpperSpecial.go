package main

import (
	"fmt"
	"strings"
	"unicode"
)

// ToUpperSpecial 将 s 中的所有字符修改为其大写格式。
// 该函数把s字符串里面的每个单词转化为大写，但是调用的是unicode.SpecialCase的ToUpper方法
// 优先使用 _case 中的规则进行转换
//func ToUpperSpecial(_case unicode.SpecialCase, s string) string
/*
_case 规则说明，以下列语句为例：
unicode.CaseRange{'A', 'Z', [unicode.MaxCase]rune{3, -3, 0}}
·其中 'A', 'Z' 表示此规则只影响 'A' 到 'Z' 之间的字符。
·其中 [unicode.MaxCase]rune 数组表示：
当使用 ToUpperSpecial 转换时，将字符的 Unicode 编码与第一个元素值（3）相加
当使用 ToLowerSpecial 转换时，将字符的 Unicode 编码与第二个元素值（-3）相加
当使用 ToTitleSpecial 转换时，将字符的 Unicode 编码与第三个元素值（0）相加

*/
func main() {
	var SC unicode.SpecialCase
	fmt.Println(strings.ToUpperSpecial(SC, "Gopher"))
	//GOPHER
}
