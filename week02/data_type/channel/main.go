package main

import "fmt"

func main(){
	var ch chan int
	if ch == nil{
		fmt.Println("ch is nil")
	}
	if len(ch) == 0{
		fmt.Println("ch length is 0")
	}
	ch = make(chan int,8)
	//向管道里写入元素
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5
	v := <-ch // 从管道里取走（recv）数据
	fmt.Println(v)
	v = <-ch
	fmt.Println(v)
	fmt.Println()

	close(ch)
//	遍历管道里剩下的元素
	for ele := range ch{
		fmt.Println(ele)
	}

//	定义只读和只写的管道
	read_only := make(<-chan int)
	write_only := make(chan<- int)

//	定义只读和只写的channel意义不大，一般用于在参数传递中

}
//只能向channel里写数据
func send(c chan<- int){
	c<-1
}
//返回一个只读的channel
func (c *Context) Done() <-chan struct{}{
	return nil
}