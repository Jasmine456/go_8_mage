package auth

import (
	"fmt"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

//// StreamServerInterceptor provides a hook to intercept the execution of a streaming RPC on the server.
//// info contains all the information of this RPC the interceptor can operate on. And handler is the
//// service method implementation. It is the responsibility of the interceptor to invoke handler to
//// complete the RPC.
//type StreamServerInterceptor func(srv interface{}, ss ServerStream, info *StreamServerInfo, handler StreamHandler) error
const (
	ClientHeaderKey = "client-id"
	ClientSecretKey = "client-secret"
)


//从header中获取认证信息
func (i *StreamServerInterceptorImpl) GetClientCredentialsFromMeta(md metadata.MD) (
	clientId, clientSecret string) {
	cids := md.Get(ClientHeaderKey)
	sids := md.Get(ClientSecretKey)
	if len(cids) > 0 {
		clientId = cids[0]
	}
	if len(sids) > 0 {
		clientSecret = sids[0]
	}
	return
}

func NewStreamServerInterceptorImpl()*StreamServerInterceptorImpl{
	return &StreamServerInterceptorImpl{
		//需要提前加载
		//zap.DevelopmentSetup() 初始化全局Logger示例，这个示例通过该方法获取：Zap.L()
		log:zap.L().Named("middleware.auth"),
	}
}

type StreamServerInterceptorImpl struct {
	log logger.Logger
}

func (i *StreamServerInterceptorImpl)Auth(srv interface{}, ss grpc.ServerStream,
info *grpc.StreamServerInfo, handler grpc.StreamHandler) error{

	//如何获取metadata,通过ServerStream
	ctx:=ss.Context()
	//grpc header，HTTP2是有header，这个header在ctx
	// 从上下文中获取认证信息，这里的md 就是类似于header
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return fmt.Errorf("ctx is not an grpc incoming context")
	}

	//认证逻辑
	cid,cs:=i.GetClientCredentialsFromMeta(md)
	i.log.Debug(cid,cs)
	if cid !="admin"|| cs!="123456"{
		return grpc.Errorf(codes.Unauthenticated,"客户端调用凭证不正确")
	}



	////请求路由到下一个环节
	//res,err:=handler(ctx,req)

	//// 处理响应后的请求
	//i.log.Debugf("resp:%s",res)

	return handler(srv,ss)
}