package registry

import (
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/hashicorp/consul/api"
	"helloworld/internal/conf"
)

func NewConsulRegistry(c *conf.Registry) *consul.Registry {
	client, err := api.NewClient(&api.Config{
		Address: c.Consul.Addr,
		Scheme:  c.Consul.Schema})
	if err != nil {
		panic(err)
	}
	return consul.New(client)
}
