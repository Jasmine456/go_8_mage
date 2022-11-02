package impl_test

import (
	"context"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/domain"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/user"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/test/tools"
	"testing"

	"github.com/infraboard/mcube/app"
)

var (
	impl user.Service
	ctx  = context.Background()
)

func TestCreateUser(t *testing.T) {
	req := user.NewCreateUserRequest()
	// 子账号: 审核管理员
	req.Domain = domain.DEFAULT_DOMAIN
	req.Username = "maudit_admin"
	req.Password = "123456"
	req.Type = user.TYPE_SUB
	r, err := impl.CreateUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}

func init() {
	tools.DevelopmentSetup()
	impl = app.GetInternalApp(user.AppName).(user.Service)
}
