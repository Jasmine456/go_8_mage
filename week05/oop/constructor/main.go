package main

import (
	"fmt"
	"sync"
)

type User struct {
	Name string
	Age int
	Sex byte
}

func NewDefaultUser() *User{
	return &User{
		Name:"",
		Age: -1,
		Sex:3,
	}
}

func NewUser(name string,age int,sex byte) *User {
	return &User{
		Name: name,
		Age: age,
		Sex: sex,
	}
}

var (
	sUser *User
	uOnce sync.Once
)

func GetUserInstance() *User {
	uOnce.Do(func(){ //确保即使在开发的情况下，下面3行代码在整个go进程里只会被执行一次
		if sUser == nil{
			sUser = NewDefaultUser()
		}
	})
	return sUser //sUser是个全局变量，每次调用GetUserInstance() 返回的都是它
}

func main(){
	u := User{} //构造一个空的User，各字段都取相应的数据类型的默认值
	fmt.Printf("name=%s,age=%d,sex=%d\n",u.Name,u.Age,u.Sex)
	up := new(User) //构造一个空的User，并返回其指针
	fmt.Printf("name=%s,age=%d,sex=%d\n",up.Name,up.Age,up.Sex)
	//通过自定义的构造函数，返回一个User指针
	up = NewDefaultUser()
	fmt.Printf("name=%s,age=%d,sex=%d\n",up.Name,up.Age,up.Sex)
	up = NewUser("张三",18,1)
	fmt.Printf("name=%s,age=%d,sex=%d\n",up.Name,up.Age,up.Sex)

	//	单例模式，调用GetUserInstance()得到的是同一个User实例
	su1 := GetUserInstance()
	su2 := GetUserInstance()
	// 修改su1会影响su2
	su1.Name = "令狐一刀"
	su1.Age = 100
	su1.Sex= 2
	fmt.Printf("name=%s,age=%d,sex=%d\n",su2.Name,su2.Age,su2.Sex)
	//GetUserInstance()
}