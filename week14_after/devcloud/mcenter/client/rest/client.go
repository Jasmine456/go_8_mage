package rest

import (
	"github.com/infraboard/mcube/client/rest"
)

func NewClient(conf *Config) (*ClientSet, error) {
	if err := conf.Validate(); err != nil {
		return nil, err
	}

	c := rest.NewRESTClient()
	c.SetBearerTokenAuth(conf.Token)
	c.SetBaseURL(conf.Address + conf.PathPrefix)
	return &ClientSet{
		c: c,
	}, nil
}

type ClientSet struct {
	c *rest.RESTClient
}

func (c *ClientSet) Token() TokenService {
	return &tokenImpl{client: c.c}
}