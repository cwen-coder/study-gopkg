package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "heLLo worLd Ａｂｃ"
	us := strings.ToUpper(s)
	ls := strings.ToLower(s)
	ts := strings.ToTitle(s)
	fmt.Printf("%q\n", us) // "HELLO WORLD ＡＢＣ"
	fmt.Printf("%q\n", ls) // "hello world ａｂｃ"
	fmt.Printf("%q\n", ts) // "HELLO WORLD ＡＢＣ"
}
