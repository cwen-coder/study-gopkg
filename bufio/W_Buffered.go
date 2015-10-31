package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	wb := bytes.NewBuffer(nil)
	w := bufio.NewWriter(wb)
	fmt.Println(w.Buffered())
	w.WriteByte('1')
	fmt.Println(w.Buffered())
	w.Flush()
	fmt.Println(w.Buffered())
}
