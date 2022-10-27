package impl

import (
	"context"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/service"
	"github.com/infraboard/mcube/exception"
)

//实例类的具体方法


// 校验服务的凭证
func (i *impl) ValidateCredential(ctx context.Context, req *service.ValidateCredentialRequest) (
	*service.Service, error) {
	svr, err := i.DescribeService(ctx, service.NewDescribeServiceRequestByClientId(req.ClientId))
	if err != nil {
		return nil, err
	}

	// 判断下 client secret是否正确
	if err := svr.Credential.Validate(req.ClientSecret); err != nil {
		return nil, err
	}

	return svr, nil
}


func (i *impl) DescribeService(ctx context.Context, req *service.DescribeServiceRequest) (
	*service.Service, error) {
	return i.get(ctx, req)
}

func (i *impl) CreateService(ctx context.Context, req *service.CreateServiceRequest) (
	*service.Service, error) {
	ins, err := service.NewService(req)
	if err != nil {
		return nil, exception.NewBadRequest("validate create book error, %s", err)
	}

	if err := i.save(ctx, ins); err != nil {
		return nil, err
	}

	return ins, nil
}
