package token

import (
	"context"
	"fmt"
	"time"

	"github.com/duke-git/lancet/convertor"
	"github.com/duke-git/lancet/slice"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-redis/redis/v8"
)

func LogginMiddleware(logger log.Logger, redis *redis.Client, forgetRouter ...string) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				// Notice往响应头中设置一下当前的时间戳 毫秒
				currTimestamp := time.Now().UnixNano() / 1e6
				tr.ReplyHeader().Set("currTimestamp", convertor.ToString(currTimestamp))

				// 白名单校验
				if len(forgetRouter) > 0 {
					ht, _ := tr.(*http.Transport)
					router := ht.Request().RequestURI // 获取请求的url
					fmt.Println("当前请求的URL：", router)
					if slice.Contain(forgetRouter, router) {
						return handler(ctx, req)
					}
				}

				var (
					token          string
					uid            string
					ApiTokenHeader = "auth_token"
					ApiTokenKey    = "token_exchange_uid" // redis中的key，保存token和uid的对应关系
				)
				// Notice 这里不关心token失效问题，token的刷新与重制用不同的方案操作redis就好了～
				// 获取header中携带的token参数
				token = tr.RequestHeader().Get(ApiTokenHeader)
				fmt.Println("表单请求token：", token)

				// 先判断下redis中是否有数据
				if len(token) <= 0 || !redis.HExists(ctx, ApiTokenKey, token).Val() {
					fmt.Println(">>>>>>>>>>> ", "redis中没有token数据")
					return nil, errors.New(401, "Invalid token", "登录失效，请重新登录")
				}

				// 数据redis中保存的的uid是否合法
				uid = redis.HGet(ctx, ApiTokenKey, token).Val()
				if len(uid) <= 0 {
					fmt.Println(">>>>>>>>>>> ", "redis中uid数据不合法！")
					// 既然保存的token对应的uid不合法，就把相应的key删掉吧
					errDel := redis.HDel(ctx, ApiTokenKey, token)
					if errDel != nil {
						// 出错只打log
						log.NewHelper(logger).WithContext(ctx).Errorw("删除redis中的token数据失败!", token, "uid", uid)
					}
					return nil, errors.New(401, "invalid token", "登录失效，请重新登录")
				}

				// Notice将redis中的uid放到请求的ctx中
				ctx = context.WithValue(ctx, "uid", uid)
				log.NewHelper(logger).WithContext(ctx).Infow("auth_token", token, "uid", uid)
			}
			return handler(ctx, req)
		}
	}
}
