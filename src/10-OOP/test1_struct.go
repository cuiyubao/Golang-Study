package main

import "fmt"

type myint int

//定义一个结构体
type Book struct {
title string
auth string
}

func changeBook(book Book)  {
	book.auth="777"
}

func changeBook1(book *Book)  {
	book.auth="888"
}

func main() {
	/*var a myint=10
	fmt.Println("a=",a)
	fmt.Printf("type of a=%T\n",a)*/
	var book1 Book
	book1.title="goLang"
	book1.auth="zhang3"
	fmt.Printf("%v\n",book1)

	//changeBook(book1)
	//fmt.Printf("changeBook %v\n",book1)
	//changeBook1(&book1)
	//fmt.Printf("changeBook1 %v\n",book1)


}
