package biz

import (
	"context"

	v1 "activity/api/activity/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

type ActivityUsecase struct {
	log *log.Helper
}

func NewActivityUsecase(logger log.Logger) *ActivityUsecase {
	return &ActivityUsecase{log: log.NewHelper(logger)}
}

/**
 * 请求订单服务，获取订单列表
 */
func (uc *ActivityUsecase) GetOrderList(ctx context.Context, orderID int64) error {
	// 构建RPC请求

	// 调用RPC接口，获取学生数据

	return nil
}
