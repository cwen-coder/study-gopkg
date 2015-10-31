package main

import (
	"fmt"
	"strings"
)

/*参数列表

s 表示需要处理的字符串
cutset 表示需要过滤的字符集
返回值：

返回string 转化之后的字符串
功能说明：

该函数把s字符串结尾里面包含字符集的字符全部过滤掉，返回过滤之后的字符串

代码实例：*/

func main() {
	fmt.Printf("[%q]", strings.TrimRight(" !!! Achtung !!! ", "! "))
	//[" !!! Achtung"]
}
