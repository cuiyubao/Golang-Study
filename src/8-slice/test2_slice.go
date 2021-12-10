package main

import "fmt"

func printArrays(myArray []int) {
	//_表示匿名变量
	for _, value := range myArray {
		fmt.Println("value=", value)
	}
}

/*func main() {
	myArray := []int{1, 2, 3, 4} //动态数组 切片slice
	fmt.Printf("myArray type is %T\n", myArray)
	printArrays(myArray)
	its := append(myArray, 5)
	fmt.Println(&its)
}
*/