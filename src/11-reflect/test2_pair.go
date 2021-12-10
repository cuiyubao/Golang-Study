package main

import (
	"fmt"
	"io"
	"os"
)

func main2() {
	// tty: pair<type:*os.File，value:"/dev/tty" 文件描述符>
	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
	if err != nil {
		fmt.Println("open file error", err)
		fmt.Println("open file error", tty)
		return
	}
	//r: pair<type:, value>
	var r io.Reader
	// r: pair<type:*os>File,value:"/dev/tty" 文件描述符>
	r = tty

	var w io.Writer
	//w：pair<type:*os>File,value:"/dev/tty" 文件描述符>
	w = r.(io.Writer)
	w.Write([]byte("HELLO THIS"))
}
