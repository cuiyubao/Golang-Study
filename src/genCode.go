package main

import (
	"github.com/atotto/clipboard"
	"github.com/xlzd/gotp"
	"time"
)

func main() {
	defaultTOTPUsage()
}

func defaultTOTPUsage() {
	tt := int(time.Now().Unix() / 60)
	//println(tt)
	// 替换成上一步得到的secret
	scret := ""
	hotp := gotp.NewDefaultHOTP(scret)
	// 111111替换成自己的固定密码
	result := "123456" + hotp.At(tt)
	clipboard.WriteAll(result)
}
