package main

import (
	"fmt"
	"reflect"
	"time"
)

func test(a string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(a)
	}
}

type Resume struct {
	Name string `info:"name" doc:"我的名字"`
	Sex  string `info："sex"`
}

func findTag(str interface{})  {
	elem := reflect.TypeOf(str).Elem()
	for i := 0; i < elem.NumField(); i++ {
		taginfo := elem.Field(i).Tag.Get("info")
		tagdoc := elem.Field(i).Tag.Get("doc")
		fmt.Println("info:",taginfo," doc:",tagdoc)
	}
}



func main() {
	//2-var myMap1 map[string]string
	//if myMap1 == nil {
	//	fmt.Println("这是一个空map")
	//}
	//myMap1 = make(map[string]string, 3)
	//myMap1["one"] = "java"
	//myMap1["two"] = "c++"
	//s, ok := myMap1["one"]
	//fmt.Println(s, ok)
	//fmt.Println(len(myMap1))
	//fmt.Printf("%#v", myMap1)
	//
	//myMap2 := map[string]string{
	//	"name": "dabao",
	//	"addr": "山西",
	//}
	//fmt.Println(myMap2)

	findTag(Resume{"cuiyubao","男"})
}
