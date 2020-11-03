//protoc --go_out=plugins=grpc:.  gocron.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.7.1
// source: gocron.proto

package protofile

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type TaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Command   string `protobuf:"bytes,2,opt,name=command,proto3" json:"command,omitempty"`     // 命令
	Timeout   int64  `protobuf:"varint,3,opt,name=timeout,proto3" json:"timeout,omitempty"`    // 任务执行超时时间
	Taskid    int64  `protobuf:"varint,4,opt,name=taskid,proto3" json:"taskid,omitempty"`      // 执行任务唯一ID
	Jobid     int64  `protobuf:"varint,5,opt,name=jobid,proto3" json:"jobid,omitempty"`        // 执行任务唯一ID
	Querytype string `protobuf:"bytes,6,opt,name=querytype,proto3" json:"querytype,omitempty"` //执行类型
}

func (x *TaskRequest) Reset() {
	*x = TaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gocron_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskRequest) ProtoMessage() {}

func (x *TaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gocron_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskRequest.ProtoReflect.Descriptor instead.
func (*TaskRequest) Descriptor() ([]byte, []int) {
	return file_gocron_proto_rawDescGZIP(), []int{0}
}

func (x *TaskRequest) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *TaskRequest) GetTimeout() int64 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

func (x *TaskRequest) GetTaskid() int64 {
	if x != nil {
		return x.Taskid
	}
	return 0
}

func (x *TaskRequest) GetJobid() int64 {
	if x != nil {
		return x.Jobid
	}
	return 0
}

func (x *TaskRequest) GetQuerytype() string {
	if x != nil {
		return x.Querytype
	}
	return ""
}

type TaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Output  string `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`   // 命令标准输出
	Err     string `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`         // 命令错误
	Status  int64  `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`  //执行状态
	Host    string `protobuf:"bytes,4,opt,name=host,proto3" json:"host,omitempty"`       //运行机器
	Endtime string `protobuf:"bytes,5,opt,name=endtime,proto3" json:"endtime,omitempty"` //执行结束时间
}

func (x *TaskResponse) Reset() {
	*x = TaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gocron_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskResponse) ProtoMessage() {}

func (x *TaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gocron_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskResponse.ProtoReflect.Descriptor instead.
func (*TaskResponse) Descriptor() ([]byte, []int) {
	return file_gocron_proto_rawDescGZIP(), []int{1}
}

func (x *TaskResponse) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

func (x *TaskResponse) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

func (x *TaskResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *TaskResponse) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *TaskResponse) GetEndtime() string {
	if x != nil {
		return x.Endtime
	}
	return ""
}

var File_gocron_proto protoreflect.FileDescriptor

var file_gocron_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x67, 0x6f, 0x63, 0x72, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03,
	0x72, 0x70, 0x63, 0x22, 0x8d, 0x01, 0x0a, 0x0b, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x6a, 0x6f, 0x62, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x6a, 0x6f, 0x62, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x71, 0x75, 0x65, 0x72, 0x79, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x71, 0x75, 0x65, 0x72, 0x79, 0x74,
	0x79, 0x70, 0x65, 0x22, 0x7e, 0x0a, 0x0c, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65,
	0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x72, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x74,
	0x69, 0x6d, 0x65, 0x32, 0x34, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x2c, 0x0a, 0x03, 0x52,
	0x75, 0x6e, 0x12, 0x10, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gocron_proto_rawDescOnce sync.Once
	file_gocron_proto_rawDescData = file_gocron_proto_rawDesc
)

func file_gocron_proto_rawDescGZIP() []byte {
	file_gocron_proto_rawDescOnce.Do(func() {
		file_gocron_proto_rawDescData = protoimpl.X.CompressGZIP(file_gocron_proto_rawDescData)
	})
	return file_gocron_proto_rawDescData
}

var file_gocron_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_gocron_proto_goTypes = []interface{}{
	(*TaskRequest)(nil),  // 0: rpc.TaskRequest
	(*TaskResponse)(nil), // 1: rpc.TaskResponse
}
var file_gocron_proto_depIdxs = []int32{
	0, // 0: rpc.Task.Run:input_type -> rpc.TaskRequest
	1, // 1: rpc.Task.Run:output_type -> rpc.TaskResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gocron_proto_init() }
func file_gocron_proto_init() {
	if File_gocron_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gocron_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskRequest); i {
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
		file_gocron_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskResponse); i {
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
			RawDescriptor: file_gocron_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gocron_proto_goTypes,
		DependencyIndexes: file_gocron_proto_depIdxs,
		MessageInfos:      file_gocron_proto_msgTypes,
	}.Build()
	File_gocron_proto = out.File
	file_gocron_proto_rawDesc = nil
	file_gocron_proto_goTypes = nil
	file_gocron_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TaskClient is the client API for Task service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TaskClient interface {
	Run(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TaskResponse, error)
}

type taskClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskClient(cc grpc.ClientConnInterface) TaskClient {
	return &taskClient{cc}
}

func (c *taskClient) Run(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TaskResponse, error) {
	out := new(TaskResponse)
	err := c.cc.Invoke(ctx, "/rpc.Task/Run", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskServer is the server API for Task service.
type TaskServer interface {
	Run(context.Context, *TaskRequest) (*TaskResponse, error)
}

// UnimplementedTaskServer can be embedded to have forward compatible implementations.
type UnimplementedTaskServer struct {
}

func (*UnimplementedTaskServer) Run(context.Context, *TaskRequest) (*TaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Run not implemented")
}

func RegisterTaskServer(s *grpc.Server, srv TaskServer) {
	s.RegisterService(&_Task_serviceDesc, srv)
}

func _Task_Run_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).Run(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Task/Run",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).Run(ctx, req.(*TaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Task_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.Task",
	HandlerType: (*TaskServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Run",
			Handler:    _Task_Run_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gocron.proto",
}