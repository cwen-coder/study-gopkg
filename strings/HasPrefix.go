package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.HasPrefix("cwenyin", "cw")) //true
	fmt.Println(strings.HasPrefix("cwenyin", "ta")) //false
}
