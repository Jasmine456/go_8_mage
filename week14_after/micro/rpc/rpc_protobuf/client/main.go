package main

import (
	"fmt"
	"net"
	"net/rpc"

	"go_8_mage/week14_after/micro/rpc/rpc_protobuf/codec/client"
	"go_8_mage/week14_after/micro/rpc/rpc_protobuf/service"
)

// 封装一个HelloServiceClient
var _ service.HelloService = (*HelloServiceClient)(nil)

func NewHelloServiceClient(network, address string) *HelloServiceClient {
	// 通过网络来和rpc server建立通信
	conn, err := net.Dial(network, address)
	if err != nil {
		panic(err)
	}

	// 把network交给 rpc框架处理, 默认使用Gob
	// client := rpc.NewClient(conn)

	// 使用自己编写的Protobuf Codec
	client := rpc.NewClientWithCodec(client.NewClientCodec(conn))

	return &HelloServiceClient{
		client: client,
	}
}

type HelloServiceClient struct {
	client *rpc.Client
}

// RPC Client
func (c *HelloServiceClient) Greet(request *service.Request, response *service.Response) error {
	return c.client.Call("HelloService.Greet", request, response)
}

// 作为RPC client的使用方
func main() {
	client := NewHelloServiceClient("tcp", "localhost:1234")
	resp := service.Response{}
	err := client.Greet(&service.Request{Value: "alice"}, &resp)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Value)
}
