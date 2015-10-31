package main

import (
	"fmt"
	"strings"
)

func main() {
	reader := strings.NewReader("Go语言学习园地")
	p := make([]byte, 6)
	n, err := reader.ReadAt(p, 2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s, %d", p, n)
}
