package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	wb := bytes.NewBuffer(nil)
	w := bufio.NewWriterSize(wb, 8192)
	w.Write([]byte("hello,"))
	w.Write([]byte("world!"))
	fmt.Printf("%d:%s\n", len(wb.Bytes()), string(wb.Bytes()))
	w.Flush()
	fmt.Printf("%d:%s\n", len(wb.Bytes()), string(wb.Bytes()))
}
