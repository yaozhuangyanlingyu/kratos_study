// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.0
// source: api/order/service/v1/order.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 获取配置信息
type GetConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetConfigRequest) Reset() {
	*x = GetConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_order_service_v1_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfigRequest) ProtoMessage() {}

func (x *GetConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_order_service_v1_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConfigRequest.ProtoReflect.Descriptor instead.
func (*GetConfigRequest) Descriptor() ([]byte, []int) {
	return file_api_order_service_v1_order_proto_rawDescGZIP(), []int{0}
}

type GetConfigReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetConfigReply) Reset() {
	*x = GetConfigReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_order_service_v1_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConfigReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfigReply) ProtoMessage() {}

func (x *GetConfigReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_order_service_v1_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConfigReply.ProtoReflect.Descriptor instead.
func (*GetConfigReply) Descriptor() ([]byte, []int) {
	return file_api_order_service_v1_order_proto_rawDescGZIP(), []int{1}
}

// 创建订单
type CreateOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId int64 `protobuf:"varint,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"` // 商品ID
	BuyNum    int64 `protobuf:"varint,2,opt,name=buy_num,json=buyNum,proto3" json:"buy_num,omitempty"`          // 购买数量
}

func (x *CreateOrderRequest) Reset() {
	*x = CreateOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_order_service_v1_order_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderRequest) ProtoMessage() {}

func (x *CreateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_order_service_v1_order_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderRequest.ProtoReflect.Descriptor instead.
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return file_api_order_service_v1_order_proto_rawDescGZIP(), []int{2}
}

func (x *CreateOrderRequest) GetProductId() int64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *CreateOrderRequest) GetBuyNum() int64 {
	if x != nil {
		return x.BuyNum
	}
	return 0
}

type CreateOrderReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId int64 `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"` // 订单ID
}

func (x *CreateOrderReply) Reset() {
	*x = CreateOrderReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_order_service_v1_order_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderReply) ProtoMessage() {}

func (x *CreateOrderReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_order_service_v1_order_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderReply.ProtoReflect.Descriptor instead.
func (*CreateOrderReply) Descriptor() ([]byte, []int) {
	return file_api_order_service_v1_order_proto_rawDescGZIP(), []int{3}
}

func (x *CreateOrderReply) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

// 获取订单列表
type ListOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid      int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	PageNum  int64 `protobuf:"varint,2,opt,name=page_num,json=pageNum,proto3" json:"page_num,omitempty"`
	PageSize int64 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *ListOrderRequest) Reset() {
	*x = ListOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_order_service_v1_order_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrderRequest) ProtoMessage() {}

