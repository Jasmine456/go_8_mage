package impl_test

import (
	"context"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/test/tools"

	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/service"
	"github.com/infraboard/mcube/app"
	"testing"

	//注册所有服务
	_ "github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/all"
)

var (
	svr service.MetaService
	ctx = context.Background()
)

func TestCreateService(t *testing.T) {
	s := service.NewCreateServiceRequest()
	s.Name = "maudit"
	s.Description = "微服务审计中心"

	ins, err := svr.CreateService(ctx, s)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)

}

func init() {
	tools.DevelopmentSetup()
	svr = app.GetInternalApp(service.AppName).(service.MetaService)

}
