package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	reader, _ := os.Open("test.txt")
	buf := make([]byte, 9)
	n, err := io.ReadAtLeast(reader, buf, 4)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
	fmt.Printf("%s", buf)
}
