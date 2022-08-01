package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main(){ //main协程
	const N = 10
	wg:= sync.WaitGroup{}
	wg.Add(N) //加N
	for i:=0;i<N;i++{
		go func(a,b int){ //开N个子协程
			defer wg.Done() //减1
			fmt.Printf("i=%d\n",a)
			time.Sleep(10 * time.Millisecond)
			_=a+b
		}(i,i+1) //这里有个闭包的问题，如果这里打印i，而不把i传入到这个匿名函数，打印这个变量是有问题的
	}
	fmt.Printf("当前协程数：%d\n",runtime.NumGoroutine())
	wg.Wait()//等待减为0
	fmt.Printf("当前协程数：%d\n",runtime.NumGoroutine())
}
