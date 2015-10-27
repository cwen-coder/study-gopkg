/*
unc FormatInt(i int64, base int) string

参数列表
i 表示需要被转换为字符串的int64值
base 表示进制数 2 <= base <= 36

返回值：
返回string 表示转换后的字符串，当进制大于10时，大于10的值将使用小写a-z表示。

功能说明：
将数值按照进制格式转换为字符串

*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(strconv.FormatInt(9223372036854775807, 2))
	fmt.Println(strconv.FormatInt(9223372036854775807, 10))

	v := int64(-42)

	s10 := strconv.FormatInt(v, 10)
	fmt.Printf("%T, %v\n", s10, s10)

	s16 := strconv.FormatInt(v, 16)
	fmt.Printf("%T, %v\n", s16, s16)
}
