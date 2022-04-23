package main

import "fmt"

func main(){
	var ch chan int
	fmt.Printf("ch is nil %t\n",ch == nil)
	fmt.Printf("len of ch is %d\n", len(ch))

	ch = make(chan int,10)
	fmt.Printf("len of ch is %d\n",len(ch))
	fmt.Printf("cap of ch is %d\n",cap(ch))
	for i := 0;i<10;i++{
		ch <- 3
	}
	fmt.Printf("len of ch is %d\n",len(ch))
	<-ch
	<-ch
	ch <- 3
	fmt.Printf("len of ch is %d\n",len(ch))
	fmt.Printf("cap of ch is %d\n",cap(ch))

	//两种循环方式是等价的，遍历管道
	close(ch)
	L:=len(ch)
	for i := 0;i<L;i++{
		ele := <-ch
		fmt.Println(ele)
	}
	fmt.Println("----------------------")

	//for ele := range ch{
	//	fmt.Println(ele)
	//}
	fmt.Printf("len of ch is %d\n", len(ch))
}
