package service

import (
	"github.com/google/wire"

	pb "order/api/order/service/v1"
	"order/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewOrderService)

type OrderService struct {
	pb.UnimplementedOrderServer

	oc *biz.OrderUsecase
}

func NewOrderService(oc *biz.OrderUsecase) *OrderService {
	return &OrderService{
		oc: oc,
	}
}
