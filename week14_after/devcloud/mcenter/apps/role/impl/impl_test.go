package impl

import (
	"context"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/domain"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/role"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/test/tools"
	"github.com/infraboard/mcube/app"
	"testing"
)

var (
	scvr role.Service
	ctx  = context.Background()
)

func TestCreateRole(t *testing.T) {
	req := role.NewCreateRoleRequest()
	req.CreateBy = "test"
	req.Domain = domain.DEFAULT_DOMAIN
	req.Type = role.RoleType_GLOBAL
	req.Name = "madut_admin"
	req.Description = "审计中心管理员"
	req.Specs = []*role.Spec{
		{
			Desc:   "单元测试测试",
			Effect: role.EffectType_ALLOW,
			//maudit_sevice_id
			ServiceId:    "2a4e174e",
			ResourceName: "book",
			LabelKey:     "action",
			LabelValues:  []string{"create", "list"},
		},
	}

	r, err := scvr.CreateRole(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}

func TestQueryRole(t *testing.T) {
	req := role.NewQueryRoleRequest()
	req.WithPermission = true
	set, err := scvr.QueryRole(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

func TestDescribeRole(t *testing.T) {
	req := role.NewDescribeRoleRequestWithID("cddnivbfqkr8ps1bdsdg")
	req.WithPermissions = true
	r, err := scvr.DescribeRole(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}

func init() {
	tools.DevelopmentSetup()
	scvr = app.GetInternalApp(role.AppName).(role.Service)

}
