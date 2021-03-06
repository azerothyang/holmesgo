// Code generated by protoc-gen-go. DO NOT EDIT.
// source: property.proto

package user

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type PropertyInfo struct {
	Age                  int32    `protobuf:"varint,1,opt,name=age,proto3" json:"age,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PropertyInfo) Reset()         { *m = PropertyInfo{} }
func (m *PropertyInfo) String() string { return proto.CompactTextString(m) }
func (*PropertyInfo) ProtoMessage()    {}
func (*PropertyInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_c191a4a5f523572b, []int{0}
}

func (m *PropertyInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PropertyInfo.Unmarshal(m, b)
}
func (m *PropertyInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PropertyInfo.Marshal(b, m, deterministic)
}
func (m *PropertyInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PropertyInfo.Merge(m, src)
}
func (m *PropertyInfo) XXX_Size() int {
	return xxx_messageInfo_PropertyInfo.Size(m)
}
func (m *PropertyInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PropertyInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PropertyInfo proto.InternalMessageInfo

func (m *PropertyInfo) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_c191a4a5f523572b, []int{1}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PropertyInfo)(nil), "user.PropertyInfo")
	proto.RegisterType((*Empty)(nil), "user.Empty")
}

func init() { proto.RegisterFile("property.proto", fileDescriptor_c191a4a5f523572b) }

var fileDescriptor_c191a4a5f523572b = []byte{
	// 142 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x28, 0xca, 0x2f,
	0x48, 0x2d, 0x2a, 0xa9, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x29, 0x2d, 0x4e, 0x2d,
	0x52, 0x52, 0xe0, 0xe2, 0x09, 0x80, 0x8a, 0x7b, 0xe6, 0xa5, 0xe5, 0x0b, 0x09, 0x70, 0x31, 0x27,
	0xa6, 0xa7, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x06, 0x81, 0x98, 0x4a, 0xec, 0x5c, 0xac, 0xae,
	0xb9, 0x05, 0x25, 0x95, 0x46, 0xe6, 0x5c, 0x1c, 0x30, 0xa5, 0x42, 0xda, 0x5c, 0x6c, 0xee, 0xa9,
	0x25, 0x8e, 0xe9, 0xa9, 0x42, 0xdc, 0x7a, 0x20, 0x73, 0xf4, 0xc0, 0x4a, 0xa4, 0x84, 0x20, 0x1c,
	0x64, 0x13, 0x95, 0x18, 0x9c, 0xa4, 0xb8, 0xb8, 0x8b, 0x2b, 0x8b, 0x4b, 0x52, 0x73, 0xc1, 0xb2,
	0x4e, 0x9c, 0xa1, 0xc5, 0xa9, 0x45, 0x01, 0x20, 0x37, 0x04, 0x30, 0x26, 0xb1, 0x81, 0x1d, 0x63,
	0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x83, 0xf1, 0x4c, 0x1b, 0x9e, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PropertyClient is the client API for Property service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PropertyClient interface {
	// get age
	GetAge(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PropertyInfo, error)
}

type propertyClient struct {
	cc *grpc.ClientConn
}

func NewPropertyClient(cc *grpc.ClientConn) PropertyClient {
	return &propertyClient{cc}
}

func (c *propertyClient) GetAge(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PropertyInfo, error) {
	out := new(PropertyInfo)
	err := c.cc.Invoke(ctx, "/user.Property/GetAge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PropertyServer is the server API for Property service.
type PropertyServer interface {
	// get age
	GetAge(context.Context, *Empty) (*PropertyInfo, error)
}

func RegisterPropertyServer(s *grpc.Server, srv PropertyServer) {
	s.RegisterService(&_Property_serviceDesc, srv)
}

func _Property_GetAge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PropertyServer).GetAge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Property/GetAge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PropertyServer).GetAge(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Property_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.Property",
	HandlerType: (*PropertyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAge",
			Handler:    _Property_GetAge_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "property.proto",
}
