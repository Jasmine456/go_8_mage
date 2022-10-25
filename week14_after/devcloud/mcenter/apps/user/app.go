package user

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/request"
	"github.com/rs/xid"
	"github.com/go_8_mage/week14_after/devcloud/mcenter/apps/domain"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
	"time"
)

const (
	AppName = "user"
)

//模块内部接口
type Service interface {
	//   创建用户
	CreateUser(context.Context,  *CreateUserRequest) (*User, error)
	//   更新用户
	UpdateUser(context.Context, *UpdateUserRequest) (*User, error)
	//    删除用户
	DeleteUser(context.Context, *DeleteUserRequest) (*User, error)
	// 修改用户密码, 用户需要知道原先密码
	UpdatePassword(context.Context, *UpdatePasswordRequest) (*Password, error)
	// 重置密码, 无需知道原先密码, 主账号执行
	ResetPassword(context.Context, *ResetPasswordRequest) (*Password, error)
	// rpc接口
	RPCServer
}

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

func NewCreateUserRequest() *CreateUserRequest {
	return &CreateUserRequest{
		Domain: domain.DEFAULT_DOMAIN,
	}
}

// Validate 校验请求是否合法
func (req *CreateUserRequest) Validate() error {
	return validate.Struct(req)
}

// New 实例
func New(req *CreateUserRequest) (*User, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	// 不能保存明文的密码, Hash
	pass, err := NewHashedPassword(req.Password)
	if err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	u := &User{
		// 随机的id
		Id:            xid.New().String(),
		CreateAt:      time.Now().UnixMilli(),
		Spec:          req,
		Password:      pass,
		Profile:       &Profile{},
		IsInitialized: false,
		Status: &Status{
			Locked: false,
		},
	}

	return u, nil
}

// NewHashedPassword 生产hash后的密码对象
// bcrypt专门用于密码hash的 散列算法
func NewHashedPassword(password string) (*Password, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil, err
	}

	return &Password{
		Password:      string(bytes),
		CreateAt:      time.Now().UnixMilli(),
		UpdateAt:      time.Now().UnixMilli(),
		ExpiredDays:   90,
		ExpiredRemind: 30,
	}, nil
}

// SetNeedReset 需要被重置
func (p *Password) SetNeedReset(format string, a ...interface{}) {
	p.NeedReset = true
	p.ResetReason = fmt.Sprintf(format, a...)
}

func NewUserSet() *UserSet {
	return &UserSet{
		Items: []*User{},
	}
}

func NewDefaultUser() *User {
	return &User{}
}

// Desensitize 关键数据脱敏
func (u *User) Desensitize() {
	if u.Password != nil {
		u.Password.Password = ""
		u.Password.History = []string{}
	}
}

func (s *UserSet) Add(item *User) {
	s.Items = append(s.Items, item)
}

func (req *QueryUserRequest) WithType(t TYPE) {
	req.Type = &t
}

// NewQueryUserRequestFromHTTP todo
func NewQueryUserRequestFromHTTP(r *http.Request) *QueryUserRequest {
	query := NewQueryUserRequest()

	// url 参数  ?a=b&c=d
	qs := r.URL.Query()
	query.Page = request.NewPageRequestFromHTTP(r)
	query.Keywords = qs.Get("keywords")
	query.SkipItems = qs.Get("skip_items") == "true"

	uids := qs.Get("user_ids")
	if uids != "" {
		query.UserIds = strings.Split(uids, ",")
	}
	return query
}

// NewQueryUserRequest 列表查询请求
func NewQueryUserRequest() *QueryUserRequest {
	return &QueryUserRequest{
		Page:      request.NewPageRequest(20, 1),
		SkipItems: false,
	}
}

// NewDescriptUserRequestWithId 查询详情请求
func NewDescriptUserRequestWithId(id string) *DescribeUserRequest {
	return &DescribeUserRequest{
		DescribeBy: DESCRIBE_BY_USER_ID,
		Id:         id,
	}
}

// NewDescriptUserRequestWithId 查询详情请求
func NewDescriptUserRequestWithName(username string) *DescribeUserRequest {
	return &DescribeUserRequest{
		DescribeBy: DESCRIBE_BY_USER_NAME,
		Username:   username,
	}
}

// CheckPassword 判断password 是否正确, 是明文的
func (p *Password) CheckPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(p.Password), []byte(password))
	if err != nil {
		return exception.NewUnauthorized("用户名或者密码错误")
	}

	return nil
}
