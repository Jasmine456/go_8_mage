package rest

import (
	"github.com/caarlos0/env/v6"
)

var (
	client *ClientSet
)

func C() *ClientSet {
	if client == nil {
		panic("mcenter rest client config not load")
	}
	return client
}

func LoadClientFromEnv() error {
	conf := NewDefaultConfig()
	err := env.Parse(conf)
	if err != nil {
		return err
	}

	c, err := NewClient(conf)
	if err != nil {
		return err
	}
	client = c
	return nil
}
