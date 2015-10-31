package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	s := strings.NewReader("hello word!")
	b, err := ioutil.ReadAll(s)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Println(string(b))
}
