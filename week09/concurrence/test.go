package main

import (
	"fmt"
	"time"
)

func main() {
	ch2 := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Millisecond*100)
			ch2 <- i
		}
	}()

	go func() {
		for i := range ch2 {
			fmt.Println("ch2取出值：", i)
			//if i == 9 {
			//	fmt.Println("结束ch2取值！")
			//	break
			//}
		}
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("main routine exit")
}
