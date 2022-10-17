package main

import (
	"go_8_mage/week14_after/micro/rpc/rpc_interface/service"
	"log"
	"net/rpc"
)

var _ service.HelloService = (*HelloServiceClient)(nil)

func NewHelloServiceClient(network, address string) *HelloServiceClient {
	//	通过网络来和rpc server建立通信
	client, err := rpc.Dial(network, address)
	if err != nil {
		log.Fatal("dialing:", err)
	}

	//	通过client来调用server方法
	//	客户端类似于这样来调用 HelloService.Greet()
	return &HelloServiceClient{
		client: client,
	}
}

type HelloServiceClient struct {
	client *rpc.Client
}

//RPC Client
func (c *HelloServiceClient) Greet(request string, response *string) error {
	return c.client.Call("HelloService.Greet", request, response)
}

//作为RPC client的使用方
func main() {
	client := NewHelloServiceClient("tcp", "localhost:1234")
	var resp string
	err := client.Greet("alice", &resp)
	if err != nil {
		panic(err)
	}
}
