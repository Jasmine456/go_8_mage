package main

import "fmt"

var sc = make(chan struct{}) //channel仅作为协程间同步的工具，不需要床底具体的数据，管道类型可以用struct{},空的结构体不占用内存空间




func subG(){
	fmt.Println("subG finish")
	sc <- struct{}{}
	//close(sc)
}

func main(){
	go subG() //启动子协程
	<-sc //等待子协程结束。关闭管道或者往管道里send一个数据，改行都可以解除阻塞
}
