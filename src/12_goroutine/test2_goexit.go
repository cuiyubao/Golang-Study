package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	//用go创建承载一个形参为空  返回值为空的一个函数

	go func() {
		defer fmt.Println("A.defer")
		func() {
			defer fmt.Println("B.defer")
			//退出当前的goroutine
			runtime.Goexit()
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()

	go func(a int, b int) {
		fmt.Println("a=", a, "b=", b)
	}(10, 20)

	for true {
		time.Sleep(1 * time.Second)
	}

}
