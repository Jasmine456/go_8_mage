package main

import (
	"fmt"
	"time"
)

var cloch = make(chan int, 1)
var cloch2 = make(chan int, 1)

func traverseChannel(){
	for ele := range cloch {
		fmt.Printf("receive %d\n",ele)
	}
}

func traverseChannel2(){
	for {
		if ele,ok := <-cloch2;ok{ // ok==true 代表管道还没有close
			fmt.Printf("receive %d\n",ele)
		}else { //管道关闭后，读操作会立即返回“0”值
			fmt.Printf("channel hava been closed,receive %d\n",ele)
			break
		}
	}
}

func main(){
	cloch <- 1
	//close(cloch)
	traverseChannel()
	fmt.Println("===========")
	//go traverseChannel2()
	//cloch2 <- 1
	//close(cloch2)
	time.Sleep(10 * time.Millisecond)

}