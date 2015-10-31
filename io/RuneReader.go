package main

import (
	"fmt"
	"strings"
)

func main() {
	index := Utf8Index("Go语言学习园地", "学习")
	fmt.Println(index)
}

func Utf8Index(str, substr string) int {
	ascilPos := strings.Index(str, substr)
	if ascilPos == -1 || ascilPos == 0 {
		return ascilPos
	}

	pos := 0
	totalSize := 0
	reader := strings.NewReader(str)
	for _, size, err := reader.ReadRune(); err == nil; _, size, err = reader.ReadRune() {
		totalSize += size
		pos++
		if totalSize == ascilPos {
			return pos
		}
	}
	return pos
}
