package auth

import (
	"context"
	"go_8_mage/week14_after/micro/rpc/grpc/hello/server/auth"
)

//// GetRequestMetadata gets the current request metadata, refreshing
//// tokens if required. This should be called by the transport layer on
//// each request, and the data should be populated in headers or other
//// context. If a status code is returned, it will be used as the status
//// for the RPC. uri is the URI of the entry point for the request.
//// When supported by the underlying implementation, ctx can be used for
//// timeout and cancellation. Additionally, RequestInfo data will be
//// available via ctx to this call.
//// TODO(zhaoq): Define the set of the qualified keys instead of leaving
//// it as an arbitrary string.
//GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error)
//// RequireTransportSecurity indicates whether the credentials requires
//// transport security.
//RequireTransportSecurity() bool
func NewPerRPCCredentialsImpl(clientId,clientSecret string)*PerRPCCredentialsImpl{
	return &PerRPCCredentialsImpl{
		clientId: clientId,
		clientSecret: clientSecret,
	}
}

type PerRPCCredentialsImpl struct {
	clientId string
	clientSecret string
}

//返回需要注入的header
func (i *PerRPCCredentialsImpl)GetRequestMetadata(ctx context.Context,
	uri ...string) (map[string]string, error){
	return map[string]string{
		auth.ClientHeaderKey:i.clientId,
		auth.ClientSecretKey: i.clientSecret,
	},nil
}

func (i *PerRPCCredentialsImpl)RequireTransportSecurity() bool{
	return false
}

