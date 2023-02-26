package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"helloworld/internal/biz"
	"time"
)

type studentRepo struct {
	// Data中封装了Mysql、Redis的曹祖
	data *Data
	log  *log.Helper
}

// Notice 需要加到data层的 ProviderSet 中
// 另外，需要studentRepo这个结构体实现 biz.StudentBiz 这个接口
func NewStudentRepo(data *Data, logger log.Logger) biz.StudentServerRepo {
	return &studentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (s *studentRepo) CreateStudent(ctx context.Context, stuReq *biz.Student) (stuReply *biz.Student, err error) {

	db := s.data.Mysql.Table(biz.StudentTableName).WithContext(ctx)
	err = db.Create(stuReq).Error
	// Notice 写入Redis测试
	s.data.RedisCli.Set(ctx, "name", stuReq.Name, time.Second*666)
	stuReply = stuReq
	return
}

func (s *studentRepo) ListStudent(ctx context.Context, pageNum, pageSize int) (stuList []*biz.Student, errList error) {

	stuList = make([]*biz.Student, 0)
	db := s.data.Mysql.Table(biz.StudentTableName).WithContext(ctx)
	if pageNum <= 1 {
		pageNum = 1
	}
	errList = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&stuList).Error
	return
}
