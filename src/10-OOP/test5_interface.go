package main

import "fmt"

//interface{}是万能的数据类型
func myFunc(arg interface{}) {
	fmt.Println("myFunc is called...")
	fmt.Println(arg)

	//interface{}该如何区分 此时引用的底层数据类型到底是什么
	//给interface提供了"类型断言"的机制
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string type")
	} else {
		fmt.Println("arg is string type,value=", value)
		fmt.Printf("value type is %T\n", value)
	}
}

type Book1 struct {
	auth string
}

func main5() {
	book := Book1{"Golan"}
	myFunc(book)
	myFunc(123)
	myFunc("abc")

}
