package rpc_test

import (
	"context"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/endpoint"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/token"
	//"os"
	"testing"

	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/client/rpc"
)

var (
	client *rpc.ClientSet
	ctx    = context.Background()
)

func TestValidateToken(t *testing.T) {
	//should := assert.New(t)

	//conf := client.NewDefaultConfig()
	// 设置GRPC服务地址
	// conf.SetAddress("127.0.0.1:8050")
	// 携带认证信息
	// conf.SetClientCredentials("secret_id", "secret_key")
	//conf:=rpc.NewConfig("127.0.0.1:18050","HlyT6JSg88phswtgbNFoa7NB","2wYf6oyT0w28jolDg0xtUdCtdE9SpT8s")

	//c, err := rpc.NewClient(client)
	//if err != nil {
	//	t.Fatal(err)
	//}
	tk, err := client.Token().ValidateToken(ctx, token.NewValidateTokenRequest("nRgNyckDCezKNx8hslR5HcEX"))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tk)
}

func TestEndpointRegistry(t *testing.T) {
	req := endpoint.NewRegistryRequest("0.1", []*endpoint.Entry{
		{
			FunctionName:     "test",
			Path:             "POST./maudit/api/v1/books",
			Method:           "POST",
			Resource:         "Book",
			AuthEnable:       true,
			PermissionEnable: true,
			Labels:           map[string]string{"action": "create"},
		},
	})
	//req.ClientId = os.Getenv("MCENTER_CLINET_ID")
	//req.ClientSecret = os.Getenv("MCENTER_CLINET_SECRET")

	req.ClientId = "HlyT6JSg88phswtgbNFoa7NB"
	req.ClientSecret = "2wYf6oyT0w28jolDg0xtUdCtdE9SpT8s"
	resp, err := client.Endpoint().RegistryEndpoint(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)

}

func init() {
	c, err := rpc.NewClientFromEnv()
	if err != nil {
		panic(err)
	}
	client = c
}
