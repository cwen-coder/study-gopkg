package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	web := bytes.NewBuffer(nil)
	w := bufio.NewWriter(web)
	w.WriteByte('a')
	w.Flush()
	fmt.Println(string(web.Bytes()))
}
