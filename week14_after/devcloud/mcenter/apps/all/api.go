package all

import (
	// 注册所有HTTP服务模块, 暴露给框架HTTP服务器加载
	_ "github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/token/api"
	_ "github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/user/api"
)
