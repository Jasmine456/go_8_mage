package main

import (
	"errors"
	"fmt"
)

func goo(x int) int {
	fmt.Printf("x=%d\n", x)
	return x
}

func foo(a, b int, p bool) int {
	c := a*3 + 9
	//	defer 是先进后出，即逆序执行
	defer fmt.Println("first defer")
	d:=c+5
	defer fmt.Println("second defer")
	e:=d/b //如果发生panic，则后面的defer不会执行
	if p{
		panic(errors.New("my error")) //主动panic,代码飘红暂时不用理会
	}
	defer fmt.Println("third defer")
	return goo(e) //defer是在函数临退出前执行，不是在代码的return语句之前执行，因为return语句不是原子操作

}

func main(){
	foo(3,4,false)
	fmt.Println("==========")
	foo(3,4,true) //panic: my error
	fmt.Println("==========")
	foo(5,0,false) //panic: runtime error: integer divide by zero
	fmt.Println("============")

}