package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	rb := bytes.NewBuffer([]byte("12345678"))
	r := bufio.NewReader(rb)
	fmt.Println(r.Buffered())
	var buf [4]byte
	r.Read(buf[:])
	fmt.Println(r.Buffered())
	r.Read(buf[:])
	fmt.Println(r.Buffered())
}
