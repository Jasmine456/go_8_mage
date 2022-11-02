package auth

import (
	"context"
	"fmt"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/service"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/client/rpc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// grpc 认证中间件

// GrpcAuthUnaryServerInterceptor returns a new unary server interceptor for auth.
func GrpcAuthUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	// 替换成rpc客户端，rpc的客户端怎么获取？
	// 1. 直接传入需要传入grpc客户端？
	// 2. mcenter服务一定会出事户的？ 为什么？ （需要注册 endpoint，使用rpc客户端）
	// 3. 使用全局变量，把rpc的客户端 作为一个全局变量，提供一个Load加载韩束，加载完成后，直接通过rpc.() 返回获取该全局变量（rpc客户端实力）
	// 4.
	return newGrpcAuther(rpc.C().Service()).Auth
}

func newGrpcAuther(svr service.RPCClient) *grpcAuther {
	return &grpcAuther{
		log:     zap.L().Named("Grpc Auther"),
		service: svr,
	}
}

// internal todo
type grpcAuther struct {
	log     logger.Logger
	service service.RPCClient
}

// grpc 中间件签名函数
func (a *grpcAuther) Auth(
	ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	// 重上下文中获取认证信息
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("ctx is not an grpc incoming context")
	}

	clientId, clientSecret := a.GetClientCredentialsFromMeta(md)

	// 校验调用的客户端凭证是否有效
	if err := a.validateServiceCredential(clientId, clientSecret); err != nil {
		return nil, err
	}

	resp, err = handler(ctx, req)
	return resp, err
}

func (a *grpcAuther) GetClientCredentialsFromMeta(md metadata.MD) (
	clientId, clientSecret string) {
	cids := md.Get(service.ClientHeaderKey)
	sids := md.Get(service.ClientSecretKey)
	if len(cids) > 0 {
		clientId = cids[0]
	}
	if len(sids) > 0 {
		clientSecret = sids[0]
	}
	return
}

func (a *grpcAuther) validateServiceCredential(clientId, clientSecret string) error {
	if clientId == "" && clientSecret == "" {
		return status.Errorf(codes.Unauthenticated, "client_id or client_secret is \"\"")
	}

	vsReq := service.NewValidateCredentialRequest(clientId, clientSecret)
	_, err := a.service.ValidateCredential(context.Background(), vsReq)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "service auth error, %s", err)
	}

	return nil
}
