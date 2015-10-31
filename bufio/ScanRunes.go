package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	const input = "我是捏第三方都是积分"
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanRunes)
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	fmt.Println(count)
}
