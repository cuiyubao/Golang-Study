package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (this User) Call() {
	fmt.Println("user is called ...")
	fmt.Printf("%v\n", this)
	fmt.Println("", this)

}

func DoFiledAndMethod(input interface{}) {
	//获取input的type
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType is:", inputType.Name())
	//获取input的value
	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue is:", inputValue)

	//分别通过type获取里面的字段
	//1.获取interface的reflect.Type,通过Type得到NumFiled,进行遍历
	//2. 得到每个field,数据类型
	//3. 通过field有一个Interface()方法得到对应的value
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()
		fmt.Printf("%s:%v=%v\n", field.Name, field.Type, value)
	}
	//通过type获取里面的方法  调用
	for i := 0; i < inputType.NumMethod(); i++ {
		method := inputType.Method(i)
		fmt.Printf("%s:%v\n", method.Name, method.Type)
	}
}

func main5() {
	user := User{1, "Aceld", 18}
	user.Call()
	DoFiledAndMethod(user)
}
