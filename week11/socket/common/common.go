package common

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

type Request struct {
	A int
	B int
}

type Response struct {
	Sum int
}

var sigChan = make(chan os.Signal)

func DealSigPipe(){
	fmt.Println("Deal 开始")
	signal.Notify(sigChan,syscall.SIGPIPE)
	fmt.Println("Deal 中间")
	for {
		fmt.Println("Deal 进入循环")
		select{
		case sig:=<-sigChan:
			fmt.Printf("receive signal %d\n",sig)
		}
		fmt.Println("Deal 退出循环")
	}
	fmt.Println("Deal 结束")
}