package main

import "fmt"

func main1() {
	//声明slice1是一个切片，并且初始化  默认值是1，2，3 长度len是3
	//slice1 := []int{1, 2, 3}

	//声明slice1是一个切片，但并没有给slice分配空间
	//var slice1 []int
	//slice1 =make([]int,3) //开辟3个空间，默认值是0

	//声明slice1是一个切片，同时给slice分配空间，3个空间，初始化值是0，通过:=推导出slice是一个切片
	slice1 := make([]int, 3)
	fmt.Printf("len=%d, slice=%v\n", len(slice1), slice1)

	//判断一个slice是否为空
	if slice1 == nil {
		fmt.Println("slice1是一个空切片")
	}

	var numbers = make([]int, 3, 5)
	fmt.Printf("len=%d,cap=%d,slice=%v\n", len(numbers), cap(numbers), numbers)
	//向numbers切片里面追加一个元素1，len=4,cap=5,slice=[0 0 0 1]
	numbers = append(numbers, 1)
	fmt.Printf("len=%d,cap=%d,slice=%v\n", len(numbers), cap(numbers), numbers)
	//len=5,cap=5,slice=[0 0 0 1 2]
	numbers = append(numbers, 2)
	fmt.Printf("len=%d,cap=%d,slice=%v\n", len(numbers), cap(numbers), numbers)

	//向一个容量cap已经满的slice append，cap会翻倍  len=6,cap=10,slice=[0 0 0 1 2 3]
	numbers = append(numbers, 3)
	fmt.Printf("len=%d,cap=%d,slice=%v\n", len(numbers), cap(numbers), numbers)

	fmt.Println("=============")

	s := []int{1, 2, 3}  //len 3 cap=3
	s1 := s[0:2]   //[1,2]
	//修改s1的同时 s也会变化   s[1,10,3]  s1[1,10]
	s1[1]=10
	fmt.Println(s1)


	s2:= make([]int,2,2)
	//copy 可以将底层数组的slice一起进行拷贝
	copy(s2,s)
	s2=append(s2,11)
    fmt.Println(s2)



}
