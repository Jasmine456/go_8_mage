package main

import "fmt"

/*
1.实现一个结构体同时“继承”另外两个结构体
 */


type Person struct {
	id int
	name string
	age int
}
type Electronic struct {
	phone string
	handset string
}

type Student2 struct {
	*Person //指针实现继承(结构体指针
	Electronic
	score int
}

func main(){
	var stu Student2
	stu.Person=new(Person) //stu.Person是指针类型，不能为父类成员直接赋值，必须需要先new() 开辟空间
	stu.name = "jasmine"
	stu.age=27
	//stu.Person=&Person{2,"jasmine",90}  // 也可以通过该方式直接为父类成员赋值
	stu.score=90
	stu.phone="苹果13"
	fmt.Println(stu)
}
