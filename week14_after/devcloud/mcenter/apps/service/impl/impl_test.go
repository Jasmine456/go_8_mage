package impl_test

import (
	"context"

	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/service"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/conf"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger/zap"
	"testing"

	//注册所有服务
	_ "github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/all"
)

var (
	svr service.MetaService
	ctx = context.Background()
)

func TestCreateService(t *testing.T) {
	s:=service.NewCreateServiceRequest()
	s.Name="maudit"
	s.Description="微服务审计中心"

	ins,err:=svr.CreateService(ctx,s)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)

}

func init() {
	//初始化日志示例
	zap.DevelopmentSetup()

	//	初始化配置,提前配置好 etc/unit_env.toml
	err := conf.LoadConfigFromEnv()
	if err != nil {
		panic(err)
	}

	//	初始化全局app
	if err:=app.InitAllApp();err!=nil{
		panic(err)
	}

	svr=app.GetInternalApp(service.AppName).(service.MetaService)

}
