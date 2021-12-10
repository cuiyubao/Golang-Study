package main

import (
	"fmt"
)

func fool1(a string, b int) int {
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	c := 100
	return c
}

//返回多个返回值 匿名的
func fool2(a string, b int) (int, int) {
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	return 666, 777
}

//返回多个返回值 有形参名的
func fool3(a string, b int) (r1 int, r2 string) {
	fmt.Println("---foo3----")
	fmt.Println("a=", a)
	fmt.Println("b=", b)

	//r1 r2属于foo3的形参 初始化的默认值是r1 0 ,r2 ""
	//r1 r2作用域空间，是foo3整个函数体的{}空间
	fmt.Println("r1=",r1)
	fmt.Println("r2=",r2)
	//给名称返回值赋值
	r1 = 20
	r2 = "ss"
	return r1, r2
}

func fool4(a string, b int) (r1, r2 int) {
	fmt.Println("---foo4----")
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	//给名称返回值赋值
	r1 = 1000
	r2 = 4000
	return r1, r2
}

func test(arg interface{}) {
	value, ok := arg.(string)
	fmt.Println(value)
	fmt.Printf("%T\n", value)
	fmt.Println(ok)
}

func main() {
	c := fool1("abc", 55)
	fmt.Println("c=", c)

	ret1, ret2 := fool2("abc", 66)
	fmt.Println("ret1=", ret1, ",ret2=", ret2)

	ret3, ret4 := fool3("eee", 44)
	fmt.Println("ret3=", ret3, ",ret4=", ret4)

	ret1, ret2 = fool4("fff", 646)
	fmt.Println("ret1=", ret1, ",ret2=", ret2)
}
