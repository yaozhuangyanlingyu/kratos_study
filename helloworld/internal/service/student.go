package service

import (
	"context"
	"helloworld/internal/biz"

	pb "helloworld/api/student/v1"
)

type StudentService struct {
	pb.UnimplementedStudentServiceServer

	bz *biz.StudentBiz
}

// Notice 需要加到service层的 ProviderSet 中
func NewStudentService(bz *biz.StudentBiz) *StudentService {
	return &StudentService{
		bz: bz,
	}
}

func (s *StudentService) CreateStudent(ctx context.Context, req *pb.CreateStudentRequest) (*pb.CreateStudentReply, error) {
	return s.bz.CreateStudent(ctx, req)
}
func (s *StudentService) ListStudent(ctx context.Context, req *pb.ListStudentRequest) (*pb.ListStudentReply, error) {
	return s.bz.ListStudent(ctx, req)
}
