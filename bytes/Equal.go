package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(bytes.Equal([]byte{1, 2, 3}, []byte{1, 2, 3}))
	fmt.Println(bytes.Equal([]byte{1, 2, 3}, []byte{1, 2}))
	fmt.Println(bytes.Equal([]byte{1, 2, 3}, nil))
	fmt.Println(bytes.Equal([]byte{}, nil))
	fmt.Println(bytes.Equal(nil, nil))
}
