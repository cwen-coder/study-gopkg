package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	reader, _ := os.Open("test.txt")
	buf := make([]byte, 13)
	n, err := io.ReadFull(reader, buf)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
	fmt.Printf("%s", buf)
}
