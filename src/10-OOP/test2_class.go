package main

import "fmt"

//如果类名首字母大写，表示其他包能够访问
type Hero struct {
	//如果类的属性首字母大写，表示该属性是对外能够访问的，否则的话只能够类的内部访问
	Name  string
	Ad    int
	Level int
}

/*func (this Hero)GetName()  {
	fmt.Println("Name=",this.Name)
}

func (this Hero)SetName(newName string)  {
	//this 是调用该方法对象的一个副本（拷贝）
	this.Name=newName
}*/

func (this *Hero) GetName() {
	fmt.Println("Name=", this.Name)
}

func (this *Hero) SetName(newName string) {
	//this 是调用该方法对象的一个副本（拷贝）
	this.Name = newName
}

func main2() {
	//创建一个对象
	hero := Hero{Name: "zhang3", Ad: 100, Level: 10}
	hero.SetName("cuiyubao")
	fmt.Println(hero)
}
