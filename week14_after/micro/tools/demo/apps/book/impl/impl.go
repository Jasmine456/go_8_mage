package impl

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"

	"go_8_mage/week14_after/micro/tools/demo/apps/book"
	"go_8_mage/week14_after/micro/tools/demo/conf"
)

var (
	// Service 服务实例
	svr = &service{}
)

type service struct {
	col *mongo.Collection
	log logger.Logger
	book.UnimplementedServiceServer
}

// 模块依赖管理（config
func (s *service) Config() error {

	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}
	s.col = db.Collection(s.Name())

	s.log = zap.L().Named(s.Name())
	return nil
}

//管理实力类的名称
//说明的说的获取实例类，都通过改名称
func (s *service) Name() string {
	return book.AppName
}

//把实例类注册给grpc server，加载到server内
func (s *service) Registry(server *grpc.Server) {
	book.RegisterServiceServer(server, svr)
}

//把实例类 托管给ioc
func init() {
	app.RegistryGrpcApp(svr)
}
