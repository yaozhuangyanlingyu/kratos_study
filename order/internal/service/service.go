package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	pb "order/api/order/service/v1"
	"order/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewOrderService)

type OrderService struct {
	pb.UnimplementedOrderServer

	oc  *biz.OrderUsecase
	log *log.Helper
}

func NewOrderService(oc *biz.OrderUsecase, logger log.Logger) *OrderService {
	return &OrderService{
		oc:  oc,
		log: log.NewHelper(logger),
	}
}
