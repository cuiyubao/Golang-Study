package lib1

import "fmt"
//当前lib1包提供的api
func Lib1Test()  {
	fmt.Println("Lib1Test执行了")
}

func init()  {
	fmt.Println("lib1.init()执行了")
}