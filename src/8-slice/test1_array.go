package main

import (
	"fmt"
)

func printArray(myArray [4]int) {
	//值拷贝
	for index, value := range myArray {
		fmt.Println("index=", index, "value=", value)
	}
}

func main() {
	//固定长度的数组 长度为10 每个值为0
	var myArray1 [10]int
	// 不指定容量 myArray11就是nil
	var myArray11 []int
	if myArray11 == nil {
		fmt.Println("myArray11是nil")
	}
	//固定长度为10 myArray2[0]=9 依次...
	myArray2 := [10]int{9, 8, 7, 6}

	//固定长度为4
	var myArray3 [4]int
	for i := 0; i < len(myArray1); i++ {
		myArray1[i] = i
		fmt.Println(myArray1[i])
	}

	for index, value := range myArray2 {
		fmt.Println("index=", index, "value=", value)
	}

	//查看数组的数据类型
	fmt.Printf("myArray1 type=%T\n", myArray1)
	fmt.Printf("myArray2 type=%T\n", myArray2)
	fmt.Printf("myArray3 type=%T\n", myArray3)

	printArray(myArray3)

}
