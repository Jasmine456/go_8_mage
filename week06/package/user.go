package mypackage //package 名和目录名可以不一致

import "fmt" //在$GOROOT/src目录下有fmt包

func init(){
	fmt.Println("enter package/user")
}

type User struct {
	Name string
	Age int
	Sex byte
}

func (u *User) Say(){
	fmt.Printf("Hi,I'm %s and %d years old\n",u.Name,u.Age)
}
