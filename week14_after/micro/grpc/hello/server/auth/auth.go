package auth

import (
	"context"
	"fmt"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"
	"github.com/infraboard/mcube/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)
const (
	ClientHeaderKey = "client-id"
	ClientSecretKey = "client-secret"
)

func NewServerAuthInceptorImpl()*ServerAuthInceptorImpl{
	return &ServerAuthInceptorImpl{
		//需要提前加载
		//zap.DevelopmentSetup() 初始化全局Logger示例，这个示例通过该方法获取：Zap.L()
		log:zap.L().Named("middleware.auth"),
	}
}

//实现服务端认证的中间件
type ServerAuthInceptorImpl struct {
	log logger.Logger
}

	//从header中获取认证信息
	func (i *ServerAuthInceptorImpl) GetClientCredentialsFromMeta(md metadata.MD) (
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

//认证逻辑
func (i ServerAuthInceptorImpl) Auth(ctx context.Context, req interface{},
info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error){
	i.log.Debugf("req:%s",req)
	i.log.Debugf("server info:server:%s,method:%s",info.Server,info.FullMethod)

	//grpc header，HTTP2是有header，这个header在ctx
	// 从上下文中获取认证信息，这里的md 就是类似于header
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("ctx is not an grpc incoming context")
	}

	//认证逻辑
	cid,cs:=i.GetClientCredentialsFromMeta(md)
	i.log.Debug(cid,cs)
	if cid !="admin"|| cs!="123456"{
		return nil, grpc.Errorf(codes.Unauthenticated,"客户端调用凭证不正确")
	}



	//请求路由到下一个环节
	res,err:=handler(ctx,req)

	// 处理响应后的请求
	i.log.Debugf("resp:%s",res)
	return res,nil
}

