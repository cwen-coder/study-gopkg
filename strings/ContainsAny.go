package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.ContainsAny("team", "i"))       //false
	fmt.Println(strings.ContainsAny("failure", "wwwi")) //true
	fmt.Println(strings.ContainsAny("foo", ""))         //false
	fmt.Println(strings.ContainsAny("", ""))            //false
}
