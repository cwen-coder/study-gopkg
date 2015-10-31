/*
func FormatFloat(f float64, fmt byte, prec, bitSize int) string

参数列表
f 需要被转换的float64值
fmt 转换格式 'b' 'e' 'E' 'f' 'g'或'G'
prec 浮点数精度
bitSize 32或64，32对应float32，64对应float64

返回值：
返回string 表示转换后的字符串

功能说明：
将浮点数转换为字符串
浮点数格式有'b' (-ddddp±ddd, 二进制指数), 'e' (-d.dddde±dd, 十进制指数), 'E' (-d.ddddE±dd, 十进制指数), 'f' (-ddd.dddd, 无指数), 'g' (大指数时相当于'e', 其他情况时相当于'f'), 'G' (大指数时相当于'E', 其他情况相当于'f').
精度用于控制当格式为'e' 'E' 'f' 'g' 'G'时除指数外的数字的个数；对于'e' 'E' 'f'指小数点后位数；对于'g' 'G'则表示总共的位数；如果使用-1，表示不改变数值的最小位数
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(strconv.FormatFloat(10.1, 'b', 5, 64))
	fmt.Println(strconv.FormatFloat(10.1, 'e', 5, 64))
	fmt.Println(strconv.FormatFloat(10.1, 'E', 5, 64))
	fmt.Println(strconv.FormatFloat(10.1, 'f', 5, 64))
	fmt.Println(strconv.FormatFloat(10.1, 'g', 5, 64))
	fmt.Println(strconv.FormatFloat(100000000.1, 'g', 5, 64))
	fmt.Println(strconv.FormatFloat(10.1, 'G', 5, 64))
	fmt.Println(strconv.FormatFloat(100000000.1, 'G', 5, 64))

	fmt.Println(strconv.FormatFloat(10.1, 'e', -1, 64))
	fmt.Println(strconv.FormatFloat(10.00001, 'e', -1, 64))

	v := 3.1415926535

	s32 := strconv.FormatFloat(v, 'E', -1, 32)
	fmt.Printf("%T, %v\n", s32, s32)

	s64 := strconv.FormatFloat(v, 'E', -1, 64)
	fmt.Printf("%T, %v\n", s64, s64)
}
