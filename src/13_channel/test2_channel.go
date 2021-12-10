package main

import (
	"fmt"
	"time"
)

func main2() {
	c := make(chan int, 3) //带有缓冲的channel
	fmt.Println("len(c)=", len(c), " cap(c)=", cap(c))

	go func() {
		defer fmt.Println("子go程结束")
		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("子go程正在运行：len(c)=", len(c), " cap(c)=", cap(c))
		}
	}()
	time.Sleep(1 * time.Second)
	for i := 0; i < 3; i++ {
		num := <-c //从c中接收数据 并赋值给nm
		fmt.Println("num=", num)

	}
	fmt.Println("main 结束")
}
