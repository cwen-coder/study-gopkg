package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	io.CopyN(os.Stdout, os.Stdin, 8)
	fmt.Println("Got EOF -- bye")
}
