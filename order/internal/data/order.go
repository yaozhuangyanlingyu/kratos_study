package data

import (
	"context"

	"order/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type OrderRepo struct {
	data *Data
	log  *log.Helper
}

// NewOrderRepo .
func NewOrderRepo(data *Data, logger log.Logger) biz.OrderRepo {
	return &OrderRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *OrderRepo) ListAll(ctx context.Context, pageNum int64, pageSize int64) ([]*biz.Order, error) {
	rsp := make([]*biz.Order, 0)
	return rsp, nil
}
