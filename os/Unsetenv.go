package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Unsetenv("go")
	if err != nil {
		fmt.Println(err)
	}
}
