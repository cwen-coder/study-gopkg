// func EvalSymlinks(path string) (string, error)
// 参数列表:
// path: 路径
// 返回值列表:
// 实际的路径
// errors
// 功能说明: 返回一个链接文件的实际的路径。例如在/home下创建一个link.log的文件，然后cd xuchdong; ln -s /home/link.log link， 使用该函数可以找到原始的文件即/home/link.log
// 官方文档：
// EvalSymlinks returns the path name after the evaluation of any symbolic links.
// If path is relative the result will be relative to the current directory,
// unless one of the components is an absolute symbolic link.
package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	// 环境准备：首先在/home目录下创建一个link.log的文件,
	// 然后在当前目录下使用ln -s /home/link.log link_other

	path, _ := filepath.EvalSymlinks("link_other")
	fmt.Printf("%s\n", path) // /home/link.log
}
