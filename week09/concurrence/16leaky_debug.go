package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

func main(){
	go func(){
		// 在8080 端口接收debug
		if err:= http.ListenAndServe("localhost:8080",nil);err!=nil{
			panic(err)
		}
	}()

	go func(){
	//	每隔1秒打印一次协程数量
		ticker:=time.NewTicker(1*time.Second)
		for {
			<-ticker.C
			fmt.Printf("当前协程数：%d\n",runtime.NumGoroutine())
		}
	}()

	for{
		//handle()
	}
}
