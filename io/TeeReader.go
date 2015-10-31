package main

import (
	"io"
	"os"
)

func main() {
	reader, _ := os.Open("test.txt")
	p := io.TeeReader(reader, os.Stdout)
	p.Read(make([]byte, 20))

}
