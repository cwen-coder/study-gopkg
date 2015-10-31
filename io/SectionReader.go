package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	reader, _ := os.Open("test.txt")
	defer reader.Close()
	sectionReader := io.NewSectionReader(reader, 5, 10)
	fmt.Println(reflect.TypeOf(sectionReader))
	var n, total int
	var err error
	p := make([]byte, 15)
	//Read
	for {
		n, err = sectionReader.Read(p)
		if err == io.EOF {
			fmt.Println("Find EOF so end total", total)
			break
		}
		total = total + n
		fmt.Println("Read value:", string(p[0:n]))
		fmt.Println("Read count:", n)
	}
	//ReadAt
	p = make([]byte, 15)
	n, _ = sectionReader.ReadAt(p, 4)
	fmt.Println("Read value:", string(p[0:n]))
	fmt.Println("Read count:", n)
	//Seek
	sectionReader1 := io.NewSectionReader(reader, 2, 20)
	sectionReader1.Seek(2, 0)
	p = make([]byte, 10)
	n, _ = sectionReader1.Read(p)
	fmt.Println("First read value:", string(p[0:n]))
	fmt.Println("First read count:", n)
	a, _ := sectionReader.Seek(2, 1)
	fmt.Println("off - base is", a)
	n, _ = sectionReader.Read(p)
	fmt.Println("Second Read value:", string(p[0:n]))
	fmt.Println("Second Read count:", n)
	sectionReader.Seek(8, 2)
	n, _ = sectionReader.Read(p)
	fmt.Println("Third read value:", string(p[0:n]))
	fmt.Println("Third read count:", n)

	//size
	sectionReader2 := io.NewSectionReader(reader, 5, 20)
	fmt.Println("Can read count:", sectionReader2.Size())
	p = make([]byte, 10)
	n, _ = sectionReader2.Read(p)
	fmt.Println("Read count:", n)
	fmt.Println("Can read count:", sectionReader2.Size())
}
