package main

import (
	"io/ioutil"
	"strings"
)

func main() {
	s := strings.NewReader("hello world")
	r := ioutil.NopCloser(s)
	defer r.Close()
}
