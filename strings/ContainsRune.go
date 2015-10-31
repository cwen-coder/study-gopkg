package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.ContainsRune("team", rune('m')))    //true
	fmt.Println(strings.ContainsRune("failure", rune('w'))) //false
	fmt.Println(strings.ContainsRune("谢foo", rune('谢')))    //true
	fmt.Println(strings.ContainsRune("", 30))               //false
}
