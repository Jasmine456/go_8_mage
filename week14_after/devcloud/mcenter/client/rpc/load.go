package rpc

import "github.com/caarlos0/env/v6"


var (
	client *ClientSet
)

// SetGlobal todo
func SetGlobal(cli *ClientSet) {
	client = cli
}

// C Global
func C() *ClientSet {
	if client == nil{
		panic("load rpc config first")
	}
	return client
}


// NewClientFromEnv 从环境变量中生成客户端实例
func LoadNewClientFromEnv() error {
	cfg := NewConfig("127.0.0.1:18050", "", "")
	if err := env.Parse(cfg); err != nil {
		return  err
	}
	// 加载全局配置单例
	c,err:=NewClient(cfg)
	if err!=nil{
		return err
	}
	client=c
	return nil
}
