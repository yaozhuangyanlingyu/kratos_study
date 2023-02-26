package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "helloworld/api/student/v1"
	"helloworld/pkg"
)

const (
	StudentTableName string = "student"
)

// Notice biz层起到了 "承上启下" 的作用：限定data层提供的方法、并调用data层提供的基础方法(MySQL/Redis)实现service相关逻辑

// 学生表
type Student struct {
	Id    uint64  `gorm:"primaryKey;column:id" json:"id"`
	Name  string  `gorm:"column:name" json:"name"`
	Age   int32   `gorm:"column:age" json:"age"`
	Score float64 `gorm:"column:score" json:"score"`
}

func (s Student) TableName() string {
	return StudentTableName
}

// Notice 定义Student model 需要的方法，在data层实现，供biz层调用
type StudentServerRepo interface {
	CreateStudent(ctx context.Context, stu *Student) (*Student, error)
	ListStudent(ctx context.Context, pageNum, pageSize int) ([]*Student, error)
}

// biz
type StudentBiz struct {
	repo StudentServerRepo
	log  *log.Helper
}

// Notice 需要加到biz层的 ProviderSet 中
func NewStudentBiz(repo StudentServerRepo, logger log.Logger) *StudentBiz {
	return &StudentBiz{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

// Notice 业务主逻辑，引用data层的方法，给service层调用
func (s *StudentBiz) CreateStudent(ctx context.Context, req *pb.CreateStudentRequest) (reply *pb.CreateStudentReply, err error) {
	// 转类型
	stu := &Student{}
	errCopyTo := pkg.CopyTo(req, stu)
	if errCopyTo != nil {
		return nil, errCopyTo
	}
	ret, err := s.repo.CreateStudent(ctx, stu)
	// 转类型
	reply = new(pb.CreateStudentReply)
	errCopyTo = pkg.CopyTo(ret, &reply.Student)
	if errCopyTo != nil {
		return nil, errCopyTo
	}
	return
}

func (s *StudentBiz) ListStudent(ctx context.Context, req *pb.ListStudentRequest) (reply *pb.ListStudentReply, err error) {
	var ret []*Student
	ret, err = s.repo.ListStudent(ctx, int(req.PageNum), int(req.PageSize))
	if err != nil {
		return
	}
	reply = new(pb.ListStudentReply)
	// 转类型
	errCopyTo := pkg.CopyTo(ret, &reply.StudentList)
	if errCopyTo != nil {
		return nil, errCopyTo
	}
	return
}
