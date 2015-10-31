package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.HasSuffix("cwenyin", "cw"))  //false
	fmt.Println(strings.HasSuffix("cwenyin", "yin")) //true
}
