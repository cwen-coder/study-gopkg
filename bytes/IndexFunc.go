package main

import (
	"bytes"
	"fmt"
)

/*IndexFunc interprets s as a sequence of UTF-8-encoded Unicode code points.
It returns the byte index in s of the first Unicode code point satisfying f(c), or -1 if none do.*/
func main() {
	s := []byte("123456677")
	f := func(a rune) bool {
		if a > '6' {
			return true
		}
		return false
	}
	fmt.Println(bytes.IndexFunc(s, f))
}
