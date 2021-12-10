package main

import "fmt"

type Human struct {
	name string
	sex  string
}

type SuperMan struct {
	Human //SuperMan类继承了Human类的方法
	level int
}

//重新定义父类的方法Eat()
func (this *SuperMan) Eat() {
	fmt.Println("SuperMan.Eat()...")
}

//子类的新方法
func (this *SuperMan) Fly() {
	fmt.Println("SuperMan.Fly()...")
}

func (this *Human) Eat() {
	fmt.Println("Human.Eat()...")
}
func (this *Human) Walk() {
	fmt.Println("Human.Walk()...")
}

func main3() {
	human := Human{"zhangsan", "12"}
	human.Eat()
	human.Walk()
	fmt.Println("====")
	//定义一个子类对象
	//s := SuperMan{Human{"cuiyubao", "nan"}, 99}
	var s SuperMan
	s.name = "cuiyubao"
	s.sex = "nan"
	s.level = 99
	s.Walk() //父类的方法
	s.Eat()  //子类的方法
	s.Fly()  //子类的方法
}
