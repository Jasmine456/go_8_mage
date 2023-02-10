package all

import (
// 注册所有内部服务模块, 无须对外暴露的服务, 用于内部依赖
	_ "github.com/Jasmine456/go_8_mage/week14_after/devcloud/maudit/apps/operate/impl"
	_ "github.com/Jasmine456/go_8_mage/week14_after/devcloud/maudit/apps/book/impl"
)
