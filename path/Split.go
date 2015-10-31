// func Split(path string) (dir, file string)

// 参数列表

// path 表示路径字符串
// 返回值：

// 返回string,string
// 功能说明：

// 这个函数主要是分离路径中的文件目录和文件
//官方文档：
//Split splits path immediately following the final slash, separating it into a directory and file name component.
//If there is no slash path, Split returns an empty dir and file set to path. The returned values have the property that path = dir+file.

package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Split("static/myfile.css")) // static/ myfile.css
	fmt.Println(path.Split("static"))            // "" static
}
