/*
func Hostname() (name string, err error)

参数列表
无

返回值：
name hostname
err 返回错误信息对象

功能说明：
这个函数主要是返回机器的hostname
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	name, err := os.Hostname()
	if err != nil {
		fmt.Printf("Error: %v\n", name)
		return
	}

	fmt.Printf("The Hostname is : %s\n", name)
}
