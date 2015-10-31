/*func Pipe() (r *File, w *File, err error)

参数列表
无

返回值：
r 只读文件指针
w 只写文件指针
error 返回错误信息对象

功能说明：
这个函数主要是获取一对文件指针，一个只读（如：stdout)，一个只写（如：stdin）

*/
package main

import (
	"fmt"
	"os"
)

func main() {
	r, w, err := os.Pipe()
	fmt.Printf("Read *Files: %+v, Write *Files: %+v, err: %v\n", r, w, err)
}
