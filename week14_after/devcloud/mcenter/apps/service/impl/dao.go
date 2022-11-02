package impl

import (
	"context"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/service"
	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// 只用于跟数据库进行交互

// mognodb 数据入库
func (i *impl) save(ctx context.Context, ins *service.Service) error {
	if _, err := i.col.InsertOne(ctx, ins); err != nil {
		return exception.NewInternalServerError("inserted Service(%s) document error, %s",
			ins.Spec.Name, err)
	}
	return nil
}

// 只用于跟数据进行交互
func (i *impl) get(ctx context.Context, req *service.DescribeServiceRequest) (*service.Service, error) {
	filter := bson.M{}
	switch req.DescribeBy {
	case service.DescribeBy_SERVICE_ID:
		filter["_id"] = req.Id
	case service.DescribeBy_SERVICE_CLIENT_ID:
		filter["credential.client_id"] = req.ClientId
	case service.DescribeBy_SERVICE_NAME:
		filter["spec.name"] = req.Name
	}

	ins := service.NewDefaultService()
	if err := i.col.FindOne(ctx, filter).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("Service %s not found", req)
		}

		return nil, exception.NewInternalServerError("find Service %s error, %s", req, err)
	}

	return ins, nil
}
