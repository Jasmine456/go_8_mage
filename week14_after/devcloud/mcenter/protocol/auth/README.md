# 服务端 认证中间件
哪些接口需要什么用户类型才能访问


http: go-restful -> 使用go-restful的认证中间件

Go Restful的路由是支持装饰（meta 就是装饰）：我们可以添加额外的标识，来控制权限，
比如 CreateUser meta{"user_type": supper}

判断权限的逻辑： 基于路由装饰的权限中间件
+ 通过token获取用户身份
+ 判断用户的身份和meta的定义是否匹配

