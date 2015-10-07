package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := strings.NewReader("天天学习Go语言")
	reader.Seek(-6, os.SEEK_END)
	r, _, _ := reader.ReadRune()
	fmt.Printf("%c\n", r)
}
