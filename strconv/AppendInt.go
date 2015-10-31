/*
func AppendInt(dst []byte, i int64, base int) []byte

参数列表
dst 表示原列表
i 表示需要添加的int64值
base 表示进制数 2 <= base <= 36

返回值：
返回[]byte 表示原列表添加数值后新生成的列表

功能说明：
类似AppendFloat，只能追加int类型，base表示int表示的进制数，返回追加后的 []byte。当进制大于10时，大于10的值将使用小写a-z表示。

*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	b10 := []byte("int (base 10):")
	b10 = strconv.AppendInt(b10, -42, 10)
	fmt.Println(string(b10))

	b16 := []byte("int (base 16):")
	b16 = strconv.AppendInt(b16, -42, 16)
	fmt.Println(string(b16))
}
