package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	rb := bytes.NewBuffer([]byte("123\r\n456"))
	r := bufio.NewReader(rb)
	line, prefix, err := r.ReadLine()
	if err == nil {
		fmt.Printf("%v, %s, %v\n", line, string(line), prefix)
	}
}
