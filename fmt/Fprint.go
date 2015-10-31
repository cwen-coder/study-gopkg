package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprint(os.Stdout, "默认样式\n")
	fmt.Fprintf(os.Stdout, "Format:%s\n", "格式打印!")
	fmt.Fprintln(os.Stdout, "默认格式加换行打印!")
}
