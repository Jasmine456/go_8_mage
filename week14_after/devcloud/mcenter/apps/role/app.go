package role

import (
	"context"
	"fmt"
	"hash/fnv"
	"time"

	"github.com/go-playground/validator/v10"
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
func NewQueryPermissionRequest(pageReq *request.PageRequest) *QueryPermissionRequest {
	return &QueryPermissionRequest{
		Page: pageReq,
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
