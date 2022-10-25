package main

import (
	"context"
	"fmt"
	"go_8_mage/week14_after/micro/rpc/grpc/hello/client/auth"
	"go_8_mage/week14_after/micro/rpc/grpc/stream/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"time"
)

//其他服务器上的客户端需要来调用grpc
func main() {
	//	1.建立网络链接,默认情况下 HTTP2 是强制要求证书的
	//	grpc.WithInsecure已经降级
	//	conn,err:=grpc.Dial("localhost:1234",grpc.WithInsecure())
	conn, err := grpc.Dial(
		"localhost:1234",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithPerRPCCredentials(auth.NewPerRPCCredentialsImpl("admin","123456")),
		)
	if err != nil {
		panic(err)
	}
	//	1. go-grpc这个插件已经生成好了grpc的client sdk
	client := pb.NewHelloServiceClient(conn)

	//3.使用sdk
	stream, err := client.Channel(context.Background())
	if err != nil {
		panic(err)
	}
	////请求的发送
	//stream.Send(nil)
	////请求的接收
	//stream.Recv()

	//	3.1 专门启动一个goroutine 用于发送请求
	go func() {
		for {
			err := stream.Send(&pb.Request{Value: "bob"})
			if err != nil {
				log.Fatal(err)
			}
			//	休息1s
			time.Sleep(1 * time.Second)
		}
	}()

	//	3.2 处理服务端响应
	for {
		resp,err:=stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		fmt.Println(resp)

	}
}
