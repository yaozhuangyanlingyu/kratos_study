package data

import (
	"context"

	v1 "middleware/api/helloworld/v1"
	"middleware/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type greeterRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger) biz.GreeterRepo {
	return &greeterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *greeterRepo) Login(ctx context.Context) (*v1.LoginReply, error) {
	return &v1.LoginReply{}, nil
}
