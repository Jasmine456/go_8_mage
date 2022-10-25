package main

import (
	"fmt"
	"go_8_mage/week14_after/micro/rpc/rpc_interface/service"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

var _ service.HelloService = (*HelloServiceClient)(nil)

func NewHelloServiceClient(network, address string) *HelloServiceClient {
	//	通过网络来和rpc server建立通信
	conn, err := net.Dial(network, address)
	if err != nil {
		panic(err)
	}

	//把network交给rpc框架处理,默认使用Gob
	//client:= rpc.NewClient(conn)

	//客户端使用json编解码器

	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

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
	fmt.Println(resp)
}
