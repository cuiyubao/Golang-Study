package main

import "fmt"

func deferFunc() int {
	fmt.Println("defer func called")
	return 0
}

func returnFunc() int {
	fmt.Println("return func called")
	return 0
}

func returnAndDefer() int {
	defer deferFunc()
	return returnFunc()
}

func changeValue(p *int)  {
	fmt.Println(p)
	*p=10
}


func main()  {

	a:=1
	changeValue(&a)
	fmt.Println(a)

}