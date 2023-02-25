package service

import (
	"context"
	"encoding/json"
	"fmt"

	pb "github.com/yaozhuangyanlingyu/kratos_study/order/api/order/service/v1"
	"github.com/yaozhuangyanlingyu/kratos_study/order/internal/conf"

	"github.com/go-kratos/kratos/contrib/config/consul/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
)

/**
 * 查询订单数据，并格式化返回
 */
func (s *OrderService) ListOrder(ctx context.Context, req *pb.ListOrderRequest) (*pb.ListOrderReply, error) {
	// 记录日志
	s.log.Info("愿2023是美好顺利的一年，与所有的不愉快说再见")

	// 查询数据库数据
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

/**
 * 创建订单
 * HTTP POST请求
 */
func (s *OrderService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (rsp *pb.CreateOrderReply, err error) {
	reqJson, _ := json.Marshal(req)
	s.log.Info("CreateOrder", zap.String("req", string(reqJson)))
	rsp = &pb.CreateOrderReply{}
	rsp.OrderId = 12345
	return rsp, nil
}

/**
 * 获取配置
 */
func (s *OrderService) GetConfig(ctx context.Context, req *pb.GetConfigRequest) (rsp *pb.GetConfigReply, err error) {
	// 从文件获取配置
	//ShowConfig()

	// 从consul中获取配置
	ShowConsulConfig()

	return &pb.GetConfigReply{}, nil
}

/**
 * 显示cunsul配置
 */
func ShowConsulConfig() {
	consulClient, err := api.NewClient(&api.Config{
		Address: "192.168.107.118:8500",
	})
	if err != nil {
		panic(err)
	}
	cs, err := consul.New(consulClient, consul.WithPath("mysql"))
	if err != nil {
		panic(err)
	}
	c := config.New(config.WithSource(cs))
	if err := c.Load(); err != nil {
		panic(err)
	}
	port, err := c.Value("user").String()
	fmt.Println(port, err)
}

/**
 * 从文件获取配置
 */
func ShowConfig() {
	flagconf := "../../configs"
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}
	cJson, err := json.Marshal(bc)
	fmt.Println(string(cJson), err)
}
