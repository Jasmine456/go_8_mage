package impl

import (
	"context"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/endpoint"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/service"

	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *impl) RegistryEndpoint(ctx context.Context, req *endpoint.RegistryRequest) (*endpoint.RegistryResponse, error) {
	// 查询该服务
	svr, err := s.app.DescribeService(ctx, service.NewDescribeServiceRequestByClientId(req.ClientId))
	if err != nil {
		return nil, err
	}
	s.log.Debugf("service %s registry endpoints", svr.Spec.Name)

	// 冗余逻辑, 避免中间件没有加认证
	if err := svr.Credential.Validate(req.ClientSecret); err != nil {
		return nil, err
	}

	// 生产该服务的Endpoint
	endpoints := req.Endpoints(svr.Id)

	// 更新已有的记录
	news := make([]interface{}, 0, len(endpoints))
	for i := range endpoints {
		if err := s.col.FindOneAndReplace(context.TODO(), bson.M{"_id": endpoints[i].Id}, endpoints[i]).Err(); err != nil {
			if err == mongo.ErrNoDocuments {
				news = append(news, endpoints[i])
			} else {
				return nil, err
			}
		}
	}

	// 插入新增记录
	if len(news) > 0 {
		if _, err := s.col.InsertMany(context.TODO(), news); err != nil {
			return nil, exception.NewInternalServerError("inserted a service document error, %s", err)
		}
	}

	return endpoint.NewRegistryResponse("ok"), nil
}

func (s *impl) DescribeEndpoint(ctx context.Context, req *endpoint.DescribeEndpointRequest) (
	*endpoint.Endpoint, error) {
	r, err := newDescribeEndpointRequest(req)
	if err != nil {
		return nil, err
	}

	ins := endpoint.NewDefaultEndpoint()
	if err := s.col.FindOne(context.TODO(), r.FindFilter()).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("tools %s not found", req)
		}

		return nil, exception.NewInternalServerError("find tools %s error, %s", req.Id, err)
	}

	return ins, nil
}

func (s *impl) QueryEndpoints(ctx context.Context, req *endpoint.QueryEndpointRequest) (
	*endpoint.EndpointSet, error) {
	r := newQueryEndpointRequest(req)
	resp, err := s.col.Find(context.TODO(), r.FindFilter(), r.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find tools error, error is %s", err)
	}

	set := endpoint.NewEndpointSet()
	// 循环
	for resp.Next(context.TODO()) {
		app := endpoint.NewDefaultEndpoint()
		if err := resp.Decode(app); err != nil {
			return nil, exception.NewInternalServerError("decode domain error, error is %s", err)
		}

		set.Add(app)
	}

	// count
	count, err := s.col.CountDocuments(context.TODO(), r.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get device count error, error is %s", err)
	}
	set.Total = count

	return set, nil
}
