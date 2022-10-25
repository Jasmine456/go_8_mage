package all

import (
	// 注册所有GRPC服务模块, 暴露给框架GRPC服务器加载, 注意 导入有先后顺序
	//_ "github.com/go_8_mage/week14_after/devcloud/mcenter/apps/book/impl"
	_ "github.com/go_8_mage/week14_after/devcloud/mcenter/apps/user/impl"
	_ "github.com/go_8_mage/week14_after/devcloud/mcenter/apps/token/impl"
)
