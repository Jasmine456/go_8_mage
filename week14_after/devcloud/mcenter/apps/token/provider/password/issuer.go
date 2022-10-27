package password

import (
	"context"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/token"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/token/provider"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/user"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	AUTH_FAILED = exception.NewUnauthorized("用户名获取密码不正确")
)

type issuer struct {
	user user.Service

	log logger.Logger
}

func (i *issuer) Init() error {
	i.user = app.GetInternalApp(user.AppName).(user.Service)
	i.log = zap.L().Named("issuer.password")
	return nil
}

func (i *issuer) GrantType() token.GRANT_TYPE {
	return token.GRANT_TYPE_PASSWORD
}

func (i *issuer) IssueToken(ctx context.Context, req *token.IssueTokenRequest) (*token.Token, error) {
	switch req.GrantType {
	case token.GRANT_TYPE_PASSWORD:
		if req.Username == "" || req.Password == "" {
			return nil, AUTH_FAILED
		}

		// 检测用户的密码是否正确
		u, err := i.user.DescribeUser(ctx, user.NewDescriptUserRequestWithName(req.Username))
		if err != nil {
			return nil, err
		}
		if err := u.Password.CheckPassword(req.Password); err != nil {
			return nil, AUTH_FAILED
		}

		// 3. 颁发Token
		tk := token.NewToken(req)
		tk.Domain = u.Spec.Domain
		tk.Username = u.Spec.Username
		tk.UserType = u.Spec.Type
		tk.UserId = u.Id
		return tk, nil
	default:
		return nil, exception.NewBadRequest("grant type %s not implemented", req.GrantType)
	}
}

func init() {
	provider.Registe(&issuer{})
}
