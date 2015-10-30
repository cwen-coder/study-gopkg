// func Abs(path string) (string, error)
// 参数列表:
// path:
// 返回值列表:
// 返回所给路径的绝对路径string
// error
// 功能说明：
// 返回所给的路径的的绝对路径。
// 官方文档：
// Abs returns an absolute representation of path.
// If the path is not absolute it will be joined with the current working directory to turn it into an absolute path.
// The absolute path name for a given file is not guaranteed to be unique.
package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	// 当前路径为/home, 如下返回的path将会是/home/abs_demo.go
	path, _ := filepath.Abs("abs_demo.go")
	fmt.Printf("%v\n", path)
}
