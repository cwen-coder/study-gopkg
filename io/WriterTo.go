package main

import (
	"bytes"
	"os"
)

func main() {
	reader := bytes.NewReader([]byte("天天学习Go语言"))
	reader.WriteTo(os.Stdout)
}
