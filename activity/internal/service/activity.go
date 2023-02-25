package service

import (
	"context"

	v1 "activity/api/activity/v1"
	"activity/internal/biz"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedActivityServer

	uc *biz.ActivityUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.ActivityUsecase) *GreeterService {
	return &GreeterService{uc: uc}
}

func (s *GreeterService) GetSku(ctx context.Context, in *v1.GetSkuRequest) (*v1.GetSkuReply, error) {
	if err := s.uc.GetOrderList(ctx, 1); err != nil {
		return nil, nil
	}
	return &v1.GetSkuReply{}, nil
}
