package conf_test

import (
	"testing"
	"go_8_mage/week14/vblog/api/apps/conf"
)

func TestLoadConfigFromEnv(t *testing.T) {
	err:=conf.LoadConfigFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(conf.C())
}
