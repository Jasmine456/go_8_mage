package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type HelloService struct {
}

// hello(request)
//Greet函数的签名：第一个参数Request 是string，第二个参数是一个指针
func (h *HelloService) Greet(request string, resp *string) error {

	*resp = fmt.Sprintf("hello,%s", request)
	return nil
}

func main() {
	//给对象的方法暴露成一个RPC到网络上
	//	HelloService：RPC网络上 这个服务的名称，类似于pkg名称
	//	&HelloService{}:提供RPC方法实现的对象 (Fun Greet)
	//	客户端类似于这样来调用 HelloService.Greet()

	//&HelloService{}: 的方法是需要约束，约束了接口的方式 Fn(request <T>,response<*T>)error
	rpc.RegisterName("HelloService", &HelloService{})

	//	已经注册需要暴露的RPC到RPC框架内，需要设置RPC
	//	使用net包(TCP/UDP),来进行网络通信
	// 然后我们建立一个唯一的TCP链接
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	// 通过rpc.ServeConn函数在该TCP链接上为对方提供RPC服务。
	// 每Accept一个请求，就创建一个goroutie进行处理
	//不断的从socket 读取报文，然后处理，如何处理交给RPC框架
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		// 前面都是tcp的知识, 到这个RPC就接管了
		// 因此 你可以认为 rpc 帮我们封装消息到函数调用的这个逻辑,
		// 提升了工作效率, 逻辑比较简洁，可以看看他代码
		//把network较为RPC框架
		go rpc.ServeConn(conn)
	}

}
