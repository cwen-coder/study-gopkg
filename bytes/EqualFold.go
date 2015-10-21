package main

import (
	"bytes"
	"fmt"
)

/*参数列表

s 字节切片
f 过滤函数
返回值

[][]byte 被分割的字节切片的切片
功能说明

FieldsFunc把s解释为UTF-8编码的字符序列，对于每个Unicode字符c，如果f(c)返回true就把c作为分隔字符对s进行拆分。如果所有都字符满足f(c)为true，则返回空的切片。
*/
func main() {
	fmt.Println(bytes.EqualFold([]byte("abc"), []byte("abc")))
	fmt.Println(bytes.EqualFold([]byte("abc"), []byte("abd")))
	fmt.Println(bytes.EqualFold([]byte("abc"), []byte("aBc")))
}
