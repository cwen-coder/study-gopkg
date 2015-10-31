package main

import (
	"fmt"
)

func main() {
	str1 := fmt.Sprint("默认样式打印字符串")
	fmt.Println(str1)
	str2 := fmt.Sprintf("Format:%s\n", "格式打印出字符串!")
	fmt.Print(str2)
	str3 := fmt.Sprintln("默认格式打印出字符串并带换行!")
	fmt.Printf(str3)
	result1 := fmt.Sprintln("studygolang.com", 2013)
	result2 := fmt.Sprint("studygolang.com", 2013)
	fmt.Println(result1)
	fmt.Println(result2)
}
