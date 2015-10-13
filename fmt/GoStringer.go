package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	Name string
	Age  int
	Sex  int
}

func (this *Person) GoString() string {
	return "&Person{Name is " + this.Name + ", Age is " + strconv.Itoa(this.Age) + ", Sex is " + strconv.Itoa(this.Sex) + "}"
}

func main() {
	p := &Person{"CWen", 22, 0}
	fmt.Printf("%#v", p)
}
