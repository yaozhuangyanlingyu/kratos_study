package data

import (
	"middleware/internal/conf"

	"github.com/duke-git/lancet/convertor"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/extra/redisotel/v8"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"go.opentelemetry.io/otel/attribute"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

// Data .
type Data struct {
	// TODO wrapped database client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{}, cleanup, nil
}

func NewRedis(c *conf.Data, logger log.Logger) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: c.Redis.Addr,
		//Password:     c.Redis.Password,
		DB:           int(c.Redis.Db),
		PoolSize:     int(c.Redis.PoolSize),     // 连接池数量
		MinIdleConns: int(c.Redis.MinIdleConns), //好比最小连接数
		MaxRetries:   int(c.Redis.MaxRetries),   // 命令执行失败时，最多重试多少次，默认为0即不重试
	})
	rdb.AddHook(redisotel.NewTracingHook(
		redisotel.WithAttributes(
			attribute.String("db.type", "redis"),
			attribute.String("db.ip", c.Redis.Addr),
			attribute.String("db.instance", c.Redis.Addr+"/"+convertor.ToString(c.Redis.Db)),
		),
	))
	log.NewHelper(logger).Infow("kind", "redis", "status", "enable")
	return rdb
}
