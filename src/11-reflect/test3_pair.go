package main

import (
	"fmt"
	"reflect"
)

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriterBook()
}

type Book struct {
}

func (this *Book) ReadBook() {
	fmt.Println("Read a book")
}

func (this *Book) WriterBook() {
	fmt.Println("Writer a book")
}

func main3() {
	//b: pair<type:Book,value:book{}地址>
	b := &Book{}

	// r: pair<type:,value:>
	var r Reader
	//r: pair<type:Book,value:book{}地址>
	r = b
	r.ReadBook()

	var w Writer
	//r: pair<type:Book,value:book{}地址>
	w = r.(Writer) //此处的断言是很么会成功，因为w r的具体type是一致的

	w.WriterBook()

	of := reflect.TypeOf(w)
	fmt.Println(of)
}
