package impl_test

import (
	"context"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/domain"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/namespace"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/policy"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/test/tools"
	"testing"

	"github.com/infraboard/mcube/app"
)

var (
	impl policy.Service
	ctx  = context.Background()
)

func TestCreatePolicy(t *testing.T) {
	req := policy.NewCreatePolicyRequest()
	req.Username = "maudit_admin"
	req.RoleId = "cdh2ff3fqkr04n557j70"
	req.Domain = domain.DEFAULT_DOMAIN
	req.Namespace = namespace.DEFAULT_NAMESPACE
	req.CreateBy = "admin"
	r, err := impl.CreatePolicy(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}

func TestQueryPolicy(t *testing.T) {
	req := policy.NewQueryPolicyRequest()
	req.Domain = domain.DEFAULT_DOMAIN
	req.Namespace = namespace.DEFAULT_NAMESPACE
	req.Username = "maudit_admin"
	req.WithRole = true
	r, err := impl.QueryPolicy(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}

func init() {
	tools.DevelopmentSetup()
	impl = app.GetInternalApp(policy.AppName).(policy.Service)
}
