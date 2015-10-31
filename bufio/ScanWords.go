package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	const input = "This is The Golang Standard Library.\nWelcome you!"
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	fmt.Println(count)
}
