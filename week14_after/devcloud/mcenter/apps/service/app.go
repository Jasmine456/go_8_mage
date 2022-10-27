package service

import (
	"context"
	"fmt"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/domain"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/namespace"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/token"
	"github.com/go-playground/validator/v10"
	"github.com/rs/xid"
	"hash/fnv"
	"time"
)

const (
	AppName = "service"
)

var (
	validate = validator.New()
)

const (
	ClientHeaderKey = "client-id"
	ClientSecretKey = "client-secret"
)

type MetaService interface {
	CreateService(context.Context, *CreateServiceRequest) (*Service, error)
	RPCServer
}

func NewCreateServiceRequest() *CreateServiceRequest {
	return &CreateServiceRequest{
		Domain:     domain.DEFAULT_DOMAIN,
		Namespace:  namespace.DEFAULT_NAMESPACE,
		Enabled:    true,
		Repository: &Repository{},
		Tags:       map[string]string{},
	}
}

func NewDescribeServiceRequestByClientId(clientId string) *DescribeServiceRequest {
	return &DescribeServiceRequest{
		DescribeBy: DescribeBy_SERVICE_CLIENT_ID,
		ClientId:   clientId,
	}
}

func (c *Credential) Validate(clientSecret string) error {
	if c.ClientSecret != clientSecret {
		return fmt.Errorf("client_id or client_secret is not conrrect")
	}

	return nil
}

func NewDefaultService() *Service {
	return &Service{
		Spec: &CreateServiceRequest{},
	}
}

func (req *CreateServiceRequest) Validate() error {
	return validate.Struct(req)
}

func NewService(req *CreateServiceRequest) (*Service, error) {
	// 校验参数的合法性
	if err := req.Validate(); err != nil {
		return nil, err
	}

	app := &Service{
		Id:       xid.New().String(),
		CreateAt: time.Now().UnixMilli(),
		Spec:     req,
		// 服务凭证, 生成的随机凭证
		Credential: NewRandomCredential(),
		// 生成的随机安全码
		Security: NewRandomSecurity(),
	}
	app.Id = app.FullNameHash()
	return app, nil
}

func (i *Service) FullNameHash() string {
	// fnv hash算法, 类似于md5, 只生成的字符串更短
	hash := fnv.New32a()
	hash.Write([]byte(i.FullName()))
	return fmt.Sprintf("%x", hash.Sum32())
}

// 服务的全称:
func (i *Service) FullName() string {
	return fmt.Sprintf("%s.%s", i.Spec.Namespace, i.Spec.Name)
}

func NewRandomCredential() *Credential {
	return &Credential{
		ClientId:     token.MakeBearer(24),
		ClientSecret: token.MakeBearer(32),
	}
}

func NewRandomSecurity() *Security {
	return &Security{
		EncryptKey: token.MakeBearer(64),
	}
}

func NewValidateCredentialRequest(clientId, clientSercet string) *ValidateCredentialRequest {
	return &ValidateCredentialRequest{
		ClientId:     clientId,
		ClientSecret: clientSercet,
	}
}
