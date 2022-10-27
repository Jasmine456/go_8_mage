package rpc

import (
	"context"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/service"

	"google.golang.org/grpc/metadata"
)

// 写grpc的客户端如何携带认证信息
// grpc.WithPerRPCCredentials, 要求实现了
// PerRPCCredentials defines the common interface for the credentials which need to
// attach security information to every RPC (e.g., oauth2).
// type PerRPCCredentials interface {
// 	// GetRequestMetadata gets the current request metadata, refreshing
// 	// tokens if required. This should be called by the transport layer on
// 	// each request, and the data should be populated in headers or other
// 	// context. If a status code is returned, it will be used as the status
// 	// for the RPC. uri is the URI of the entry point for the request.
// 	// When supported by the underlying implementation, ctx can be used for
// 	// timeout and cancellation. Additionally, RequestInfo data will be
// 	// available via ctx to this call.
// 	// TODO(zhaoq): Define the set of the qualified keys instead of leaving
// 	// it as an arbitrary string.
// 	GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error)
// 	// RequireTransportSecurity indicates whether the credentials requires
// 	// transport security.
// 	RequireTransportSecurity() bool
// }

// 客户端携带的凭证
func NewAuthentication(clientId, clientSecret string) *Authentication {
	return &Authentication{
		clientID:     clientId,
		clientSecret: clientSecret,
	}
}

// Authentication todo
type Authentication struct {
	clientID     string
	clientSecret string
}

// SetClientCredentials todo
func (a *Authentication) SetClientCredentials(clientID, clientSecret string) {
	a.clientID = clientID
	a.clientSecret = clientSecret
}

// GetRequestMetadata todo
func (a *Authentication) GetRequestMetadata(context.Context, ...string) (
	map[string]string, error,
) {
	return map[string]string{
		service.ClientHeaderKey: a.clientID,
		service.ClientSecretKey: a.clientSecret,
	}, nil
}

// RequireTransportSecurity todo
func (a *Authentication) RequireTransportSecurity() bool {
	return false
}

func GetClientCredential(ctx context.Context) (clientId, clientSecret string) {
	// 重上下文中获取认证信息
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return
	}

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
