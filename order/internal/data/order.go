package data

import (
	"context"

	"order/internal/biz"
	"order/pkg/util/pagination"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type OrderRepo struct {
	data *Data
	log  *log.Helper
}

type Order struct {
	gorm.Model
	Id     int64
	UserId int64
}

// NewOrderRepo .
func NewOrderRepo(data *Data, logger log.Logger) biz.OrderRepo {
	return &OrderRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

/**
 * 获取订单列表
 */
func (r *OrderRepo) ListAll(ctx context.Context, pageNum int64, pageSize int64) ([]*biz.Order, error) {
	var os []Order
	result := r.data.db.WithContext(ctx).
		Limit(int(pageSize)).
		Offset(int(pagination.GetPageOffset(pageNum, pageSize))).
		Find(&os)

	if result.Error != nil {
		return nil, result.Error
	}
	rsp := make([]*biz.Order, 0)
	for _, o := range os {
		rsp = append(rsp, &biz.Order{
			Id:     o.Id,
			UserId: o.UserId,
		})
	}
	return rsp, nil
}
