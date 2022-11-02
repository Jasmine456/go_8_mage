package impl

import (
	"context"
	"fmt"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/endpoint"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/permission"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/policy"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/role"

	"github.com/infraboard/mcube/exception"
)

func (s *service) QueryPermission(ctx context.Context, req *permission.QueryPermissionRequest) (
	*role.PermissionSet, error) {
	return nil, nil
	// if err := req.Validate(); err != nil {
	// 	return nil, exception.NewBadRequest("validate param error, %s", err)
	// }

	// // 获取用户的策略列表
	// preq := policy.NewQueryPolicyRequest()
	// preq.Page = request.NewPageRequest(100, 1)
	// preq.Username = req.Username
	// preq.Namespace = req.Namespace

	// policySet, err := s.policy.QueryPolicy(ctx, preq)
	// if err != nil {
	// 	return nil, err
	// }

	// // 获取用户的角色列表
	// rset, err := policySet.GetRoles(ctx, s.role, true)
	// if err != nil {
	// 	return nil, err
	// }

	// return rset.Permissions(), nil
}

func (s *service) QueryRole(ctx context.Context, req *permission.QueryRoleRequest) (
	*role.RoleSet, error) {
	return nil, nil
	// if err := req.Validate(); err != nil {
	// 	return nil, exception.NewBadRequest("validate param error, %s", err)
	// }

	// // 获取用户的策略列表
	// preq := policy.NewQueryPolicyRequest()
	// preq.Page = request.NewPageRequest(100, 1)
	// preq.Username = req.Username
	// preq.Namespace = req.Namespace

	// policySet, err := s.policy.QueryPolicy(ctx, preq)
	// if err != nil {
	// 	return nil, err
	// }

	// return policySet.GetRoles(ctx, s.role, req.WithPermission)
}

func (s *service) CheckPermission(ctx context.Context, req *permission.CheckPermissionRequest) (*role.Permission, error) {
	// 1.  用户在namespace下关联的策略
	policyReq := policy.NewQueryPolicyRequest()
	policyReq.Domain = req.Domain
	policyReq.Namespace = req.Namespace
	policyReq.Username = req.Username
	policyReq.WithRole = true
	policySet, err := s.policy.QueryPolicy(ctx, policyReq)
	if err != nil {
		return nil, err
	}

	// TODO: 判断策略是否由有效，是否过期
	// 获取用户的角色列表
	rset, err := policySet.GetRoles(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Println(policySet)

	// 判断这些角色是否包含 Endpoint
	// 根据 serviceId和Path 查询Endpoint
	descReq := &endpoint.DescribeEndpointRequest{
		Id: endpoint.GenHashID(req.ServiceId, req.Path),
	}
	ep, err := s.endpoint.DescribeEndpoint(ctx, descReq)
	if err != nil {
		return nil, err
	}

	perm := rset.HasPerm(ep)
	if perm != nil {
		return perm, nil
	}

	return nil, exception.NewPermissionDeny("no permission")
}
