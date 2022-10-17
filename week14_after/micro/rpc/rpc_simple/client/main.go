package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	//	通过网络来和rpc server建立通信
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	//	通过client来调用server方法
	//	客户端类似于这样来调用 HelloService.Greet()
	var resp string
	if err := client.Call("HelloService.Greet", "alice", &resp); err != nil {
		panic(err)
	}

	//	处理返回结果
	fmt.Println(resp)
}
