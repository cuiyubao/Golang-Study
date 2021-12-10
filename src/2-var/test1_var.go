package main

import "fmt"

//声明全局变量：方法一 方法二  方法三是可以的
// :=只能用在函数体中

func main() {
	//方法一：声明一个变量 默认的值为0
	var a int
	fmt.Println("a=", a)
	fmt.Printf("type of a=%T\n", a)

	//方法二：声明一个变量，初始化一个值
	var b int = 100
	fmt.Println("b=", b)
	fmt.Printf("type of b=%T\n", b)

	//方法三  在初始化的时候可以省去数据类型，通过值自动匹配当前变量的数据类型
	var c = 120
	fmt.Println("c=", c)
	fmt.Printf("type of c=%T\n", c)

	var cc = "abcd"
	fmt.Printf("cc =%s ,type of cc =%T\n", cc, cc)

	//方法四（常用方法） 省去var关键字，直接自动匹配
	e := 140
	fmt.Printf("type of e =%T\n", e)

	f := "wozhidao"
	fmt.Println("f=", f)
	fmt.Printf("type of f=%T\n", f)

	//声明多个变量
	var xx, yy int = 100, 200
	fmt.Println("xx=", xx, ",yy=", yy)
	var kk, ll = 100, "aceld"
	fmt.Println("kk=", kk, ",ll=", ll)

	//多行的多变量声明
	var (
		vv int  = 23
		jj bool = true
	)
	fmt.Println("vv=", vv, ",jj=", jj)
}
