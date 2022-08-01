package main

import (
	"fmt"
	"time"
)

func Add(a,b int) int {
	fmt.Println("Add")
	return a+b
}

var add = func(a,b int) int{
	fmt.Println("add")
	return a+b
}

func main(){
	fmt.Println("add begin")
	go Add(2,3)
	go Add(3,9)
	fmt.Println("add over")

	go func(a,b int) int{
		fmt.Println("add")
		return a+b
	}(2,4)
	go func(a,b int) int{
		fmt.Println("add")
		return a+b
	}(3,9)

	go add(2,4)
	go add(3,9)

	time.Sleep(10*time.Millisecond)


}