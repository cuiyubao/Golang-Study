package main

import "fmt"

//本质是指针
type AnimalIF interface {
	Sleep()
	GetColor() string //获取动物的颜色
	GetType() string  //获取动物的种类
}


type AnimalIF1 interface {
	Sleep()
	GetColor() string //获取动物的颜色
}

//具体的类
type Cat struct {
	color string //🐱的颜色
}

func (this *Cat) Sleep() {
	fmt.Println("Cat is Sleep")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "cat"
}

type Dog struct {
	color string //狗的颜色
}

func (this *Dog) Sleep() {
	fmt.Println("Dog is Sleep")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "dog"
}

func showAnimal(animal AnimalIF) {
	animal.Sleep() //多态
	fmt.Println("color=", animal.GetColor())
	fmt.Println("kind=", animal.GetType())
}

func main4() {
	/*    var animal AnimalIF //接口的数据类型，父类指针
	      animal=&Cat{"Green"}
	      animal.Sleep()  //调用的是Cat的Sleep()方法，多态的现象

	      animal=&Dog{"Yellow"}
	      animal.Sleep() //调用的是Dog的Sleep()方法，多态的现象
	*/
	cat := Cat{"Green"}
	showAnimal(&cat)
	fmt.Println("=====")
	dog := Dog{"Yellow"}
	showAnimal(&dog)


}
