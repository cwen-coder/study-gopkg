package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	patterns := []string{"abc", "efg"}
	replacer := strings.NewReplacer(patterns...)
	format := replacer.Replace("abc is abc is abc")
	fmt.Println(format)
	//efg is efg is efg
	replacer.WriteString(os.Stdout, "abc is abc is abc")
	//efg is efg is efg
}
