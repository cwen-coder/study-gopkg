package main

import (
	"fmt"
	"os"
)

func main() {
	fi, err := os.Create("Hello.go")
	if err != nil {
		fmt.Printf("Error:%v\n", err)
		return
	}

	defer fi.Close()
	fmt.Printf("The %s has been created!\n", fi.Name())
}
