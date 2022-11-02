package policy

import (
	context "context"
	"fmt"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/role"
	"hash/fnv"
	"time"

	"github.com/go-playground/validator/v10"
	request "github.com/infraboard/mcube/http/request"
)

const (
	AppName = "policy"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

type Service interface {
	CreatePolicy(context.Context, *CreatePolicyRequest) (*Policy, error)
	DeletePolicy(context.Context, *DeletePolicyRequest) (*Policy, error)
	RPCServer
}

// New 新实例
func New(req *CreatePolicyRequest) (*Policy, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	p := &Policy{
		CreateAt: time.Now().UnixMilli(),
		UpdateAt: time.Now().UnixMilli(),
		Spec:     req,
	}
	p.genID()

	return p, nil
}

// Validate 校验请求合法
func (req *CreatePolicyRequest) Validate() error {
	return validate.Struct(req)
}

// 保证策略的唯一性
func (p *Policy) genID() {
	// 短hash算法，生成比较短的hash串
	h := fnv.New32a()
	hashedStr := fmt.Sprintf("%s-%s-%s-%s-%s",
		p.Spec.Domain, p.Spec.Namespace, p.Spec.Group, p.Spec.Username, p.Spec.RoleId)

	h.Write([]byte(hashedStr))
	p.Id = fmt.Sprintf("%x", h.Sum32())
}

// Validate 校验请求是否合法
func (req *QueryPolicyRequest) Validate() error {
	return validate.Struct(req)
}

// NewPolicySet todo
func NewPolicySet() *PolicySet {
	return &PolicySet{
		Items: []*Policy{},
	}
}

// NewDefaultPolicy todo
func NewDefaultPolicy() *Policy {
	return &Policy{
		Spec: NewCreatePolicyRequest(),
	}
}

// Add 添加
func (s *PolicySet) Add(e *Policy) {
	s.Items = append(s.Items, e)
}

// Validate todo
func (req *DescribePolicyRequest) Validate() error {
	if req.Id == "" {
		return fmt.Errorf("policy id required")
	}

	return nil
}

// NewDescriptPolicyRequest new实例
func NewDescriptPolicyRequest() *DescribePolicyRequest {
	return &DescribePolicyRequest{}
}

// NewDeletePolicyRequestWithID todo
func NewDeletePolicyRequestWithID(id string) *DeletePolicyRequest {
	return &DeletePolicyRequest{Id: id}
}

// NewCreatePolicyRequest 请求实例
func NewCreatePolicyRequest() *CreatePolicyRequest {
	return &CreatePolicyRequest{}
}

// NewQueryPolicyRequest 列表查询请求
func NewQueryPolicyRequest() *QueryPolicyRequest {
	return &QueryPolicyRequest{
		Page:          request.NewDefaultPageRequest(),
		WithRole:      false,
		WithNamespace: false,
	}
}

// GetRoles todo
func (s *PolicySet) GetRoles(ctx context.Context) (*role.RoleSet, error) {
	set := role.NewRoleSet()
	for i := range s.Items {
		set.Add(s.Items[i].Role)
	}
	return set, nil
}
