package main

import "fmt"

/*
2.定义结构体方法时，用结构体的值和指针有什么区别？用代码验证一下有没有对结构体进行拷贝
 */

type User struct {
	Name string
	Age int
}

//用结构体的值
func (u User)hello(man string){
	u.Name="jasmine"
	fmt.Printf("hi,%s,my name is %s\n",u.Name,man)
}

//用结构体的指针
func (u *User)hello2(man string){
	u.Name="jasmine"
	fmt.Printf("hi,%s,my name is %s\n",u.Name,man)
}


func main(){
	u:= new(User)
	u.Name="nina"
	//用结构体的值，是深拷贝，在函数中修改不会影响当前结构体的值
	u.hello(u.Name)
	println(u.Name)
	//用结构体的指针，是浅拷贝，在函数中的修改会直接修改结构体
	u.hello2(u.Name)
	println(u.Name)

}
