package main

import (
	"context"
	"fmt"
	"github.com/infraboard/mcube/logger/zap"
	"go_8_mage/week14_after/micro/rpc/grpc/hello/pb"
	"go_8_mage/week14_after/micro/rpc/grpc/hello/server/auth"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	//初始化全局Logger实例
	zap.DevelopmentSetup()


	//	如何把我们实现了接口的对象HelloServiceServerImpl 注册给grpc框架
	// 1.首先是通过grpc.NewServer()构造一个gRPC服务对象
	////1.1添加一个中间件
	//grpc.UnaryServerInterceptor(mid1)
	////1.2 可以
	//grpc.ChainUnaryInterceptor(mid1,mid2,mid3)
	//1.3构造一个添加中间件的服务端参数，实现一个auth中间件

	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(auth.NewServerAuthInceptorImpl().Auth))

	//	2.把实现了接口的对象注册给grpc server
	// go-grpc插件，帮忙生成了对象的注册方法
	//把 HelloServiceServer的实现 HelloServiceServerImpl 注册给grpcServer
	pb.RegisterHelloServiceServer(grpcServer, &HelloServerImpl{})

	//	3. 启动grpc服务器，grpcServer是建立在HTTP2网络层之上的
	lis,err:=net.Listen("tcp",":1234")
	if err != nil {
		log.Fatal(err)
	}
	if err:=grpcServer.Serve(lis);err!=nil{
		panic(err)
	}

}

//1.写一个对象，实现HelloServiceServer
type HelloServerImpl struct {
	//嵌套UnimplementedHelloServiceServer，用于向前兼容（框架）
	//UnimplementedHelloServiceServer该对象有很多默认参数，嵌套后才能实现继承
	pb.UnimplementedHelloServiceServer
}

func (i *HelloServerImpl) Greet(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	return &pb.Response{Value: fmt.Sprintf("hello,%s",req.Value)}, nil
}
