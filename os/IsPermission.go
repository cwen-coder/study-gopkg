/*func IsPermission(err error) bool

参数列表
err 错误信息对象

返回值：
bool 是否为无权限错误

功能说明：
这个函数主要是判断一个错误是否为无权限错误

*/
package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%t\n", os.IsPermission(os.ErrPermission))
	fmt.Printf("%t\n", os.IsPermission(errors.New("Custom Error")))
}
