package token

import (
	"context"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

const (
	AppName = "token"
)

type Service interface {
	// 颁发Token
	IssueToken(context.Context, *IssueTokenRequest) (*Token, error)
	// 撤销Token
	RevolkToken(context.Context, *RevolkTokenRequest) (*Token, error)
	// 切换Token空间
	ChangeNamespace(context.Context, *ChangeNamespaceRequest) (*Token, error)
	// 查询Token, 用于查询Token颁发记录, 也就是登陆日志
	QueryToken(context.Context, *QueryTokenRequest) (*TokenSet, error)
	// RPC
	RPCServer
}

func NewToken(req *IssueTokenRequest) *Token {
	tk := &Token{
		AccessToken:      MakeBearer(24),
		RefreshToken:     MakeBearer(32),
		IssueAt:          time.Now().UnixMilli(),
		AccessExpiredAt:  req.ExpiredAt,
		RefreshExpiredAt: req.ExpiredAt * 4,
		GrantType:        req.GrantType,
		Type:             req.Type,
		Description:      req.Description,
		Status:           NewStatus(),
		Location:         req.Location,
	}
	switch req.GrantType {
	case GRANT_TYPE_PRIVATE_TOKEN:
		tk.Platform = PLATFORM_API
	default:
		tk.Platform = PLATFORM_WEB
	}
	return tk
}

func NewStatus() *Status {
	return &Status{
		IsBlock: false,
	}
}

func (t *Token) IsAccessTokenExpired() bool {
	// now > expired
	now := time.Now().UnixMilli()
	expired := t.IssueAt + t.AccessExpiredAt*1000
	return now > expired
}

// MakeBearer https://tools.ietf.org/html/rfc6750#section-2.1
// b64token    = 1*( ALPHA / DIGIT /"-" / "." / "_" / "~" / "+" / "/" ) *"="
func MakeBearer(lenth int) string {
	charlist := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	t := make([]string, lenth)
	rand.Seed(time.Now().UnixNano() + int64(lenth) + rand.Int63n(10000))
	for i := 0; i < lenth; i++ {
		rn := rand.Intn(len(charlist))
		w := charlist[rn : rn+1]
		t = append(t, w)
	}

	token := strings.Join(t, "")
	return token
}

// NewIssueTokenRequest 默认请求
func NewIssueTokenRequest() *IssueTokenRequest {
	return &IssueTokenRequest{
		// 单位秒
		ExpiredAt: 60*60,
	}
}

// NewRevolkTokenRequest 撤销Token请求
func NewRevolkTokenRequest() *RevolkTokenRequest {
	return &RevolkTokenRequest{}
}

// Authorization xxxxxxx
func GetTokenFromHTTPHeader(r *http.Request) string {
	auth := r.Header.Get("Authorization")
	info := strings.Split(auth, " ")
	if len(info) > 1 {
		return info[1]
	}

	return ""
}

func NewChangeNamespaceRequest() *ChangeNamespaceRequest {
	return &ChangeNamespaceRequest{}
}

func NewValidateTokenRequest(ak string) *ValidateTokenRequest {
	return &ValidateTokenRequest{
		AccessToken: ak,
	}
}

func NewDefaultToken() *Token {
	return &Token{
		Status:   NewStatus(),
		Location: NewLocation(),
	}
}

func NewLocation() *Location {
	return &Location{}
}
