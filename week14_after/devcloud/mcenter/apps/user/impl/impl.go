package impl

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"

	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/user"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/conf"
)

// 内部服务service接口的实现类

var (
	// Service 服务实例
	svr = &impl{}
)

type impl struct {
	col *mongo.Collection
	log logger.Logger
	//作为rpc服务的一个实现类
	user.UnimplementedRPCServer
}

func (i *impl) Config() error {

	db, err := conf.C().Mongo.GetDB()

	if err != nil {
		return err
	}
	//表的名字，约定为 服务的名字
	i.col = db.Collection(i.Name())

	// 设置唯一建
	indexs := []mongo.IndexModel{
		{
			Keys: bsonx.Doc{{Key: "create_at", Value: bsonx.Int32(-1)}},
		},
		{
			Keys: bsonx.Doc{
				{Key: "spec.domain", Value: bsonx.Int32(-1)},
				{Key: "spec.username", Value: bsonx.Int32(-1)},
			},
			Options: options.Index().SetUnique(true),
		},
	}
	_, err = i.col.Indexes().CreateMany(context.Background(), indexs)
	if err != nil {
		return err
	}

	i.log = zap.L().Named(i.Name())
	return nil
}

func (i *impl) Name() string {
	return user.AppName
}

//把该实例类 作为grpc对外提供服务
func (i *impl) Registry(server *grpc.Server) {
	user.RegisterRPCServer(server, svr)
}

func init() {
	app.RegistryGrpcApp(svr)
	app.RegistryInternalApp(svr)

}
