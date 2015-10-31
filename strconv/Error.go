package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	err := errors.New("too large")
	er := strconv.NumError{Func: "anyFunc", Num: "1e100", Err: err}
	fmt.Println(er.Error())
}
