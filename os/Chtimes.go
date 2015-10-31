/*
func Chtimes(name string, atime time.Time, mtime time.Time) error
参数列表
name 修改的文件名
atime 访问时间
mtime 修改时间

返回值：
返回error 返回错误信息对象

功能说明：
这个函数主要是修改一个文件的访问时间和修改时间
*/
package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

func main() {
	fi, err := os.Stat("Hello.go")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	fmt.Printf("Hello.go: atime=%+v, mtime=%v\n", fi.Sys().(*syscall.Stat_t).Atim, fi.ModTime())

	err = os.Chtimes("Hello.go", time.Now(), time.Now())
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	fi, err = os.Stat("Hello.go")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("Now Hello.go: atime=%+v, mtime=%v\n", fi.Sys().(*syscall.Stat_t).Atim, fi.ModTime())
}