func (x *ListOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_order_service_v1_order_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrderRequest.ProtoReflect.Descriptor instead.
func (*ListOrderRequest) Descriptor() ([]byte, []int) {
	return file_api_order_service_v1_order_proto_rawDescGZIP(), []int{4}
}

func (x *ListOrderRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *ListOrderRequest) GetPageNum() int64 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

func (x *ListOrderRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ListOrderReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orders []*ListOrderReply_Order `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
}

func (x *ListOrderReply) Reset() {
	*x = ListOrderReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_order_service_v1_order_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOrderReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrderReply) ProtoMessage() {}

func (x *ListOrderReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_order_service_v1_order_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrderReply.ProtoReflect.Descriptor instead.
func (*ListOrderReply) Descriptor() ([]byte, []int) {
	return file_api_order_service_v1_order_proto_rawDescGZIP(), []int{5}
}

func (x *ListOrderReply) GetOrders() []*ListOrderReply_Order {
	if x != nil {
		return x.Orders
	}
	return nil
}

type ListOrderReply_Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId int64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *ListOrderReply_Order) Reset() {
	*x = ListOrderReply_Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_order_service_v1_order_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOrderReply_Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrderReply_Order) ProtoMessage() {}

func (x *ListOrderReply_Order) ProtoReflect() protoreflect.Message {
	mi := &file_api_order_service_v1_order_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrderReply_Order.ProtoReflect.Descriptor instead.
func (*ListOrderReply_Order) Descriptor() ([]byte, []int) {
	return file_api_order_service_v1_order_proto_rawDescGZIP(), []int{5, 0}
}

func (x *ListOrderReply_Order) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ListOrderReply_Order) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

var File_api_order_service_v1_order_proto protoreflect.FileDescriptor

var file_api_order_service_v1_order_proto_rawDesc = []byte{
	0x0a, 0x20, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x10, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x12, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x4c, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x62, 0x75, 0x79, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x62, 0x75, 0x79, 0x4e, 0x75, 0x6d, 0x22, 0x2d, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x5c, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x70,
	0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x22, 0x82, 0x01, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3e, 0x0a, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x06,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x1a, 0x30, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x32, 0xcc, 0x02, 0x0a, 0x05, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x12, 0x66, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12,
	0x22, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x68, 0x0a, 0x09, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x22, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x15, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x71, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x12, 0x24, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x18, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22, 0x0d, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x42, 0x19, 0x5a, 0x17, 0x61, 0x70, 0x69, 0x2f, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_order_service_v1_order_proto_rawDescOnce sync.Once
	file_api_order_service_v1_order_proto_rawDescData = file_api_order_service_v1_order_proto_rawDesc
)

func file_api_order_service_v1_order_proto_rawDescGZIP() []byte {
	file_api_order_service_v1_order_proto_rawDescOnce.Do(func() {
		file_api_order_service_v1_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_order_service_v1_order_proto_rawDescData)
	})
	return file_api_order_service_v1_order_proto_rawDescData
}

var file_api_order_service_v1_order_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_order_service_v1_order_proto_goTypes = []interface{}{
	(*GetConfigRequest)(nil),     // 0: order.service.v1.GetConfigRequest
	(*GetConfigReply)(nil),       // 1: order.service.v1.GetConfigReply
	(*CreateOrderRequest)(nil),   // 2: order.service.v1.CreateOrderRequest
	(*CreateOrderReply)(nil),     // 3: order.service.v1.CreateOrderReply
	(*ListOrderRequest)(nil),     // 4: order.service.v1.ListOrderRequest
	(*ListOrderReply)(nil),       // 5: order.service.v1.ListOrderReply
	(*ListOrderReply_Order)(nil), // 6: order.service.v1.ListOrderReply.Order
}
var file_api_order_service_v1_order_proto_depIdxs = []int32{
	6, // 0: order.service.v1.ListOrderReply.orders:type_name -> order.service.v1.ListOrderReply.Order
	4, // 1: order.service.v1.Order.ListOrder:input_type -> order.service.v1.ListOrderRequest
	0, // 2: order.service.v1.Order.GetConfig:input_type -> order.service.v1.GetConfigRequest
	2, // 3: order.service.v1.Order.CreateOrder:input_type -> order.service.v1.CreateOrderRequest
	5, // 4: order.service.v1.Order.ListOrder:output_type -> order.service.v1.ListOrderReply
	1, // 5: order.service.v1.Order.GetConfig:output_type -> order.service.v1.GetConfigReply
	3, // 6: order.service.v1.Order.CreateOrder:output_type -> order.service.v1.CreateOrderReply
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_order_service_v1_order_proto_init() }
func file_api_order_service_v1_order_proto_init() {
	if File_api_order_service_v1_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_order_service_v1_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConfigRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_order_service_v1_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConfigReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_order_service_v1_order_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrderRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_order_service_v1_order_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrderReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_order_service_v1_order_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOrderRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_order_service_v1_order_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOrderReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_order_service_v1_order_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOrderReply_Order); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_order_service_v1_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_order_service_v1_order_proto_goTypes,
		DependencyIndexes: file_api_order_service_v1_order_proto_depIdxs,
		MessageInfos:      file_api_order_service_v1_order_proto_msgTypes,
	}.Build()
	File_api_order_service_v1_order_proto = out.File
	file_api_order_service_v1_order_proto_rawDesc = nil
	file_api_order_service_v1_order_proto_goTypes = nil
	file_api_order_service_v1_order_proto_depIdxs = nil
}
