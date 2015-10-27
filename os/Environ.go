/*func Environ() []string

参数列表
无

返回值：
[]string 环境变量

功能说明：
这个函数主要是获取当前环境变量*/

package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("The path is: %+v\n", os.Environ())
}
