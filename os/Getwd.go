package main

import (
	"fmt"
	"os"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("The current directory id: %s\n", pwd)
}
