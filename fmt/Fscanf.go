package main

import "fmt"
import "os"

var (
	n, a, b, c int
	err        error
)

func main() {

	n, err = fmt.Fscan(os.Stdin, &a, &b, &c)
	fmt.Println(a, b, c)
	fmt.Printf("输入正确参数%v个，错误参数原因:%v", n, err)
	fmt.Fscanf(os.Stdin, "%d %d %d", &a, &b, &c)
	fmt.Println(a, b, c)
	fmt.Fscanln(os.Stdin, &a, &b, &c)
	fmt.Println(a, b, c)
}
