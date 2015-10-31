package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	wb := bytes.NewBuffer(nil)
	w := bufio.NewWriter(wb)
	n, err := w.WriteString("123456")
	if err != nil {
		return
	}
	fmt.Println(n)
}
