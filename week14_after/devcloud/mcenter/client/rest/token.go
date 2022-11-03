package rest

import (
	"context"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/token"
	"github.com/infraboard/mcube/client/rest"
	"github.com/infraboard/mcube/http/response"
)

type TokenService interface {
//	校验token
	ValidateToken(ctx context.Context,request *token.ValidateTokenRequest)(*token.Token,error)
}

type tokenImpl struct {
	client *rest.RESTClient
}


func (i *tokenImpl) ValidateToken(ctx context.Context, req *token.ValidateTokenRequest) (*token.Token, error) {
	ins := token.NewDefaultToken()
	resp := response.NewData(ins)

	err := i.client.
		Get("token").
		Header(token.VALIDATE_TOKEN_HEADER_KEY, req.AccessToken).
		Do(ctx).
		Into(resp)
	if err != nil {
		return nil, err
	}

	if resp.Error() != nil {
		return nil, err
	}

	return ins, nil
}
