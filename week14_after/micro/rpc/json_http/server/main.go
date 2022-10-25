package main

import (
	"fmt"
	"go_8_mage/week14_after/micro/rpc/rpc_interface/service"
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
)

type HelloService struct {
}

//如何约束HelloService实现HelloService接口
//var a1 service.HelloService = &HelloService{}
var _ service.HelloService = (*HelloService)(nil)

// hello(request)
//Greet函数的签名：第一个参数Request 是string，第二个参数是一个指针
func (h *HelloService) Greet(request string, resp *string) error {

	*resp = fmt.Sprintf("hello,%s", request)
	return nil
}

func NewRPCReadWriteCloserFromHTTP(w http.ResponseWriter,r *http.Request)*RPCReadWriteCloser{
	return &RPCReadWriteCloser{
		w,
		r.Body,
	}
}

type RPCReadWriteCloser struct {
	io.Writer
	io.ReadCloser
}


func main() {
	//给对象的方法暴露成一个RPC到网络上
	//	HelloService：RPC网络上 这个服务的名称，类似于pkg名称
	//	&HelloService{}:提供RPC方法实现的对象 (Fun Greet)
	//	客户端类似于这样来调用 HelloService.Greet()

	//&HelloService{}: 的方法是需要约束，约束了接口的方式 Fn(request <T>,response<*T>)error
	rpc.RegisterName("HelloService", &HelloService{})


	// RPC的服务架设在“/jsonrpc”路径，
	// 在处理函数中基于http.ResponseWriter和http.Request类型的参数构造一个io.ReadWriteCloser类型的conn通道。
	// 然后基于conn构建针对服务端的json编码解码器。
	// 最后通过rpc.ServeRequest函数为每次请求处理一次RPC方法调用
	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		//  Read(p []byte) (n int, err error)
		//Write(p []byte) (n int, err error)
		//Close() error

		//defer r.Body.Close()

		//http response: Write([]byte) (int, error)
		//r.Body: Read(p []byte) (n int, err error)
		//r.Body: Close bool

		//构造一个HTTP Conn
		conn := NewRPCReadWriteCloserFromHTTP(w, r)
		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})

	err:=http.ListenAndServe("localhost:1234",nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}
