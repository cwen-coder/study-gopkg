package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Index("chickenkenken", "ken")) //4
	fmt.Println(strings.Index("chicken", "dmr"))       //-1
}
