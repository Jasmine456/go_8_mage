package main

import (
	"fmt"
	"time"
)

var syncChann = make(chan int)

// 从管道取(拿)
func takeFromSyncChann() {
	if v, ok := <-syncChann; ok { // ok==true 说明管道还没有关闭close
		fmt.Printf("take %d from syncchronous channel\n", v)
	}
}

// 放入管道
func putToSyncChann() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("putToSyncChann 发生panic：%s\n", err)
		} else {
			fmt.Println("putToSyncChann success")
		}
	}()
	syncChann <- 10 //取操作没有准备好时，写管道会发生fatal error(注意不是panic，通过recover不能捕获)
}

func TestsSyncChann1() {
	go takeFromSyncChann() //消费者协程会阻塞2秒钟，等待put
	time.Sleep(2 * time.Second)
	putToSyncChann()
}

func TestsSyncChann2() {
	putToSyncChann() //取操作没有准备好时，写管道会发生fatal error
	go takeFromSyncChann()
}

func testSyncChan3() {
	go func() {
		for {
			if v, ok := <-syncChann; ok { //ok=true 说明管道还没有关闭close
				fmt.Printf("receive %d\n", v)
			} else {
				break
			}
		}
	}()
	for i := 0; i < 10; i++ {
		syncChann <- i
		fmt.Printf("send %d\n",i)
	}
	close(syncChann)
}

func main() {
	TestsSyncChann1()
	time.Sleep(3*time.Second)
	fmt.Println("=================")

	//TestsSyncChann2()
	//time.Sleep(3 * time.Second)
	//fmt.Println("=================")
	//testSyncChan3()
	//fmt.Println("==================")

}
