package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"

	"go_8_mage/week14_after/micro/rpc/rpc_protobuf/codec/server"
	"go_8_mage/week14_after/micro/rpc/rpc_protobuf/service"
)

// 如何约束HelloService实现 HelloService接口
// var a1 service.HelloService = &HelloService{}
// 多个一个变量a1 --> _
// 多个一个值 &HelloService{} 多个一个开销, nil

// a1.(service.HelloService)
// (*HelloService)(nil) 表示的是一个*HelloService的nil   (int32)(10)  (*int32)(nil)
var _ service.HelloService = (*HelloService)(nil)

type HelloService struct{}

// hello(request)
// Greet函数的签名: 第一个参数Request 是string, 第二个参数是指针
// alice --->
// alice <-- hello alice
func (s *HelloService) Greet(request *service.Request, resp *service.Response) error {
	// 往指针内 塞入值
	// *resp = "xxxx"
	// StructA{}, 外出RPC框架层 感知不到该对象的变化(值和指针的区别)
	resp.Value = fmt.Sprintf("hello, %s", request.Value)
	return nil
}

func main() {
	// 把又给对象的方法暴露成一个RPC到网络上
	// HelloService: RPC网络上 这个服务的名称, 类似于pkg名称
	// &HelloService{}: 提供RPC方法实现的对象(Fn Greet)
	// 客户端类似于这样来调用  HelloService.Greet()
	//
	// &HelloService{}: 的方法是需要约束, 约束了接口方式 Fn(request <T>, response<*T>) error
	rpc.RegisterName("HelloService", &HelloService{})

	// 已经注册需要暴露的RPC到RPC框架内, 需要设置RPC
	// 使用net包(TCP/UPD), 来进行网络通信
	// 然后我们建立一个唯一的TCP链接，
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	// 不断的从socket 读取报文, 然后处理, 如何处理交接给RPC框架
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
			continue
		}
		// 把network交给RPC框架
		// buf := bufio.NewWriter(conn)
		// srv := &gobServerCodec{
		// 	rwc:    conn,
		// 	dec:    gob.NewDecoder(conn),
		// 	enc:    gob.NewEncoder(buf),
		// 	encBuf: buf,
		// }
		// server.ServeCodec(srv)
		// rcp.ServeCodec(srv)
		// 自己写的基于Protobuf 的 ServerCodec
		svc := server.NewServerCodec(conn)
		go rpc.ServeCodec(svc)
	}
}
