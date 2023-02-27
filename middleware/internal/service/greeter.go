package service

import (
	"context"

	pb "middleware/api/helloworld/v1"
	"middleware/internal/biz"
)

type GreeterService struct {
	pb.UnimplementedGreeterServer

	uc *biz.GreeterUsecase
}

func NewGreeterService() *GreeterService {
	return &GreeterService{}
}

func (s *GreeterService) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{
		Message: "hello" + req.Name,
	}, nil
}

func (s *GreeterService) Login(ctx context.Context, req *pb.Empty) (*pb.LoginReply, error) {
	return s.uc.Login(ctx, req)
}
