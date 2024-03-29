package main

import (
	"fmt"
	"time"
)

var asyncChann = make(chan int,1) //缓冲长度为1，put可以比take多一次

func producer(){
	for i:=0;i<10;i++{
		asyncChann <-i //往channel里send 一个元素
		fmt.Printf("SEND %d\n",i) // SEND 0 1 2 3
	}
}

func consumer(){
	for i:=0;i<3;i++{
		v:= <-asyncChann //从channel里take一个元素
		fmt.Printf("take %d\n",v) // take 0 1 2
	}
}

func main(){
	go consumer()
	//producer() //缓冲满时，在main协程是继续send会fatal error : all goroutines are asleep - deadlock!
	go producer() //缓冲满时，在子协程是继续send 会阻塞
	time.Sleep(time.Duration(1) * time.Second)
}
