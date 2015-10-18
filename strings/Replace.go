package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))      //oinky oinky oink
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1)) //moo moo moo
}
