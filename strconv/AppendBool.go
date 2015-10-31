package main

import (
	"fmt"
	"strconv"
)

/*
参数列表

dst 表示原列表
b 表示需要添加的bool值，true或者false

返回值：
[]byte 返回原列表追加bool后新生成的列表

功能说明：
将布尔值 b 转换为字符串 "true" 或 "false" 然后将结果追加到 dst 的尾部，返回追加后的 []byte
*/
func main() {
	b := []byte("bool:")
	b = strconv.AppendBool(b, true)
	fmt.Println(string(b))
}
