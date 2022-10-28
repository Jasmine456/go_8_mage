package rpc

import (
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/endpoint"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/service"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/token"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// NewClient todo
func NewClient(conf *Config) (*ClientSet, error) {
	zap.DevelopmentSetup()
	log := zap.L()

	conn, err := grpc.Dial(
		conf.Address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithPerRPCCredentials(conf.Credentials()),
	)
	if err != nil {
		return nil, err
	}

	return &ClientSet{
		conn: conn,
		log:  log,
		conf: conf,
	}, nil
}

// Client 客户端
type ClientSet struct {
	conn *grpc.ClientConn
	log  logger.Logger
	conf *Config
}

// Token服务的SDK
func (c *ClientSet) Token() token.RPCClient {
	return token.NewRPCClient(c.conn)
}

// Endpoint 服务的SDK
func (c *ClientSet) Endpoint() endpoint.RPCClient {
	return endpoint.NewRPCClient(c.conn)
}

// 如何注册，获取到当前所有的路由定义，把这些路由定义转化为Endpoint，然后进行注册
// Gin 不支持
// GoRestful支持

// Service 内部服务管理
func (c *ClientSet) Service() service.RPCClient {
	return service.NewRPCClient(c.conn)
}

//这个地方有安全风险
func (c *ClientSet)Config() Config{
	// c.conf.ClientSecret 是有风险的
	// ？什么情况下 client secret会泄露
	return *c.conf
}