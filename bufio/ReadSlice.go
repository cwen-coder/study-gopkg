package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	rb := bytes.NewBuffer([]byte("1234$56789"))
	r := bufio.NewReader(rb)
	line, err := r.ReadSlice('$')
	if err != nil {
		return
	}

	fmt.Println(string(line))
	r.ReadSlice('$')
	fmt.Println(string(line))
}
