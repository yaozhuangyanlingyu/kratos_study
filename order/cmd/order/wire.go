//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/yaozhuangyanlingyu/kratos_study/order/internal/biz"
	"github.com/yaozhuangyanlingyu/kratos_study/order/internal/conf"
	"github.com/yaozhuangyanlingyu/kratos_study/order/internal/data"
	"github.com/yaozhuangyanlingyu/kratos_study/order/internal/registry"
	"github.com/yaozhuangyanlingyu/kratos_study/order/internal/server"
	"github.com/yaozhuangyanlingyu/kratos_study/order/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *conf.Registry, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, registry.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
