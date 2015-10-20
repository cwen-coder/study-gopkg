package main

import (
	"bytes"
	"fmt"
)

func main() {
	b := []byte("12345678")
	s1 := []byte("456")
	s2 := []byte("789")
	fmt.Println(bytes.Contains(b, s1))
	fmt.Println(bytes.Contains(b, s2))
}
