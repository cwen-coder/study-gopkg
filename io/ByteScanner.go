package main

import (
	"bytes"
	"fmt"
)

func main() {
	buffer := bytes.NewBuffer([]byte{'a', 'b'})
	ch, _ := buffer.ReadByte()
	err := buffer.UnreadByte()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%c\n", ch)
	ch, _ = buffer.ReadByte()
	fmt.Printf("%c\n", ch)

}
