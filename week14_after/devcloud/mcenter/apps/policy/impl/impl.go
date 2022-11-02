package impl

import (
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/policy"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/role"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/user"
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
	col *mongo.Collection
	log logger.Logger
	policy.UnimplementedRPCServer

	user user.Service
	role role.Service
}

func (i *impl) Config() error {
	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}
	i.col = db.Collection(i.Name())
	i.log = zap.L().Named(i.Name())

	i.user = app.GetInternalApp(user.AppName).(user.Service)
	i.role = app.GetInternalApp(role.AppName).(role.Service)
	return nil
}

func (i *impl) Name() string {
	return policy.AppName
}

func (i *impl) Registry(server *grpc.Server) {
	policy.RegisterRPCServer(server, svr)
}

func init() {
	app.RegistryInternalApp(svr)
	app.RegistryGrpcApp(svr)
}
