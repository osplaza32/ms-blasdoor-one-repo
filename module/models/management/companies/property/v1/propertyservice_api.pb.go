// Code generated by protoc-gen-go. DO NOT EDIT.
// source: management/companies/property/v1/propertyservice_api.proto

package propertyv1

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
	return fileDescriptor_222c3d87adfac69d, []int{0}
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
	return fileDescriptor_222c3d87adfac69d, []int{1}
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
	proto.RegisterType((*EchoRequest)(nil), "managment.companies.property.v1.EchoRequest")
	proto.RegisterType((*EchoResponse)(nil), "managment.companies.property.v1.EchoResponse")
}

func init() {
	proto.RegisterFile("management/companies/property/v1/propertyservice_api.proto", fileDescriptor_222c3d87adfac69d)
}

var fileDescriptor_222c3d87adfac69d = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0xca, 0x4d, 0xcc, 0x4b,
	0x4c, 0x4f, 0xcd, 0x4d, 0xcd, 0x2b, 0xd1, 0x4f, 0xce, 0xcf, 0x2d, 0x48, 0xcc, 0xcb, 0x4c, 0x2d,
	0xd6, 0x2f, 0x28, 0xca, 0x2f, 0x48, 0x2d, 0x2a, 0xa9, 0xd4, 0x2f, 0x33, 0x84, 0xb3, 0x8b, 0x53,
	0x8b, 0xca, 0x32, 0x93, 0x53, 0xe3, 0x13, 0x0b, 0x32, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85,
	0xe4, 0xc1, 0x7a, 0x41, 0x5a, 0xf5, 0xe0, 0x5a, 0xf5, 0x60, 0xca, 0xf5, 0xca, 0x0c, 0xa5, 0x64,
	0xd2, 0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0x13, 0x0b, 0x32, 0xf5, 0x13, 0xf3, 0xf2, 0xf2, 0x4b,
	0x12, 0x4b, 0x32, 0xf3, 0xf3, 0x8a, 0x21, 0xda, 0x95, 0x94, 0xb9, 0xb8, 0x5d, 0x93, 0x33, 0xf2,
	0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x44, 0xb8, 0x58, 0xcb, 0x12, 0x73, 0x4a, 0x53,
	0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x20, 0x1c, 0x25, 0x15, 0x2e, 0x1e, 0x88, 0xa2, 0xe2,
	0x82, 0xfc, 0xbc, 0xe2, 0x54, 0xec, 0xaa, 0x8c, 0xa6, 0x33, 0x72, 0x09, 0x05, 0xa0, 0xba, 0xd3,
	0x31, 0xc0, 0x53, 0xa8, 0x81, 0x91, 0x8b, 0x05, 0xa4, 0x5b, 0x48, 0x47, 0x8f, 0x80, 0x53, 0xf5,
	0x90, 0x5c, 0x22, 0xa5, 0x4b, 0xa4, 0x6a, 0x88, 0x93, 0x94, 0xa4, 0x9b, 0x2e, 0x3f, 0x99, 0xcc,
	0x24, 0xaa, 0x24, 0x00, 0x0a, 0xad, 0xd4, 0x8a, 0xc4, 0xdc, 0x82, 0x9c, 0x54, 0xfd, 0xd4, 0xe4,
	0x8c, 0x7c, 0x2b, 0x46, 0x2d, 0xa7, 0x69, 0x8c, 0x5c, 0xca, 0xc9, 0xf9, 0xb9, 0x84, 0x4c, 0x74,
	0x12, 0x47, 0x77, 0x7e, 0x41, 0x66, 0x00, 0x28, 0x94, 0x02, 0x18, 0xa3, 0xb8, 0x60, 0xea, 0xca,
	0x0c, 0x17, 0x31, 0x31, 0xfb, 0x3a, 0x07, 0xac, 0x62, 0x92, 0xf7, 0x85, 0x1b, 0xe7, 0x0c, 0x37,
	0x0e, 0x66, 0x82, 0x5e, 0x98, 0xe1, 0x29, 0x24, 0x15, 0x31, 0x70, 0x15, 0x31, 0x30, 0x15, 0x31,
	0x61, 0x86, 0x49, 0x6c, 0xe0, 0x48, 0x30, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xeb, 0xc1, 0x86,
	0xa0, 0x01, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PropertyserviceAPIClient is the client API for PropertyserviceAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PropertyserviceAPIClient interface {
	//Doc.
	Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error)
}

type propertyserviceAPIClient struct {
	cc *grpc.ClientConn
}

func NewPropertyserviceAPIClient(cc *grpc.ClientConn) PropertyserviceAPIClient {
	return &propertyserviceAPIClient{cc}
}

func (c *propertyserviceAPIClient) Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error) {
	out := new(EchoResponse)
	err := c.cc.Invoke(ctx, "/managment.companies.property.v1.PropertyserviceAPI/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PropertyserviceAPIServer is the server API for PropertyserviceAPI service.
type PropertyserviceAPIServer interface {
	//Doc.
	Echo(context.Context, *EchoRequest) (*EchoResponse, error)
}

// UnimplementedPropertyserviceAPIServer can be embedded to have forward compatible implementations.
type UnimplementedPropertyserviceAPIServer struct {
}

func (*UnimplementedPropertyserviceAPIServer) Echo(ctx context.Context, req *EchoRequest) (*EchoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}

func RegisterPropertyserviceAPIServer(s *grpc.Server, srv PropertyserviceAPIServer) {
	s.RegisterService(&_PropertyserviceAPI_serviceDesc, srv)
}

func _PropertyserviceAPI_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PropertyserviceAPIServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/managment.companies.property.v1.PropertyserviceAPI/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PropertyserviceAPIServer).Echo(ctx, req.(*EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PropertyserviceAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "managment.companies.property.v1.PropertyserviceAPI",
	HandlerType: (*PropertyserviceAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _PropertyserviceAPI_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "management/companies/property/v1/propertyservice_api.proto",
}
