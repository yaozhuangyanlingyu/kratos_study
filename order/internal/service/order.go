package service

import (
	"context"

	pb "order/api/order/service/v1"
)

func (s *OrderService) ListOrder(ctx context.Context, req *pb.ListOrderRequest) (*pb.ListOrderReply, error) {
	rv, err := s.oc.List(ctx, req.PageNum, req.PageSize)
	if err != nil {
		return nil, err
	}
	rspOrders := make([]*pb.ListOrderReply_Order, 0)
	for _, v := range rv {
		rspOrders = append(rspOrders, &pb.ListOrderReply_Order{
			Id:     v.Id,
			UserId: v.UserId,
		})
	}

	return &pb.ListOrderReply{
		Orders: rspOrders,
	}, nil
}
