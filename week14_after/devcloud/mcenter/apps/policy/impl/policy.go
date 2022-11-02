package impl

import (
	"context"
	"fmt"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/policy"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/role"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/user"

	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *impl) CreatePolicy(ctx context.Context, req *policy.CreatePolicyRequest) (
	*policy.Policy, error) {
	ins, err := policy.New(req)
	if err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	u, err := s.CheckDependence(ctx, ins)
	if err != nil {
		return nil, err
	}
	s.log.Debugf("user: %s", u.Spec.Username)

	if _, err := s.col.InsertOne(ctx, ins); err != nil {
		return nil, exception.NewInternalServerError("inserted policy(%s) document error, %s",
			ins.Id, err)
	}

	return ins, nil
}

// CheckDependence todo
func (i *impl) CheckDependence(ctx context.Context, p *policy.Policy) (*user.User, error) {
	account, err := i.user.DescribeUser(ctx, user.NewDescriptUserRequestWithName(p.Spec.Username))
	if err != nil {
		return nil, fmt.Errorf("check user error, %s", err)
	}

	_, err = i.role.DescribeRole(ctx, role.NewDescribeRoleRequestWithID(p.Spec.RoleId))
	if err != nil {
		return nil, fmt.Errorf("check role error, %s", err)
	}

	return account, nil
}

//QueryPolicy
func (s *impl) QueryPolicy(ctx context.Context, req *policy.QueryPolicyRequest) (
	*policy.PolicySet, error) {
	r, err := newQueryPolicyRequest(req)
	if err != nil {
		return nil, err
	}

	s.log.Debugf("query policy filter: %s", r.FindFilter())

	resp, err := s.col.Find(ctx, r.FindFilter(), r.FindOptions())
	if err != nil {
		return nil, exception.NewInternalServerError("find policy error, error is %s", err)
	}

	set := policy.NewPolicySet()

	// 循环
	for resp.Next(ctx) {

		ins := policy.NewDefaultPolicy()
		if err := resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode policy error, error is %s", err)
		}
		// 补充关联的角色信息
		if req.WithRole {
			descRole := role.NewDescribeRoleRequestWithID(ins.Spec.RoleId)
			descRole.WithPermissions = true
			ins.Role, err = s.role.DescribeRole(ctx, descRole)
			if err != nil {
				return nil, err
			}
		}

		set.Add(ins)
	}

	// count
	count, err := s.col.CountDocuments(ctx, r.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get policy count error, error is %s", err)
	}
	set.Total = count

	return set, nil
}

func (s *impl) DescribePolicy(ctx context.Context, req *policy.DescribePolicyRequest) (
	*policy.Policy, error) {
	r, err := newDescribePolicyRequest(req)
	if err != nil {
		return nil, err
	}

	ins := policy.NewDefaultPolicy()
	s.log.Debugf("describe policy filter: %s", r.FindFilter())
	if err := s.col.FindOne(ctx, r.FindFilter()).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("policy %s not found", req)
		}

		return nil, exception.NewInternalServerError("find policy %s error, %s", req.Id, err)
	}

	return ins, nil
}

func (s *impl) DeletePolicy(ctx context.Context, req *policy.DeletePolicyRequest) (*policy.Policy, error) {
	descReq := policy.NewDescriptPolicyRequest()
	descReq.Id = req.Id
	p, err := s.DescribePolicy(ctx, descReq)
	if err != nil {
		return nil, err
	}

	r, err := newDeletePolicyRequest(req)
	if err != nil {
		return nil, err
	}

	_, err = s.col.DeleteOne(ctx, r.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("delete policy(%s) error, %s", req.Id, err)
	}

	return p, nil
}
