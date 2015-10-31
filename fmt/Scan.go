package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	fmt.Println(a, b, c)
	fmt.Scanf("%d %d %d", &a, &b, &c)
	fmt.Println(a, b, c)
	fmt.Scanln(&a, &b, &c)
	fmt.Println(a, b, c)
}
