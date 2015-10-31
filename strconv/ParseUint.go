/*
func ParseUint(s string, base int, bitSize int) (n uint64, err error)

参数列表
str 可以表示int64值的字符串
base 进制 (2 to 36)
bitSize 精度 0、8、16、32、64对应uint、uint8、uint16、uint32、uint64

返回值：
i 通过str转换的uint64值.
err 当str无法转换为uinit64值返回错误，否则为nil.

功能说明：
ParseUint is like ParseInt but for unsigned numbers.
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	v := "42"
	if s, err := strconv.ParseUint(v, 10, 32); err == nil {
		fmt.Printf("%T,%v\n", s, s)
	}
	if s, err := strconv.ParseUint(v, 10, 64); err == nil {
		fmt.Printf("%T,%v\n", s, s)
	}
}
