package impl

import (
	"context"
	"fmt"
	"github.com/infraboard/mcube/exception"
	"github.com/go_8_mage/week14_after/devcloud/mcenter/apps/token"
	"github.com/go_8_mage/week14_after/devcloud/mcenter/apps/token/provider"
)

func (s *service) IssueToken(ctx context.Context, req *token.IssueTokenRequest) (
	*token.Token, error) {
	// 颁发令牌
	tk, err := s.IssueTokenNow(ctx, req)
	if err != nil {
		return nil, err
	}
	return tk, nil
}

func (s *service) IssueTokenNow(ctx context.Context, req *token.IssueTokenRequest) (*token.Token, error) {
	// 获取令牌颁发器
	issuer := provider.Get(token.GRANT_TYPE_PASSWORD)
	if issuer == nil {
		return nil, fmt.Errorf("%s issuer not found", token.GRANT_TYPE_PASSWORD)
	}
	tk, err := issuer.IssueToken(ctx, req)
	if err != nil {
		return nil, err
	}

	// 入库保存
	if !req.DryRun {
		if err := s.save(ctx, tk); err != nil {
			return nil, err
		}
	}

	return tk, nil
}


// 撤销Token
func (s *service) RevolkToken(ctx context.Context, req *token.RevolkTokenRequest) (
	*token.Token, error) {
	// 查询Token
	tk, err := s.get(ctx, req.AccessToken)
	if err != nil {
		return nil, err
	}

	if tk.RefreshToken != req.RefreshToken {
		return nil, exception.NewBadRequest("refresh token not connrect")
	}

	if err := s.delete(ctx, tk); err != nil {
		return nil, err
	}
	return tk, nil
}

// 切换Token空间
func (s *service) ChangeNamespace(ctx context.Context, req *token.ChangeNamespaceRequest) (
	*token.Token, error) {
	return nil, nil
}

// 校验Token
func (s *service) ValidateToken(ctx context.Context, req *token.ValidateTokenRequest) (
	*token.Token, error) {
	// 查询Token
	tk, err := s.get(ctx, req.AccessToken)
	if err != nil {
		return nil, err
	}

	if tk.Status.IsBlock {
		return nil, fmt.Errorf("token is blocked")
	}

	if tk.IsAccessTokenExpired() {
		return nil, fmt.Errorf("token has expired")
	}

	return tk, nil
}

// 查询Token, 用于查询Token颁发记录, 也就是登陆日志
func (s *service) QueryToken(ctx context.Context, req *token.QueryTokenRequest) (*token.TokenSet, error) {
	return nil, nil
}

