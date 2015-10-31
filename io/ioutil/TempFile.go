package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	f, err := ioutil.TempFile("", "temp")
	defer f.Close()
	fmt.Println(f.Name())
	if err != nil {
		fmt.Println("create tempfile error")
		return
	}
	f.WriteString("hellp world")
	b, err1 := ioutil.ReadFile(f.Name())
	if err1 != nil {
		fmt.Println("read error")
		return
	}
	fmt.Println(string(b))
}
