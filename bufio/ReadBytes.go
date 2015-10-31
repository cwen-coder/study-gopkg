package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	rb := bytes.NewBuffer([]byte("123$456"))
	r := bufio.NewReader(rb)
	b, err := r.ReadBytes('$')
	fmt.Printf("%s, %v\n", string(b), err)
}
