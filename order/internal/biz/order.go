package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// Order is a Order model.
type Order struct {
	Id     int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId int64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

// OrderRepo is a Greater repo.
type OrderRepo interface {
	ListAll(ctx context.Context, pageNum int64, pageSize int64) ([]*Order, error)
}

// OrderUsecase is a Order usecase.
type OrderUsecase struct {
	repo OrderRepo
	log  *log.Helper
}

// NewOrderUsecase new a Order usecase.
func NewOrderUsecase(repo OrderRepo, logger log.Logger) *OrderUsecase {
	return &OrderUsecase{repo: repo, log: log.NewHelper(logger)}
}

/**
 * 订单列表
 */
func (uc *OrderUsecase) List(ctx context.Context, pageNum, pageSize int64) ([]*Order, error) {
	uc.log.WithContext(ctx).Infof("ListOrder pageNum: %v pageSize：%v", pageNum, pageSize)
	return uc.repo.ListAll(ctx, pageNum, pageSize)
}
