/*返回值

string 字符串
功能说明

String把Buffer中的未读数据作为string返回。如果Buffer是个nil指针，则返回
代码示例*/

package main

import (
	"bytes"
	"fmt"
)

func main() {
	var b *bytes.Buffer
	fmt.Println(b.String())
	b = bytes.NewBuffer([]byte("123"))
	fmt.Println(b.String())
}
