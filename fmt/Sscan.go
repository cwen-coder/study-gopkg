package main

import "fmt"

func main() {
	str := "34  343  245"
	var a, b, c int
	fmt.Sscan(str, &a, &b, &c)
	fmt.Println(a, b, c)
	fmt.Sscanf(str, "%d %d %d", &a, &b, &c)
	fmt.Println(a, b, c)
	fmt.Sscanln(str, &a, &b, &c)
	fmt.Println(a, b, c)
}
