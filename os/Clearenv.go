/*
func Clearenv()
参数列表
无

返回值：
无

功能说明：
这个函数主要是清空当前环境变量
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("The path is: %+v\n", os.Environ())
	os.Clearenv()
	fmt.Printf("Now the path is: %+v\n", os.Environ())

}
