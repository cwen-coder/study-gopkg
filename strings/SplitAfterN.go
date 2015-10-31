package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("%q\n", strings.SplitAfterN("a,b,c", ",", 2)) //["a," "b,c"]
	fmt.Printf("%q\n", strings.SplitAfterN("a,b,c", "", 5))  //["a" "," "b" "," "c"]
	fmt.Printf("%q\n", strings.SplitAfterN("a,b,c", ",", 1)) //["a,b,c"]
}
