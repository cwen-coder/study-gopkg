package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func main() {
	readers := []io.Reader{
		strings.NewReader("from strings reader"),
		bytes.NewBufferString("from bytes buffer"),
	}
	reader := io.MultiReader(readers...)
	data := make([]byte, 0, 1024)
	var (
		err error
		n   int
	)
	for err != io.EOF {
		tmp := make([]byte, 512)
		n, err = reader.Read(tmp)
		if err == nil {
			data = append(data, tmp[:n]...)
		} else {
			if err != io.EOF {
				panic(err)
			}
		}
	}
	fmt.Printf("%s\n", data)
}
