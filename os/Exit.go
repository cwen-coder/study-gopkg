package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("runing...")
	defer fmt.Println("defer function")
	os.Exit(1)
	fmt.Println("finished")
}
