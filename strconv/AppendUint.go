/*
参数列表
dst 原列表
i 需要追加到列表的unit64值
base unit64的进制

返回值:
[]byte 返回列表

功能说明：
类似AppendInt，只能追加uint类型，base表示uint表示的进制数，返回追加后的 []byte
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	b10 := []byte("uint (base 10):")
	b10 = strconv.AppendUint(b10, 42, 10)
	fmt.Println(string(b10))

	b16 := []byte("uint (base 16):")
	b16 = strconv.AppendUint(b16, 42, 16)
	fmt.Println(string(b16))
}
