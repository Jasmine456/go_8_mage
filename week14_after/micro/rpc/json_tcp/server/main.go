package main

import (
	"fmt"
	"go_8_mage/week14_after/micro/rpc/rpc_interface/service"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
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

func main() {
	//给对象的方法暴露成一个RPC到网络上
	//	HelloService：RPC网络上 这个服务的名称，类似于pkg名称
	//	&HelloService{}:提供RPC方法实现的对象 (Fun Greet)
	//	客户端类似于这样来调用 HelloService.Greet()

	//&HelloService{}: 的方法是需要约束，约束了接口的方式 Fn(request <T>,response<*T>)error
	rpc.RegisterName("HelloService", &HelloService{})

	//	已经注册需要报楼的RPC到RPC框架内，需要设置RPC
	//	使用net包(TCP/UDP),来进行网络通信
	// 然后我们建立一个唯一的TCP链接
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	// 通过rpc.ServeConn函数在该TCP链接上为对方提供RPC服务。
	// 没Accept一个请求，就创建一个goroutie进行处理
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		//把network交给RPC框架
		//buf := bufio.NewWriter(conn)
		//srv := &gobServerCodec{
		//	rwc:    conn,
		//	dec:    gob.NewDecoder(conn),
		//	enc:    gob.NewEncoder(buf),
		//	encBuf: buf,
		//}
		//server.serverCodec(srv)
		//rpc.ServerCodec(svc)

		//rpc框架为我们提供了基于 json 服务端的的编解码器 jsonrpc.ServerCodec
		svc:=jsonrpc.NewServerCodec(conn)
		go rpc.ServeCodec(svc)
	}

}
