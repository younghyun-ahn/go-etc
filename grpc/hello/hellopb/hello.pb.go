// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: hello/hellopb/hello.proto

package hellopb

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

type Hello struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
}

func (x *Hello) Reset() {
	*x = Hello{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_hellopb_hello_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hello) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hello) ProtoMessage() {}

func (x *Hello) ProtoReflect() protoreflect.Message {
	mi := &file_hello_hellopb_hello_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hello.ProtoReflect.Descriptor instead.
func (*Hello) Descriptor() ([]byte, []int) {
	return file_hello_hellopb_hello_proto_rawDescGZIP(), []int{0}
}

func (x *Hello) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Hello) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hello *Hello `protobuf:"bytes,1,opt,name=hello,proto3" json:"hello,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_hellopb_hello_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hello_hellopb_hello_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_hello_hellopb_hello_proto_rawDescGZIP(), []int{1}
}

func (x *HelloRequest) GetHello() *Hello {
	if x != nil {
		return x.Hello
	}
	return nil
}

type HelloResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *HelloResponse) Reset() {
	*x = HelloResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_hellopb_hello_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloResponse) ProtoMessage() {}

func (x *HelloResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hello_hellopb_hello_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloResponse.ProtoReflect.Descriptor instead.
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return file_hello_hellopb_hello_proto_rawDescGZIP(), []int{2}
}

func (x *HelloResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type HelloWithDeadlineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hello *Hello `protobuf:"bytes,1,opt,name=hello,proto3" json:"hello,omitempty"`
}

func (x *HelloWithDeadlineRequest) Reset() {
	*x = HelloWithDeadlineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_hellopb_hello_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloWithDeadlineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloWithDeadlineRequest) ProtoMessage() {}

func (x *HelloWithDeadlineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hello_hellopb_hello_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloWithDeadlineRequest.ProtoReflect.Descriptor instead.
func (*HelloWithDeadlineRequest) Descriptor() ([]byte, []int) {
	return file_hello_hellopb_hello_proto_rawDescGZIP(), []int{3}
}

func (x *HelloWithDeadlineRequest) GetHello() *Hello {
	if x != nil {
		return x.Hello
	}
	return nil
}

type HelloWithDeadlineResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *HelloWithDeadlineResponse) Reset() {
	*x = HelloWithDeadlineResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_hellopb_hello_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloWithDeadlineResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloWithDeadlineResponse) ProtoMessage() {}

func (x *HelloWithDeadlineResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hello_hellopb_hello_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloWithDeadlineResponse.ProtoReflect.Descriptor instead.
func (*HelloWithDeadlineResponse) Descriptor() ([]byte, []int) {
	return file_hello_hellopb_hello_proto_rawDescGZIP(), []int{4}
}

func (x *HelloWithDeadlineResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_hello_hellopb_hello_proto protoreflect.FileDescriptor

var file_hello_hellopb_hello_proto_rawDesc = []byte{
	0x0a, 0x19, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x70, 0x62, 0x2f,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x22, 0x43, 0x0a, 0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61,
	0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c,
	0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x32, 0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x22, 0x27, 0x0a, 0x0d, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x22, 0x3e, 0x0a, 0x18, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x69, 0x74,
	0x68, 0x44, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x22, 0x0a, 0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x05, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x22, 0x33, 0x0a, 0x19, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x69, 0x74,
	0x68, 0x44, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x9e, 0x01, 0x0a, 0x0c, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x05, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x12, 0x13, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x58, 0x0a, 0x11, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x69, 0x74, 0x68, 0x44, 0x65, 0x61,
	0x64, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x1f, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x57, 0x69, 0x74, 0x68, 0x44, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x69, 0x74, 0x68, 0x44, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_hello_hellopb_hello_proto_rawDescOnce sync.Once
	file_hello_hellopb_hello_proto_rawDescData = file_hello_hellopb_hello_proto_rawDesc
)

func file_hello_hellopb_hello_proto_rawDescGZIP() []byte {
	file_hello_hellopb_hello_proto_rawDescOnce.Do(func() {
		file_hello_hellopb_hello_proto_rawDescData = protoimpl.X.CompressGZIP(file_hello_hellopb_hello_proto_rawDescData)
	})
	return file_hello_hellopb_hello_proto_rawDescData
}

var file_hello_hellopb_hello_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_hello_hellopb_hello_proto_goTypes = []interface{}{
	(*Hello)(nil),                     // 0: hello.Hello
	(*HelloRequest)(nil),              // 1: hello.HelloRequest
	(*HelloResponse)(nil),             // 2: hello.HelloResponse
	(*HelloWithDeadlineRequest)(nil),  // 3: hello.HelloWithDeadlineRequest
	(*HelloWithDeadlineResponse)(nil), // 4: hello.HelloWithDeadlineResponse
}
var file_hello_hellopb_hello_proto_depIdxs = []int32{
	0, // 0: hello.HelloRequest.hello:type_name -> hello.Hello
	0, // 1: hello.HelloWithDeadlineRequest.hello:type_name -> hello.Hello
	1, // 2: hello.HelloService.Hello:input_type -> hello.HelloRequest
	3, // 3: hello.HelloService.HelloWithDeadline:input_type -> hello.HelloWithDeadlineRequest
	2, // 4: hello.HelloService.Hello:output_type -> hello.HelloResponse
	4, // 5: hello.HelloService.HelloWithDeadline:output_type -> hello.HelloWithDeadlineResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_hello_hellopb_hello_proto_init() }
func file_hello_hellopb_hello_proto_init() {
	if File_hello_hellopb_hello_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hello_hellopb_hello_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hello); i {
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
		file_hello_hellopb_hello_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_hello_hellopb_hello_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloResponse); i {
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
		file_hello_hellopb_hello_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloWithDeadlineRequest); i {
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
		file_hello_hellopb_hello_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloWithDeadlineResponse); i {
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
			RawDescriptor: file_hello_hellopb_hello_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hello_hellopb_hello_proto_goTypes,
		DependencyIndexes: file_hello_hellopb_hello_proto_depIdxs,
		MessageInfos:      file_hello_hellopb_hello_proto_msgTypes,
	}.Build()
	File_hello_hellopb_hello_proto = out.File
	file_hello_hellopb_hello_proto_rawDesc = nil
	file_hello_hellopb_hello_proto_goTypes = nil
	file_hello_hellopb_hello_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HelloServiceClient is the client API for HelloService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloServiceClient interface {
	// Unary
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	HelloWithDeadline(ctx context.Context, in *HelloWithDeadlineRequest, opts ...grpc.CallOption) (*HelloWithDeadlineResponse, error)
}

type helloServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloServiceClient(cc grpc.ClientConnInterface) HelloServiceClient {
	return &helloServiceClient{cc}
}

func (c *helloServiceClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/hello.HelloService/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloServiceClient) HelloWithDeadline(ctx context.Context, in *HelloWithDeadlineRequest, opts ...grpc.CallOption) (*HelloWithDeadlineResponse, error) {
	out := new(HelloWithDeadlineResponse)
	err := c.cc.Invoke(ctx, "/hello.HelloService/HelloWithDeadline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloServiceServer is the server API for HelloService service.
type HelloServiceServer interface {
	// Unary
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
	HelloWithDeadline(context.Context, *HelloWithDeadlineRequest) (*HelloWithDeadlineResponse, error)
}

// UnimplementedHelloServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHelloServiceServer struct {
}

func (*UnimplementedHelloServiceServer) Hello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (*UnimplementedHelloServiceServer) HelloWithDeadline(context.Context, *HelloWithDeadlineRequest) (*HelloWithDeadlineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HelloWithDeadline not implemented")
}

func RegisterHelloServiceServer(s *grpc.Server, srv HelloServiceServer) {
	s.RegisterService(&_HelloService_serviceDesc, srv)
}

func _HelloService_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.HelloService/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloService_HelloWithDeadline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloWithDeadlineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).HelloWithDeadline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.HelloService/HelloWithDeadline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).HelloWithDeadline(ctx, req.(*HelloWithDeadlineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HelloService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hello.HelloService",
	HandlerType: (*HelloServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _HelloService_Hello_Handler,
		},
		{
			MethodName: "HelloWithDeadline",
			Handler:    _HelloService_HelloWithDeadline_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello/hellopb/hello.proto",
}
