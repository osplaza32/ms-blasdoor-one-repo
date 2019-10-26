// Code generated by protoc-gen-go. DO NOT EDIT.
// source: management/permissions/resource/v1/resourceservice_api.proto

package resourcev1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type EchoRequest struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EchoRequest) Reset()         { *m = EchoRequest{} }
func (m *EchoRequest) String() string { return proto.CompactTextString(m) }
func (*EchoRequest) ProtoMessage()    {}
func (*EchoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed93a65bbbd21b79, []int{0}
}

func (m *EchoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoRequest.Unmarshal(m, b)
}
func (m *EchoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoRequest.Marshal(b, m, deterministic)
}
func (m *EchoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoRequest.Merge(m, src)
}
func (m *EchoRequest) XXX_Size() int {
	return xxx_messageInfo_EchoRequest.Size(m)
}
func (m *EchoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EchoRequest proto.InternalMessageInfo

func (m *EchoRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type EchoResponse struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EchoResponse) Reset()         { *m = EchoResponse{} }
func (m *EchoResponse) String() string { return proto.CompactTextString(m) }
func (*EchoResponse) ProtoMessage()    {}
func (*EchoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed93a65bbbd21b79, []int{1}
}

func (m *EchoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoResponse.Unmarshal(m, b)
}
func (m *EchoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoResponse.Marshal(b, m, deterministic)
}
func (m *EchoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoResponse.Merge(m, src)
}
func (m *EchoResponse) XXX_Size() int {
	return xxx_messageInfo_EchoResponse.Size(m)
}
func (m *EchoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EchoResponse proto.InternalMessageInfo

func (m *EchoResponse) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*EchoRequest)(nil), "management.permissions.resource.v1.EchoRequest")
	proto.RegisterType((*EchoResponse)(nil), "management.permissions.resource.v1.EchoResponse")
}

func init() {
	proto.RegisterFile("management/permissions/resource/v1/resourceservice_api.proto", fileDescriptor_ed93a65bbbd21b79)
}

var fileDescriptor_ed93a65bbbd21b79 = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0xc9, 0x4d, 0xcc, 0x4b,
	0x4c, 0x4f, 0xcd, 0x4d, 0xcd, 0x2b, 0xd1, 0x2f, 0x48, 0x2d, 0xca, 0xcd, 0x2c, 0x2e, 0xce, 0xcc,
	0xcf, 0x2b, 0xd6, 0x2f, 0x4a, 0x2d, 0xce, 0x2f, 0x2d, 0x4a, 0x4e, 0xd5, 0x2f, 0x33, 0x84, 0xb3,
	0x8b, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0xe3, 0x13, 0x0b, 0x32, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b,
	0xf2, 0x85, 0x94, 0x10, 0xba, 0xf5, 0x90, 0x74, 0xeb, 0xc1, 0x74, 0xe8, 0x95, 0x19, 0x4a, 0xc9,
	0xa4, 0xe7, 0xe7, 0xa7, 0xe7, 0xa4, 0xea, 0x27, 0x16, 0x64, 0xea, 0x27, 0xe6, 0xe5, 0xe5, 0x97,
	0x24, 0x96, 0x80, 0xd5, 0x80, 0x4d, 0x50, 0x52, 0xe6, 0xe2, 0x76, 0x4d, 0xce, 0xc8, 0x0f, 0x4a,
	0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe1, 0x62, 0x2d, 0x4b, 0xcc, 0x29, 0x4d, 0x95, 0x60,
	0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x70, 0x94, 0x54, 0xb8, 0x78, 0x20, 0x8a, 0x8a, 0x0b, 0xf2,
	0xf3, 0x8a, 0x53, 0xb1, 0xab, 0x32, 0x9a, 0xcb, 0xc8, 0x25, 0x14, 0x84, 0xea, 0x54, 0xc7, 0x00,
	0x4f, 0xa1, 0x36, 0x46, 0x2e, 0x16, 0x90, 0x6e, 0x21, 0x7d, 0x3d, 0xc2, 0xae, 0xd5, 0x43, 0x72,
	0x8c, 0x94, 0x01, 0xf1, 0x1a, 0x20, 0x0e, 0x53, 0x92, 0x6e, 0xba, 0xfc, 0x64, 0x32, 0x93, 0xa8,
	0x92, 0x00, 0x28, 0xd8, 0x52, 0x2b, 0x12, 0x73, 0x0b, 0x72, 0x52, 0xf5, 0x53, 0x93, 0x33, 0xf2,
	0xad, 0x18, 0xb5, 0x9c, 0xe6, 0x33, 0x72, 0xa9, 0x25, 0xe7, 0xe7, 0x12, 0x61, 0xa8, 0x93, 0x38,
	0xba, 0x3f, 0x0a, 0x32, 0x03, 0x40, 0xc1, 0x15, 0xc0, 0x18, 0xc5, 0x05, 0x53, 0x57, 0x66, 0xb8,
	0x88, 0x89, 0xd9, 0x37, 0x20, 0x68, 0x15, 0x93, 0x92, 0x2f, 0xc2, 0xc4, 0x00, 0x24, 0x13, 0x61,
	0x86, 0xe8, 0x85, 0x19, 0x9e, 0x42, 0x56, 0x14, 0x83, 0xa4, 0x28, 0x06, 0xa6, 0x28, 0x26, 0xcc,
	0x30, 0x89, 0x0d, 0x1c, 0x27, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3f, 0x80, 0xbf, 0x10,
	0x15, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ResourceserviceAPIClient is the client API for ResourceserviceAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ResourceserviceAPIClient interface {
	//Doc.
	Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error)
}

type resourceserviceAPIClient struct {
	cc *grpc.ClientConn
}

func NewResourceserviceAPIClient(cc *grpc.ClientConn) ResourceserviceAPIClient {
	return &resourceserviceAPIClient{cc}
}

func (c *resourceserviceAPIClient) Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error) {
	out := new(EchoResponse)
	err := c.cc.Invoke(ctx, "/management.permissions.resource.v1.ResourceserviceAPI/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResourceserviceAPIServer is the server API for ResourceserviceAPI service.
type ResourceserviceAPIServer interface {
	//Doc.
	Echo(context.Context, *EchoRequest) (*EchoResponse, error)
}

// UnimplementedResourceserviceAPIServer can be embedded to have forward compatible implementations.
type UnimplementedResourceserviceAPIServer struct {
}

func (*UnimplementedResourceserviceAPIServer) Echo(ctx context.Context, req *EchoRequest) (*EchoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}

func RegisterResourceserviceAPIServer(s *grpc.Server, srv ResourceserviceAPIServer) {
	s.RegisterService(&_ResourceserviceAPI_serviceDesc, srv)
}

func _ResourceserviceAPI_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceserviceAPIServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.permissions.resource.v1.ResourceserviceAPI/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceserviceAPIServer).Echo(ctx, req.(*EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ResourceserviceAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "management.permissions.resource.v1.ResourceserviceAPI",
	HandlerType: (*ResourceserviceAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _ResourceserviceAPI_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "management/permissions/resource/v1/resourceservice_api.proto",
}
