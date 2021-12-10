package main

import "fmt"

func main1() {
	//====>声明方式一
	//声明myMap1是一种map类型 key是string value是string
	var myMap1 map[string]string
	if myMap1 == nil {
		fmt.Println("myMap1是一个空map")
	}
	//在使用map前，需要先使用make给map分配数据空间
	myMap1 = make(map[string]string, 10)
	myMap1["one"] = "java"
	myMap1["two"] = "c++"
	myMap1["three"] = "python"
	fmt.Println(myMap1)

	//===>声明方式二
	myMap2 := make(map[string]string)
	myMap2["one"] = "java"
	myMap2["two"] = "c++"
	myMap2["three"] = "python"
	fmt.Println(myMap2)

	//===>声明方式三
	myMap3 := map[string]string{
		"name": "db",
		"age":  "12",
		"":     "12", //空key放不进map
		"addr": "",
	}
	fmt.Println(myMap3)
	myMap3=make(map[string]string)
	if myMap3 == nil {
		fmt.Println("myMap1是一个空map")
	}
    fmt.Println(len(myMap3))
}
