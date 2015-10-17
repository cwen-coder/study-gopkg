package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Count("cheese", "e"))  //3
	fmt.Println(strings.Count("cheese", "ee")) //1
	fmt.Println(strings.Count("five", ""))     // 5
}
