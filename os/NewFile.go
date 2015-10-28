package main

import (
	"fmt"
	"os"
)

func main() {
	fi, err := os.Open("Hello.go")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer fi.Close()
	file := os.NewFile(fi.Fd(), "Word.go")
	defer file.Close()
	fmt.Printf("The %s has been new!\n", file.Name())
}
