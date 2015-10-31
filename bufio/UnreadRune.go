package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	rb := bytes.NewBuffer([]byte("123456"))
	r := bufio.NewReader(rb)
	r.ReadByte()
	// error occurs
	fmt.Println(r.UnreadRune())
	c, _, _ := r.ReadRune()
	fmt.Printf("read %s\n", string(c))
	// no error happens
	fmt.Println(r.UnreadRune())
	c, _, _ = r.ReadRune()
	fmt.Printf("read %s\n", string(c))
}
