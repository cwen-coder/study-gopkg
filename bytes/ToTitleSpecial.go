package main

import (
	"bytes"
	"fmt"
	"unicode"
)

func main() {
	s := []byte("hello, world!")
	fmt.Println(string(bytes.ToTitleSpecial(unicode.AzeriCase, s)))
	fmt.Println(string(s))
}
