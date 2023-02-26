package biz

import (
	"context"
	"fmt"

	v1 "github.com/yaozhuangyanlingyu/kratos_study/activity/api/activity/v1"

	orderProto "ordersrv/api/order/service/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

type ActivityUsecase struct {
	log *log.Helper

	OrderRPC orderProto.OrderClient
}

func NewActivityUsecase(logger log.Logger, OrderRpc orderProto.OrderClient) *ActivityUsecase {
	return &ActivityUsecase{
		log:      log.NewHelper(logger),
		OrderRPC: OrderRpc,
	}
}

/**
 * 请求订单服务，获取订单列表
 */
func (uc *ActivityUsecase) GetOrderList(ctx context.Context, orderID int64) error {
	// 构建RPC请求
	req := &orderProto.ListOrderRequest{
		PageNum:  1,
		PageSize: 10,
		Uid:      46399747,
	}

	// 调用RPC接口，获取学生数据
	rsp, err := uc.OrderRPC.ListOrder(ctx, req)
	if err != nil {
		panic(err)
	}
	for _, o := range rsp.Orders {
		fmt.Println("OrderID：", o.Id, " UserID：", o.UserId)
	}
	return nil
}
