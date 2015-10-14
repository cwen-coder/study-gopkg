package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	rb := bytes.NewBuffer([]byte("123456789"))
	r := bufio.NewReader(rb)
	b, err := r.ReadByte()
	fmt.Printf("%c, %v\n", b, err)
}
