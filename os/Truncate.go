/*func Truncate(name string, size int64) error

参数列表
name 文件名
size 目标长度，单位为byte

返回值：
返回error 返回错误信息对象

功能说明：
这个函数主要是截短一个文件到新的长度

*/

package main

import (
	"fmt"
	"os"
)

func main() {
	fi, err := os.Stat("Hello.go")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("The Hello.go's size is: %d\n", fi.Size())

	if err = os.Truncate("Hello.go", 10); err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fi, err = os.Stat("Hello.go")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("Now the Hello.go's size is: %d\n", fi.Size())
}
