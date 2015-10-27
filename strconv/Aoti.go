package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(strconv.Atoi("12345"))
	fmt.Println(strconv.Atoi("abcd"))
	v := "10"
	if s, err := strconv.Atoi(v); err == nil {
		fmt.Printf("%T, %v", s, s)
	}
}
