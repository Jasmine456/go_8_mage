package rest_test

import (
	"context"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/token"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/client/rest"
	"testing"
)

var (
	c *rest.ClientSet
	ctx=context.Background()
)

func Test_ValidateToken(t *testing.T) {
	req:=token.NewValidateTokenRequest("r9amhekRs9FAYU8EdXpSRTCA")
	tk,err:=c.Token().ValidateToken(ctx,req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tk)
}


func init(){
	err:=rest.LoadClientFromEnv()
	if err != nil {
		panic(err)
	}
	c=rest.C()
}

