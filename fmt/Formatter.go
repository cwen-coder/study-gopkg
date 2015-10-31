package main

import (
	"bytes"
	"fmt"
	"strconv"
)

type Person struct {
	Name string
	Age  int
	Sex  int
}

func (this *Person) String() string {
	buffer := bytes.NewBufferString("This is ")
	buffer.WriteString(this.Name + ", ")
	if this.Sex == 0 {
		buffer.WriteString("He ")
	} else {
		buffer.WriteString("She ")
	}
	buffer.WriteString("is ")
	buffer.WriteString(strconv.Itoa(this.Age))
	buffer.WriteString(" year old.")
	return buffer.String()
}

func (this *Person) Format(f fmt.State, c rune) {
	if c == 'L' {
		f.Write([]byte(this.String()))
		f.Write([]byte(" Person has three fields."))
	} else {
		// 没有此句，会导致 fmt.Printf("%s", p) 啥也不输出
		f.Write([]byte(fmt.Sprintln(this.String())))
	}
}

func main() {
	p := &Person{"CWen", 22, 0}
	fmt.Printf("%L", p)
}
