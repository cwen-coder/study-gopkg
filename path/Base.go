// func Base(path string) string
// 参数列表
// path 表示需要取路径的字符串
// 返回值：
// 返回string
// 功能说明：
// 这个函数主要是用来返回最后一个元素的路径,如果路径为空返回.如果路径由斜线组成,返回/
// 官方文档：
// Base returns the last element of path. Trailing slashes are removed before extracting the last element.
// If the path is empty,
// Base returns ".". If the path consists entirely of slashes, Base returns "/".
package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Base("/a/b"))
	fmt.Println(path.Base(""))
	fmt.Println(path.Base("//"))
}
