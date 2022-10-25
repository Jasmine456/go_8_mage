package main

import (
	"context"
	"fmt"
	"go_8_mage/week14_after/micro/rpc/grpc/hello/client/auth"
	"go_8_mage/week14_after/micro/rpc/grpc/hello/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

//其他服务器上的客户端需要来调用grpc
func main(){
//	1.建立网络链接,默认情况下 HTTP2 是强制要求证书的
//	grpc.WithInsecure已经降级
//	conn,err:=grpc.Dial("localhost:1234",grpc.WithInsecure())

	//1.1 客户端如何添加header，在header认证信息
	//1.2客户端认证中间件参数grpc.WithTransportCredentials(),需要实现接口：credentials.PerRPCCredentials


	conn,err:=grpc.Dial(
		"localhost:1234",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithPerRPCCredentials(auth.NewPerRPCCredentialsImpl("admin","123456")),
		)
	if err != nil {
		panic(err)
	}
//	1. go-grpc这个插件已经生成好了grpc的client sdk
	client:=pb.NewHelloServiceClient(conn)

	//3.使用sdk
	//header:=metadata.New(map[string]string{
	//	auth.ClientHeaderKey:"admin",
	//	auth.ClientSecretKey: "123456",
	//})
	resp,err:=client.Greet(context.Background(),&pb.Request{Value: "alice"})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)

}
