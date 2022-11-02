package tools

import (
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/conf"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger/zap"

	//注册所有服务
	_ "github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/all"
)

func DevelopmentSetup() {
	//初始化日志示例
	zap.DevelopmentSetup()

	//	初始化配置,提前配置好 etc/unit_env.toml
	err := conf.LoadConfigFromEnv()
	if err != nil {
		panic(err)
	}

	//	初始化全局app
	if err := app.InitAllApp(); err != nil {
		panic(err)
	}

}
