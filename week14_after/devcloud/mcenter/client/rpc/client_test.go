package rpc_test

import (
	"context"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/token"
	"testing"

	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/client/rpc"
)

func TestTokenQuery(t *testing.T) {
	//should := assert.New(t)

	//conf := client.NewDefaultConfig()
	// 设置GRPC服务地址
	// conf.SetAddress("127.0.0.1:8050")
	// 携带认证信息
	// conf.SetClientCredentials("secret_id", "secret_key")
	c, err := rpc.NewClient("127.0.0.1:18050")
	if err != nil {
		t.Fatal(err)
	}
	tk,err:=c.Token().ValidateToken(context.Background(),token.NewValidateTokenRequest("nRgNyckDCezKNx8hslR5HcEX"))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tk)
}
