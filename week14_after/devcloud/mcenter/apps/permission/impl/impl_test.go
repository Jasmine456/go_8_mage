package impl_test

import (
	"context"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/namespace"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/permission"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/test/tools"
	"github.com/infraboard/mcube/app"
	"testing"
)

var (
	impl permission.Service
	ctx  = context.Background()
)

// 1.创建角色
// 2.创建策略

func TestCheckPermission(t *testing.T) {
	req := permission.NewCheckPermissionRequest()
	req.Username = "maudit_admin"
	req.Namespace = namespace.DEFAULT_NAMESPACE
	req.Group = ""
	req.ServiceId = "2a4e174e"
	req.Path = "GET./maudit/api/v1/book/"
	r, err := impl.CheckPermission(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)

}

func init() {
	tools.DevelopmentSetup()
	impl = app.GetInternalApp(permission.AppName).(permission.Service)
}
