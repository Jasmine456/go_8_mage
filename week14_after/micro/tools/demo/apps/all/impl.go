package all

import (
	// 注册所有GRPC服务模块, 暴露给框架GRPC服务器加载, 注意 导入有先后顺序
	_ "go_8_mage/week14_after/micro/tools/demo/apps/book/impl"
)
