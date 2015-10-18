package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.LastIndexByte("chickenkenken", 'k')) //10
	fmt.Println(strings.LastIndexByte("chicken", 'd'))       //-1
}
