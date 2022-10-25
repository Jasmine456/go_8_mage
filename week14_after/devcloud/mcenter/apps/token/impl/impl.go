package impl

import (
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/go_8_mage/week14_after/devcloud/mcenter/apps/token/provider"
	"google.golang.org/grpc"

	"github.com/go_8_mage/week14_after/devcloud/mcenter/conf"
	"github.com/go_8_mage/week14_after/devcloud/mcenter/apps/token"

	// 加载所有的provider
	_ "github.com/go_8_mage/week14_after/devcloud/mcenter/apps/token/provider/all"
)

var (
	// Service 服务实例
	svr = &service{}
)

type service struct {
	col *mongo.Collection
	log logger.Logger
	token.UnimplementedRPCServer
}

func (s *service) Config() error {

	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}
	s.col = db.Collection(s.Name())

	s.log = zap.L().Named(s.Name())

	// 初始化provider
	return provider.Init()
}

func (s *service) Name() string {
	return token.AppName
}

func (s *service) Registry(server *grpc.Server) {
	token.RegisterRPCServer(server, svr)
}

func init() {
	app.RegistryGrpcApp(svr)
	app.RegistryInternalApp(svr)
}
