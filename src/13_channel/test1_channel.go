package main

import (
	"fmt"
	"time"
)

func main1() {
	c := make(chan int)
	go func() {
		defer fmt.Println("goroutine结束")
		fmt.Println("goroutine正在运行...")
		//time.Sleep(2 * time.Second)
		c <- 666 //将666 发送给c
		fmt.Println("goroutine数据发送完成...")
	}()

	time.Sleep(10 * time.Second)
	num := <-c //从c中接收数据  并赋值给num
	fmt.Println("num=", num)
	fmt.Println("main goroutine 结束..,")
}
