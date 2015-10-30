// func Dir(path string) string
// 参数列表
// path 表示路径字符串
// 返回值：
// 返回string
// 功能说明：
// 这个函数主要是返回path中最后一个元素的路径
// 规则：
// 1.通常是路径最后一个元素的路径目录
// 2.路径为空返回.
// 官方文档：
// Dir returns all but the last element of path, typically the path's directory.
// After dropping the final element using Split, the path is Cleaned and trailing slashes are removed.
// If the path is empty, Dir returns ".". If the path consists entirely of slashes followed by non-slash bytes,
// Dir returns a single slash. In any other case, the returned path does not end in a slash.
package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Dir("/a/b/c")) // /a/b
	fmt.Println(path.Dir(""))       // .
}
