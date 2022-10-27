package rpc

import "github.com/caarlos0/env/v6"

// NewClientFromEnv 从环境变量中生成客户端实例
func NewClientFromEnv() (*ClientSet, error) {
	cfg := NewConfig("127.0.0.1:18050", "", "")
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	// 加载全局配置单例
	return NewClient(cfg)
}
