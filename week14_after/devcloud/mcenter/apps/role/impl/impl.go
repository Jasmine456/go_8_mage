package impl

import (
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/role"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/conf"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

var (
	// Service 服务实例
	svr = &impl{}
)

type impl struct {
	role *mongo.Collection
	perm *mongo.Collection
	log  logger.Logger
	role.UnimplementedRPCServer
}

func (i *impl) Config() error {
	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}
	i.role = db.Collection("role")
	i.perm = db.Collection("permission")

	i.log = zap.L().Named(i.Name())
	return nil
}

func (i *impl) Name() string {
	return role.AppName
}

func (i *impl) Registry(server *grpc.Server) {
	role.RegisterRPCServer(server, svr)
}

func init() {
	app.RegistryInternalApp(svr)
	app.RegistryGrpcApp(svr)
}
