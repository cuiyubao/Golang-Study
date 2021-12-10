package main

import (
	"fmt"
)

func main() {

	cityMap := make(map[string]string)
	//添加
	cityMap["China"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "NewYork"
	//遍历
	for key, value := range cityMap {
		fmt.Println("key=", key)
		fmt.Println("value=", value)
	}

	//删除
	delete(cityMap, "China")

	//修改
	cityMap["USA"]="Dc"

	fmt.Println("====")
	//遍历
	for key, value := range cityMap {
		fmt.Println("key=", key)
		fmt.Println("value=", value)
	}

	intMap := make(map[int]string)
	intMap[1]="a"
	intMap[2]="a"
	intMap[3]="a"
	intMap[4]="a"

	cnumList:=make([]int,0)
	for key, _ := range intMap {
		cnumList=append(cnumList, key)
	}
	omoType:=0
	//lnum:="5f1eddfd9fc141a5a468923bec8d9fb7"
	type CnumParam struct {
		Cnum int `json:"cnum"`
		OmoType int `json:"omo_type"`
	}
	type VideoAndLessonPlansParam struct {
		Lnum string `json:"lnum"`
		cnum int `json:"cnum"`
		CnumParamList []CnumParam `json:"cnum_list"`
	}
	cnumParamList:=make([]CnumParam,0)
	for _, value := range cnumList {
		cnumParam:=CnumParam{Cnum: value,OmoType: omoType}
		cnumParamList=append(cnumParamList, cnumParam)
	}
	_=VideoAndLessonPlansParam{CnumParamList: cnumParamList}

	fmt.Println("============")


	if _ ,ok:= cityMap["USA"];ok{
         fmt.Println("result:"+cityMap["China"])
         fmt.Println("result1:"+cityMap["China1"])
	}else {
		fmt.Println("错误")
	}

	


}