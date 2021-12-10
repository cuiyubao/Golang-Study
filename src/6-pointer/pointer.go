package main

import "fmt"

//func swap(a int, b int) {
//	var temp int
//	temp = a
//	a = b
//	b = temp
//}

func swap(a *int, b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp
}

func main() {
	var a int = 10
	var b int = 20
	swap(&a, &b)
	fmt.Println("a=", a, "b=", b)

}
