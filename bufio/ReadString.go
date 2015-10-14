package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	rb := bytes.NewBuffer([]byte("1234$5678"))
	r := bufio.NewReader(rb)
	line, err := r.ReadString('$')
	if err == nil {
		fmt.Println(line)
	}
}
