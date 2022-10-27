package impl

import (
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/service"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/conf"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

//实例类

//定义对象
var (
	// Service 服务实例
	svr = &impl{}
)

type impl struct {
	col *mongo.Collection
	log logger.Logger
	service.UnimplementedRPCServer
}

func (i *impl) Config() error {
	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}
	i.col = db.Collection(i.Name())

	i.log = zap.L().Named(i.Name())
	return nil
}

func (i *impl) Name() string {
	return service.AppName
}

func (i *impl) Registry(server *grpc.Server) {
	service.RegisterRPCServer(server, svr)
}

func init() {
	// 通过托管给内部服务, 进程内调用的模块
	app.RegistryInternalApp(svr)
	// 这些模块是需要需要对外提供GRPC服务, 至于那些是对外的GRPC,rpc里面有定义
	app.RegistryGrpcApp(svr)
}

