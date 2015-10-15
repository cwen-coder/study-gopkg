package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	wb := bytes.NewBuffer(nil)
	w := bufio.NewWriter(wb)
	n, err := w.Write([]byte("123456"))
	if err != nil {
		return
	}
	fmt.Println(n)
	w.Flush()
	fmt.Println(string(wb.Bytes()))
}
