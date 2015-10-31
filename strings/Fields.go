package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   ")) //Fields are: ["foo" "bar" "baz"]
	fmt.Printf("Fields are: %q", strings.Fields(" baz "))             //Fields are: ["bar"]
}
