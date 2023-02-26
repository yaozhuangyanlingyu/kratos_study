//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/yaozhuangyanlingyu/kratos_study/activity/internal/biz"
	"github.com/yaozhuangyanlingyu/kratos_study/activity/internal/client"
	"github.com/yaozhuangyanlingyu/kratos_study/activity/internal/conf"
	"github.com/yaozhuangyanlingyu/kratos_study/activity/internal/data"
	"github.com/yaozhuangyanlingyu/kratos_study/activity/internal/registry"
	"github.com/yaozhuangyanlingyu/kratos_study/activity/internal/server"
	"github.com/yaozhuangyanlingyu/kratos_study/activity/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, client.ProviderSet, registry.ProviderSet, newApp))
}
