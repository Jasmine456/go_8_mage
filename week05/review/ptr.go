package main

import (
	"fmt"
)
type User struct {
	name string
	age int
}

func update_user(arr []User){
	 arr[0].age=18

}
func main(){

	var arr  []User
	arr = []User{
		{"nina",25},
	}

	//arr[0].name="Jasmine"
	//fmt.Println(arr[0])
	update_user(arr)
	fmt.Println(arr[0])
}
