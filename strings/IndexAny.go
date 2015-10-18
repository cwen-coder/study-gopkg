package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.IndexAny("chickenkenkenken", "iken")) //2
	fmt.Println(strings.IndexAny("chicken", "dmr"))           //-1
}
