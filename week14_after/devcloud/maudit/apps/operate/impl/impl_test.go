package impl

import (
	"context"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/maudit/apps/operate"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/maudit/test/tools"
	"github.com/infraboard/mcube/app"
	"testing"
)

var (
	impl operate.Service
	ctx  = context.Background()
)

func TestSaveOperateLog(t *testing.T) {
	req := operate.NewOperateLog()
	req.Domain = "default"
	req.Namespace="default"
	req.Username="username"
	req.ResourceType="ecs"
	req.Action="power off"
	resp, err := impl.SaveOperateLog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

func TestQueryOperateLog(t *testing.T) {
	req := operate.NewQueryOperateLogRequest()
	resp, err := impl.QueryOperateLog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

func init() {
	tools.DevelopmentSetup()
	impl = app.GetInternalApp(operate.AppName).(operate.Service)
}
