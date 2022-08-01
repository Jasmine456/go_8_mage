package main

import (
	"fmt"
	"runtime"
	"time"
)

// 模拟孙子协程
func grandson(){
	fmt.Println("grandson begin")
	fmt.Printf("routine num %d\n",runtime.NumGoroutine())
	time.Sleep(8*time.Second)
	fmt.Printf("routine num %d\n",runtime.NumGoroutine())
	fmt.Println("grandson finish")
}

// 模拟儿子辈协程
func child(){
	fmt.Println("child begin")
	go  grandson()
	//time.Sleep(100 * time.Millisecond)
	fmt.Println("child finish") //子协程退出后，孙子协程还在运行
}

func main(){
	go child()
	time.Sleep(10 * time.Second)
}

