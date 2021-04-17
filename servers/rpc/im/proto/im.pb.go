// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: im.proto

package proto

import (
	context "context"
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

type IsOnlineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *IsOnlineRequest) Reset() {
	*x = IsOnlineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_im_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsOnlineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsOnlineRequest) ProtoMessage() {}

func (x *IsOnlineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_im_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsOnlineRequest.ProtoReflect.Descriptor instead.
func (*IsOnlineRequest) Descriptor() ([]byte, []int) {
	return file_im_proto_rawDescGZIP(), []int{0}
}

func (x *IsOnlineRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type IsOnlineResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsOnline bool `protobuf:"varint,1,opt,name=isOnline,proto3" json:"isOnline,omitempty"`
}

func (x *IsOnlineResponse) Reset() {
	*x = IsOnlineResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_im_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsOnlineResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsOnlineResponse) ProtoMessage() {}

func (x *IsOnlineResponse) ProtoReflect() protoreflect.Message {
	mi := &file_im_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsOnlineResponse.ProtoReflect.Descriptor instead.
func (*IsOnlineResponse) Descriptor() ([]byte, []int) {
	return file_im_proto_rawDescGZIP(), []int{1}
}

func (x *IsOnlineResponse) GetIsOnline() bool {
	if x != nil {
		return x.IsOnline
	}
	return false
}

type SendToUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetUid  int64  `protobuf:"varint,1,opt,name=targetUid,proto3" json:"targetUid,omitempty"`  //接收消息方
	MsgType    int64  `protobuf:"varint,2,opt,name=msgType,proto3" json:"msgType,omitempty"`      //消息类型
	MsgContent string `protobuf:"bytes,3,opt,name=msgContent,proto3" json:"msgContent,omitempty"` //消息内容
}

func (x *SendToUserRequest) Reset() {
	*x = SendToUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_im_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendToUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendToUserRequest) ProtoMessage() {}

func (x *SendToUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_im_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendToUserRequest.ProtoReflect.Descriptor instead.
func (*SendToUserRequest) Descriptor() ([]byte, []int) {
	return file_im_proto_rawDescGZIP(), []int{2}
}

func (x *SendToUserRequest) GetTargetUid() int64 {
	if x != nil {
		return x.TargetUid
	}
	return 0
}

func (x *SendToUserRequest) GetMsgType() int64 {
	if x != nil {
		return x.MsgType
	}
	return 0
}

func (x *SendToUserRequest) GetMsgContent() string {
	if x != nil {
		return x.MsgContent
	}
	return ""
}

type SendToUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"` // 发送结果
}

func (x *SendToUserResponse) Reset() {
	*x = SendToUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_im_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendToUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendToUserResponse) ProtoMessage() {}

func (x *SendToUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_im_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendToUserResponse.ProtoReflect.Descriptor instead.
func (*SendToUserResponse) Descriptor() ([]byte, []int) {
	return file_im_proto_rawDescGZIP(), []int{3}
}

func (x *SendToUserResponse) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

var File_im_proto protoreflect.FileDescriptor

var file_im_proto_rawDesc = []byte{
	0x0a, 0x08, 0x69, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x69, 0x6d, 0x22, 0x23,
	0x0a, 0x0f, 0x49, 0x73, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03,
	0x75, 0x69, 0x64, 0x22, 0x2e, 0x0a, 0x10, 0x49, 0x73, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x4f, 0x6e, 0x6c,
	0x69, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x4f, 0x6e, 0x6c,
	0x69, 0x6e, 0x65, 0x22, 0x6b, 0x0a, 0x11, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x6f, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x55, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x55, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x73, 0x67, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x73, 0x67, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x22, 0x2c, 0x0a, 0x12, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x7c,
	0x0a, 0x02, 0x49, 0x6d, 0x12, 0x37, 0x0a, 0x08, 0x49, 0x73, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65,
	0x12, 0x13, 0x2e, 0x69, 0x6d, 0x2e, 0x49, 0x73, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x69, 0x6d, 0x2e, 0x49, 0x73, 0x4f, 0x6e, 0x6c,
	0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a,
	0x0a, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x69, 0x6d,
	0x2e, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x69, 0x6d, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x6f, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x05, 0x5a, 0x03,
	0x2f, 0x69, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_im_proto_rawDescOnce sync.Once
	file_im_proto_rawDescData = file_im_proto_rawDesc
)

func file_im_proto_rawDescGZIP() []byte {
	file_im_proto_rawDescOnce.Do(func() {
		file_im_proto_rawDescData = protoimpl.X.CompressGZIP(file_im_proto_rawDescData)
	})
	return file_im_proto_rawDescData
}

var file_im_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_im_proto_goTypes = []interface{}{
	(*IsOnlineRequest)(nil),    // 0: im.IsOnlineRequest
	(*IsOnlineResponse)(nil),   // 1: im.IsOnlineResponse
	(*SendToUserRequest)(nil),  // 2: im.SendToUserRequest
	(*SendToUserResponse)(nil), // 3: im.SendToUserResponse
}
var file_im_proto_depIdxs = []int32{
	0, // 0: im.Im.IsOnline:input_type -> im.IsOnlineRequest
	2, // 1: im.Im.SendToUser:input_type -> im.SendToUserRequest
	1, // 2: im.Im.IsOnline:output_type -> im.IsOnlineResponse
	3, // 3: im.Im.SendToUser:output_type -> im.SendToUserResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_im_proto_init() }
func file_im_proto_init() {
	if File_im_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_im_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsOnlineRequest); i {
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
		file_im_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsOnlineResponse); i {
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
		file_im_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendToUserRequest); i {
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
		file_im_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendToUserResponse); i {
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
			RawDescriptor: file_im_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_im_proto_goTypes,
		DependencyIndexes: file_im_proto_depIdxs,
		MessageInfos:      file_im_proto_msgTypes,
	}.Build()
	File_im_proto = out.File
	file_im_proto_rawDesc = nil
	file_im_proto_goTypes = nil
	file_im_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ImClient is the client API for Im service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ImClient interface {
	// 查询用户是否在线方法
	IsOnline(ctx context.Context, in *IsOnlineRequest, opts ...grpc.CallOption) (*IsOnlineResponse, error)
	// 发送实时消息给用户
	SendToUser(ctx context.Context, in *SendToUserRequest, opts ...grpc.CallOption) (*SendToUserResponse, error)
}

type imClient struct {
	cc grpc.ClientConnInterface
}

func NewImClient(cc grpc.ClientConnInterface) ImClient {
	return &imClient{cc}
}

func (c *imClient) IsOnline(ctx context.Context, in *IsOnlineRequest, opts ...grpc.CallOption) (*IsOnlineResponse, error) {
	out := new(IsOnlineResponse)
	err := c.cc.Invoke(ctx, "/im.Im/IsOnline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imClient) SendToUser(ctx context.Context, in *SendToUserRequest, opts ...grpc.CallOption) (*SendToUserResponse, error) {
	out := new(SendToUserResponse)
	err := c.cc.Invoke(ctx, "/im.Im/SendToUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImServer is the server API for Im service.
type ImServer interface {
	// 查询用户是否在线方法
	IsOnline(context.Context, *IsOnlineRequest) (*IsOnlineResponse, error)
	// 发送实时消息给用户
	SendToUser(context.Context, *SendToUserRequest) (*SendToUserResponse, error)
}

// UnimplementedImServer can be embedded to have forward compatible implementations.
type UnimplementedImServer struct {
}

func (*UnimplementedImServer) IsOnline(context.Context, *IsOnlineRequest) (*IsOnlineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsOnline not implemented")
}
func (*UnimplementedImServer) SendToUser(context.Context, *SendToUserRequest) (*SendToUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendToUser not implemented")
}

func RegisterImServer(s *grpc.Server, srv ImServer) {
	s.RegisterService(&_Im_serviceDesc, srv)
}

func _Im_IsOnline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsOnlineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImServer).IsOnline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/im.Im/IsOnline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImServer).IsOnline(ctx, req.(*IsOnlineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Im_SendToUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendToUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImServer).SendToUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/im.Im/SendToUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImServer).SendToUser(ctx, req.(*SendToUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Im_serviceDesc = grpc.ServiceDesc{
	ServiceName: "im.Im",
	HandlerType: (*ImServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsOnline",
			Handler:    _Im_IsOnline_Handler,
		},
		{
			MethodName: "SendToUser",
			Handler:    _Im_SendToUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "im.proto",
}
