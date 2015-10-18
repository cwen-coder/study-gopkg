package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.IndexRune("chicken", 'k')) //4
	fmt.Println(strings.IndexRune("chicken", 'd')) //-1
}
