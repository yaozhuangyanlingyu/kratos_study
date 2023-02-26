// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: student/v1/student.proto

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

type StudentBase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    uint64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Age   int32   `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	Score float32 `protobuf:"fixed32,4,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *StudentBase) Reset() {
	*x = StudentBase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_student_v1_student_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StudentBase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StudentBase) ProtoMessage() {}

func (x *StudentBase) ProtoReflect() protoreflect.Message {
	mi := &file_student_v1_student_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StudentBase.ProtoReflect.Descriptor instead.
func (*StudentBase) Descriptor() ([]byte, []int) {
	return file_student_v1_student_proto_rawDescGZIP(), []int{0}
}

func (x *StudentBase) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *StudentBase) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StudentBase) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *StudentBase) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

type CreateStudentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age   int32   `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Score float32 `protobuf:"fixed32,3,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *CreateStudentRequest) Reset() {
	*x = CreateStudentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_student_v1_student_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStudentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStudentRequest) ProtoMessage() {}

func (x *CreateStudentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_student_v1_student_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStudentRequest.ProtoReflect.Descriptor instead.
func (*CreateStudentRequest) Descriptor() ([]byte, []int) {
	return file_student_v1_student_proto_rawDescGZIP(), []int{1}
}

func (x *CreateStudentRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateStudentRequest) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *CreateStudentRequest) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

type CreateStudentReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Student *StudentBase `protobuf:"bytes,1,opt,name=student,proto3" json:"student,omitempty"`
}

func (x *CreateStudentReply) Reset() {
	*x = CreateStudentReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_student_v1_student_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStudentReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStudentReply) ProtoMessage() {}

func (x *CreateStudentReply) ProtoReflect() protoreflect.Message {
	mi := &file_student_v1_student_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStudentReply.ProtoReflect.Descriptor instead.
func (*CreateStudentReply) Descriptor() ([]byte, []int) {
	return file_student_v1_student_proto_rawDescGZIP(), []int{2}
}

func (x *CreateStudentReply) GetStudent() *StudentBase {
	if x != nil {
		return x.Student
	}
	return nil
}

type ListStudentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageNum  int32 `protobuf:"varint,1,opt,name=pageNum,proto3" json:"pageNum,omitempty"`
	PageSize int32 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
}

func (x *ListStudentRequest) Reset() {
	*x = ListStudentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_student_v1_student_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListStudentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStudentRequest) ProtoMessage() {}

func (x *ListStudentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_student_v1_student_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStudentRequest.ProtoReflect.Descriptor instead.
func (*ListStudentRequest) Descriptor() ([]byte, []int) {
	return file_student_v1_student_proto_rawDescGZIP(), []int{3}
}

func (x *ListStudentRequest) GetPageNum() int32 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

func (x *ListStudentRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ListStudentReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StudentList []*StudentBase `protobuf:"bytes,1,rep,name=studentList,proto3" json:"studentList,omitempty"`
}

func (x *ListStudentReply) Reset() {
	*x = ListStudentReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_student_v1_student_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListStudentReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStudentReply) ProtoMessage() {}

func (x *ListStudentReply) ProtoReflect() protoreflect.Message {
	mi := &file_student_v1_student_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStudentReply.ProtoReflect.Descriptor instead.
func (*ListStudentReply) Descriptor() ([]byte, []int) {
	return file_student_v1_student_proto_rawDescGZIP(), []int{4}
}

func (x *ListStudentReply) GetStudentList() []*StudentBase {
	if x != nil {
		return x.StudentList
	}
	return nil
}

var File_student_v1_student_proto protoreflect.FileDescriptor

var file_student_v1_student_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x59, 0x0a, 0x0b, 0x53, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x42, 0x61, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x73, 0x63,
	0x6f, 0x72, 0x65, 0x22, 0x52, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x4b, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x35, 0x0a,
	0x07, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x42, 0x61, 0x73, 0x65, 0x52, 0x07, 0x73, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x22, 0x4a, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61,
	0x67, 0x65, 0x4e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x61, 0x67,
	0x65, 0x4e, 0x75, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x22, 0x51, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x3d, 0x0a, 0x0b, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x42, 0x61, 0x73, 0x65, 0x52, 0x0b, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x32, 0x81, 0x02, 0x0a, 0x0e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7c, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x22, 0x16, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x3a, 0x01, 0x2a, 0x12, 0x71, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x12, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x16, 0x12, 0x14, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x42, 0x30, 0x0a, 0x0e, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x1c, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_student_v1_student_proto_rawDescOnce sync.Once
	file_student_v1_student_proto_rawDescData = file_student_v1_student_proto_rawDesc
)

func file_student_v1_student_proto_rawDescGZIP() []byte {
	file_student_v1_student_proto_rawDescOnce.Do(func() {
		file_student_v1_student_proto_rawDescData = protoimpl.X.CompressGZIP(file_student_v1_student_proto_rawDescData)
	})
	return file_student_v1_student_proto_rawDescData
}

var file_student_v1_student_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_student_v1_student_proto_goTypes = []interface{}{
	(*StudentBase)(nil),          // 0: api.student.v1.StudentBase
	(*CreateStudentRequest)(nil), // 1: api.student.v1.CreateStudentRequest
	(*CreateStudentReply)(nil),   // 2: api.student.v1.CreateStudentReply
	(*ListStudentRequest)(nil),   // 3: api.student.v1.ListStudentRequest
	(*ListStudentReply)(nil),     // 4: api.student.v1.ListStudentReply
}
var file_student_v1_student_proto_depIdxs = []int32{
	0, // 0: api.student.v1.CreateStudentReply.student:type_name -> api.student.v1.StudentBase
	0, // 1: api.student.v1.ListStudentReply.studentList:type_name -> api.student.v1.StudentBase
	1, // 2: api.student.v1.StudentService.CreateStudent:input_type -> api.student.v1.CreateStudentRequest
	3, // 3: api.student.v1.StudentService.ListStudent:input_type -> api.student.v1.ListStudentRequest
	2, // 4: api.student.v1.StudentService.CreateStudent:output_type -> api.student.v1.CreateStudentReply
	4, // 5: api.student.v1.StudentService.ListStudent:output_type -> api.student.v1.ListStudentReply
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_student_v1_student_proto_init() }
func file_student_v1_student_proto_init() {
	if File_student_v1_student_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_student_v1_student_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StudentBase); i {
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
		file_student_v1_student_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStudentRequest); i {
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
		file_student_v1_student_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStudentReply); i {
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
		file_student_v1_student_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListStudentRequest); i {
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
		file_student_v1_student_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListStudentReply); i {
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
			RawDescriptor: file_student_v1_student_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_student_v1_student_proto_goTypes,
		DependencyIndexes: file_student_v1_student_proto_depIdxs,
		MessageInfos:      file_student_v1_student_proto_msgTypes,
	}.Build()
	File_student_v1_student_proto = out.File
	file_student_v1_student_proto_rawDesc = nil
	file_student_v1_student_proto_goTypes = nil
	file_student_v1_student_proto_depIdxs = nil
}