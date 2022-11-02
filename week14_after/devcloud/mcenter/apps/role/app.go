package role

import (
	"context"
	"fmt"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/endpoint"
	"hash/fnv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/exception"
	request "github.com/infraboard/mcube/http/request"
	"github.com/rs/xid"
)

const (
	AppName = "role"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

type Service interface {
	CreateRole(context.Context, *CreateRoleRequest) (*Role, error)
	DeleteRole(context.Context, *DeleteRoleRequest) (*Role, error)
	AddPermissionToRole(context.Context, *AddPermissionToRoleRequest) (*PermissionSet, error)
	RemovePermissionFromRole(context.Context, *RemovePermissionFromRoleRequest) (*PermissionSet, error)
	UpdatePermission(context.Context, *UpdatePermissionRequest) (*Permission, error)
	RPCServer
}

// Validate 请求校验
func (req *CreateRoleRequest) Validate() error {
	return validate.Struct(req)
}

// New 新创建一个Role
func New(req *CreateRoleRequest) (*Role, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	r := &Role{
		Id:       xid.New().String(),
		CreateAt: time.Now().UnixMilli(),
		UpdateAt: time.Now().UnixMilli(),
		Spec:     req,
	}

	return r, nil
}

// Validate todo
func (req *DescribeRoleRequest) Validate() error {
	if req.Id == "" && req.Name == "" {
		return fmt.Errorf("id or name required")
	}

	return nil
}

// Validate todo
func (req *QueryRoleRequest) Validate() error {
	return nil
}

// NewRoleSet 实例化make
func NewRoleSet() *RoleSet {
	return &RoleSet{
		Items: []*Role{},
	}
}

// Add todo
func (s *RoleSet) Add(item *Role) {
	s.Total++
	s.Items = append(s.Items, item)
}

func (s *RoleSet) HasPerm(ed *endpoint.Endpoint) *Permission {
	for i := range s.Items {
		perm := s.Items[i].HasPerm(ed)
		if perm != nil {
			return perm
		}
	}

	return nil
}

// 判断该角色是否具有权限A
func (r *Role) HasPerm(ed *endpoint.Endpoint) *Permission {
	for i := range r.Permissions {
		perm := r.Permissions[i]
		if perm.HasPerm(ed) {
			return perm
		}
	}

	return nil
}

func (p *Permission) HasPerm(ed *endpoint.Endpoint) bool {
	if ed == nil {
		return false
	}

	// 开启了鉴权的
	// ed.ServiceId
	// ed.Entry.Resource
	// ed.Entry.Labels
	if ed.Entry.PermissionEnable {
		// 1. service能匹配
		if p.Spec.ServiceId == "*" || p.Spec.ServiceId == ed.ServiceId {
			// 2. resource 匹配
			if p.Spec.ResourceName == "*" || p.Spec.ResourceName == ed.Entry.Resource {
				// 3. labels是否匹配: 判断ep的label 是否匹配 permisson允许的label
				for k, v := range ed.Entry.Labels {
					// 3.1匹配Key
					if p.Spec.LabelKey == "*" || k == p.Spec.LabelKey {
						// 3.2匹配value
						if p.Spec.MatchAll {
							return true
						}
						for _, allow := range p.Spec.LabelValues {
							if allow == v {
								return true
							}
						}
					}
				}
			}
		}
	}

	return false
}

// NewDefaultRole 默认实例
func NewDefaultRole() *Role {
	spec := NewCreateRoleRequest()
	return &Role{
		Spec: spec,
	}
}

// NewCreateRoleRequest 实例化请求
func NewCreateRoleRequest() *CreateRoleRequest {
	return &CreateRoleRequest{
		Type: RoleType_CUSTOM,
		Meta: map[string]string{},
	}
}

// NewDescribeRoleRequestWithID todo
func NewDescribeRoleRequestWithID(id string) *DescribeRoleRequest {
	return &DescribeRoleRequest{
		Id: id,
	}
}

// NewDescribeRoleRequestWithName todo
func NewDescribeRoleRequestWithName(name string) *DescribeRoleRequest {
	return &DescribeRoleRequest{
		Name: name,
	}
}

func (req *RemovePermissionFromRoleRequest) Validate() error {
	return validate.Struct(req)
}

// NewPermissionSet todo
func NewPermissionSet() *PermissionSet {
	return &PermissionSet{
		Items: []*Permission{},
	}
}

// NewRemovePermissionFromRoleRequest todo
func NewRemovePermissionFromRoleRequest() *RemovePermissionFromRoleRequest {
	return &RemovePermissionFromRoleRequest{
		PermissionId: []string{},
	}
}

func (req *AddPermissionToRoleRequest) Validate() error {
	return validate.Struct(req)
}

func (req *QueryPermissionRequest) Validate() error {
	return validate.Struct(req)
}

func NewDeaultPermission() *Permission {
	return &Permission{
		Spec: &Spec{},
	}
}

// Add todo
func (s *PermissionSet) Add(items ...*Permission) {
	s.Items = append(s.Items, items...)
}

// NewQueryPermissionRequest todo
func NewQueryPermissionRequest() *QueryPermissionRequest {
	return &QueryPermissionRequest{
		Page: request.NewDefaultPageRequest(),
	}
}

func (req *AddPermissionToRoleRequest) Length() int {
	return len(req.Permissions)
}

func NewPermission(roleId string, perms []*Spec) []*Permission {
	set := []*Permission{}
	for i := range perms {
		set = append(set, &Permission{
			Id:       perms[i].HashID(roleId),
			CreateAt: time.Now().UnixMilli(),
			RoleId:   roleId,
			Spec:     perms[i],
		})
	}
	return set
}

// cmdb host [create, list]
func (req *Spec) HashID(roleId string) string {
	h := fnv.New32a()

	h.Write([]byte(roleId + req.Effect.String() + req.ServiceId + req.ResourceName))
	return fmt.Sprintf("%x", h.Sum32())
}

func (req *UpdatePermissionRequest) Validate() error {
	if req.Id == "" {
		return exception.NewBadRequest("id required")
	}

	return nil
}

func NewDescribePermissionRequestWithID(id string) *DescribePermissionRequest {
	return &DescribePermissionRequest{Id: id}
}

// NewAddPermissionToRoleRequest todo
func NewAddPermissionToRoleRequest() *AddPermissionToRoleRequest {
	return &AddPermissionToRoleRequest{
		Permissions: []*Spec{},
	}
}

// NewQueryRoleRequest 列表查询请求
func NewQueryRoleRequest() *QueryRoleRequest {
	return &QueryRoleRequest{
		Page: request.NewDefaultPageRequest(),
	}
}
