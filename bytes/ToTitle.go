package main

import (
	"bytes"
	"fmt"
)

func main() {
	s := []byte("hello, world!")
	fmt.Println(string(bytes.ToTitle(s)))
	fmt.Println(string(s))
}
