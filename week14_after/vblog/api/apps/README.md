# 所有业务领域模块

+ blog:文章管理
+ tag: 标签管理

# vblog 项目后端

提供RESTfulAPI接口

## 工程配置对象管理
直接os.GetEnv太简单，复杂的工程，支持多种方式的配置
+ 基于文件(json,yaml,toml),基于toml的格式来作为程序的配置
    +json：json.Marshal 标准库
    +yaml：第三方库
    +toml：第三方库  "github.com/BurntSushi/toml"
    +env: 基于环境变量，容器部署时很有用。如何解析环境变量，os.GetEnv,
  // "github.com/caarlos0/env/v6",通过定义Struct Tag 直接帮你完成 环境变量的映射
    +配置中心：nacos，etcd，console，这里不做实现
  
选择支持： toml，env

需要把工程的配置，统一为一个全局对象，当程序配置加载后，程序的任何地方都能使用（全局变量）

项目加载的配置文件，一般放在当前项目的etc目录下

ioc
控制反转
对象管理 思维
mindmaster 画图软件


