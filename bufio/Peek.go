package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	rb := bytes.NewBuffer([]byte("12345678"))
	r := bufio.NewReader(rb)
	b1, _ := r.Peek(4)
	fmt.Println(string(b1))
	b2, _ := r.Peek(8)
	fmt.Println(string(b2))
}
