package main

import (
	"bytes"
	"fmt"
)

func main() {
	s := []byte("12345677777777")
	fmt.Println(bytes.Count(s, []byte("123")))
	fmt.Println(bytes.Count(s, []byte("77")))
	fmt.Println(bytes.Count(s, []byte("777")))
}
