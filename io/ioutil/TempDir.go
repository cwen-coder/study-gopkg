package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	f, err := ioutil.TempDir("", "temp")
	if err != nil {
		fmt.Println("create tempDir error")
		return
	}
	fmt.Println(f)
}
