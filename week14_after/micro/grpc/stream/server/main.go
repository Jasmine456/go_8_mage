package main

import (
	"fmt"
	"github.com/infraboard/mcube/logger/zap"
	"go_8_mage/week14_after/micro/rpc/grpc/stream/server/auth"
	"go_8_mage/week14_after/micro/rpc/grpc/stream/pb"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

func main() {
	//初始化全局Logger实例
	zap.DevelopmentSetup()
	//	如何把我们实现了接口的对象HelloServiceServerImpl 注册给grpc框架
	// 1.首先是通过grpc.NewServer()构造一个gRPC服务对象
	//1.1 流拦截器

	grpcServer := grpc.NewServer(grpc.ChainStreamInterceptor(auth.NewStreamServerInterceptorImpl().Auth))

	//	2.把实现了接口的对象注册给grpc server
	// go-grpc插件，帮忙生成了对象的注册方法
	//把 HelloServiceServer的实现 HelloServiceServerImpl 注册给grpcServer
	pb.RegisterHelloServiceServer(grpcServer, &helloServiceChannelServerImpl{})

	//	3. 启动grpc服务器，grpcServer是建立在HTTP2网络层之上的
	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}

//写一个接口的实现类
type helloServiceChannelServerImpl struct {
	//嵌套一个标准继承来实现
	pb.UnimplementedHelloServiceServer
}

//实现流式响应
func (i *helloServiceChannelServerImpl) Channel(stream pb.HelloService_ChannelServer) error {
	//1.接受客户请求（也可以不处理）
	for {
		req, err := stream.Recv()
		//	如果遇到io.EOF表示客户端流被关闭
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	//	2。处理用户响应
	err=stream.Send(&pb.Response{Value: fmt.Sprintf("hello,%s",req.Value)})
		if err != nil {
			return err
		}
	}

	return nil
}
