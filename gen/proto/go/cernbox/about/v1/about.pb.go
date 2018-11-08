// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cernbox/about/v1/about.proto

package aboutv1pb

import (
	context "context"
	fmt "fmt"
	rpc "github.com/cernbox/cernboxapis/gen/proto/go/cernbox/rpc"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ListMembersRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListMembersRequest) Reset()         { *m = ListMembersRequest{} }
func (m *ListMembersRequest) String() string { return proto.CompactTextString(m) }
func (*ListMembersRequest) ProtoMessage()    {}
func (*ListMembersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1bfa3bac23793f1f, []int{0}
}

func (m *ListMembersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMembersRequest.Unmarshal(m, b)
}
func (m *ListMembersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMembersRequest.Marshal(b, m, deterministic)
}
func (m *ListMembersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMembersRequest.Merge(m, src)
}
func (m *ListMembersRequest) XXX_Size() int {
	return xxx_messageInfo_ListMembersRequest.Size(m)
}
func (m *ListMembersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMembersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListMembersRequest proto.InternalMessageInfo

type ListMembersResponse struct {
	Status               *rpc.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Members              []*Member   `protobuf:"bytes,2,rep,name=members,proto3" json:"members,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ListMembersResponse) Reset()         { *m = ListMembersResponse{} }
func (m *ListMembersResponse) String() string { return proto.CompactTextString(m) }
func (*ListMembersResponse) ProtoMessage()    {}
func (*ListMembersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1bfa3bac23793f1f, []int{1}
}

func (m *ListMembersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMembersResponse.Unmarshal(m, b)
}
func (m *ListMembersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMembersResponse.Marshal(b, m, deterministic)
}
func (m *ListMembersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMembersResponse.Merge(m, src)
}
func (m *ListMembersResponse) XXX_Size() int {
	return xxx_messageInfo_ListMembersResponse.Size(m)
}
func (m *ListMembersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMembersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListMembersResponse proto.InternalMessageInfo

func (m *ListMembersResponse) GetStatus() *rpc.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ListMembersResponse) GetMembers() []*Member {
	if m != nil {
		return m.Members
	}
	return nil
}

type GetDocumentationRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDocumentationRequest) Reset()         { *m = GetDocumentationRequest{} }
func (m *GetDocumentationRequest) String() string { return proto.CompactTextString(m) }
func (*GetDocumentationRequest) ProtoMessage()    {}
func (*GetDocumentationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1bfa3bac23793f1f, []int{2}
}

func (m *GetDocumentationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDocumentationRequest.Unmarshal(m, b)
}
func (m *GetDocumentationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDocumentationRequest.Marshal(b, m, deterministic)
}
func (m *GetDocumentationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDocumentationRequest.Merge(m, src)
}
func (m *GetDocumentationRequest) XXX_Size() int {
	return xxx_messageInfo_GetDocumentationRequest.Size(m)
}
func (m *GetDocumentationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDocumentationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDocumentationRequest proto.InternalMessageInfo

type GetDocumentationResponse struct {
	Status               *rpc.Status    `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Documentation        *Documentation `protobuf:"bytes,2,opt,name=documentation,proto3" json:"documentation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetDocumentationResponse) Reset()         { *m = GetDocumentationResponse{} }
func (m *GetDocumentationResponse) String() string { return proto.CompactTextString(m) }
func (*GetDocumentationResponse) ProtoMessage()    {}
func (*GetDocumentationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1bfa3bac23793f1f, []int{3}
}

func (m *GetDocumentationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDocumentationResponse.Unmarshal(m, b)
}
func (m *GetDocumentationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDocumentationResponse.Marshal(b, m, deterministic)
}
func (m *GetDocumentationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDocumentationResponse.Merge(m, src)
}
func (m *GetDocumentationResponse) XXX_Size() int {
	return xxx_messageInfo_GetDocumentationResponse.Size(m)
}
func (m *GetDocumentationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDocumentationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDocumentationResponse proto.InternalMessageInfo

func (m *GetDocumentationResponse) GetStatus() *rpc.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *GetDocumentationResponse) GetDocumentation() *Documentation {
	if m != nil {
		return m.Documentation
	}
	return nil
}

func init() {
	proto.RegisterType((*ListMembersRequest)(nil), "cernbox.aboutv1.ListMembersRequest")
	proto.RegisterType((*ListMembersResponse)(nil), "cernbox.aboutv1.ListMembersResponse")
	proto.RegisterType((*GetDocumentationRequest)(nil), "cernbox.aboutv1.GetDocumentationRequest")
	proto.RegisterType((*GetDocumentationResponse)(nil), "cernbox.aboutv1.GetDocumentationResponse")
}

func init() { proto.RegisterFile("cernbox/about/v1/about.proto", fileDescriptor_1bfa3bac23793f1f) }

var fileDescriptor_1bfa3bac23793f1f = []byte{
	// 389 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcd, 0x4a, 0xeb, 0x50,
	0x14, 0x85, 0x49, 0x2e, 0xf4, 0x72, 0x4f, 0xaf, 0x54, 0x4f, 0x84, 0xa6, 0xa1, 0x68, 0x89, 0x0e,
	0x2a, 0x42, 0x42, 0xea, 0x13, 0x34, 0xad, 0x38, 0xf1, 0xa7, 0xa4, 0x20, 0xa2, 0x4e, 0x92, 0xb8,
	0x29, 0x01, 0x93, 0x1d, 0xcf, 0x39, 0x29, 0x1d, 0x3b, 0x72, 0xa2, 0x13, 0xdf, 0xc0, 0xa1, 0x8f,
	0xe2, 0xd4, 0x57, 0xf0, 0x41, 0xa4, 0xf9, 0xd1, 0xa6, 0x41, 0x11, 0x67, 0x61, 0xaf, 0xb5, 0x57,
	0xbe, 0xb3, 0xd8, 0xa4, 0xed, 0x03, 0x8b, 0x3c, 0x9c, 0x99, 0xae, 0x87, 0x89, 0x30, 0xa7, 0x56,
	0xf6, 0x61, 0xc4, 0x0c, 0x05, 0xd2, 0x46, 0xae, 0x1a, 0xe9, 0x70, 0x6a, 0x69, 0x9d, 0x8a, 0x9d,
	0x01, 0xc7, 0x84, 0xf9, 0xc0, 0xb3, 0x15, 0x4d, 0x2d, 0x1c, 0x2c, 0xf6, 0x4d, 0x2e, 0x5c, 0x91,
	0x14, 0x4a, 0x7b, 0x82, 0x38, 0xb9, 0x06, 0xd3, 0x8d, 0x03, 0xd3, 0x8d, 0x22, 0x14, 0xae, 0x08,
	0x30, 0xca, 0x55, 0x7d, 0x9d, 0xd0, 0xc3, 0x80, 0x8b, 0x23, 0x08, 0x3d, 0x60, 0xdc, 0x81, 0x9b,
	0x04, 0xb8, 0xd0, 0x13, 0xa2, 0x94, 0xa6, 0x3c, 0xc6, 0x88, 0x03, 0xdd, 0x25, 0xb5, 0x2c, 0x5a,
	0x95, 0x3a, 0x52, 0xb7, 0xde, 0x53, 0x8c, 0x02, 0x94, 0xc5, 0xbe, 0x31, 0x4e, 0x25, 0x27, 0xb7,
	0x50, 0x8b, 0xfc, 0x0d, 0xb3, 0x7d, 0x55, 0xee, 0xfc, 0xe9, 0xd6, 0x7b, 0x4d, 0x63, 0xe9, 0x59,
	0x46, 0x96, 0xef, 0x14, 0x3e, 0xbd, 0x45, 0x9a, 0x07, 0x20, 0x86, 0xe8, 0x27, 0x21, 0x44, 0x19,
	0x67, 0x41, 0x74, 0x2f, 0x11, 0xb5, 0xaa, 0xfd, 0x86, 0x6b, 0x48, 0x56, 0xae, 0x16, 0x53, 0x54,
	0x39, 0xdd, 0xd9, 0xa8, 0xd0, 0x95, 0xff, 0x55, 0x5e, 0xea, 0x3d, 0xc8, 0xe4, 0x7f, 0x7f, 0x6e,
	0x1c, 0x03, 0x9b, 0x06, 0x3e, 0x50, 0x4e, 0xea, 0x0b, 0x95, 0xd1, 0xad, 0x4a, 0x5c, 0xb5, 0x66,
	0x6d, 0xfb, 0x7b, 0x53, 0xf6, 0x3a, 0xbd, 0x75, 0xfb, 0xfa, 0xf6, 0x28, 0x2b, 0x74, 0xed, 0xe3,
	0x4c, 0xcc, 0xbc, 0x30, 0x7a, 0x27, 0x91, 0xd5, 0xe5, 0x56, 0x68, 0xb7, 0x92, 0xfa, 0x45, 0xa9,
	0xda, 0xce, 0x0f, 0x9c, 0x39, 0xc4, 0x66, 0x0a, 0xd1, 0xa2, 0xcd, 0x4f, 0x88, 0x52, 0x21, 0xf6,
	0x05, 0x51, 0x7c, 0x0c, 0x97, 0x03, 0x6d, 0x92, 0x96, 0x34, 0x9a, 0xdf, 0xda, 0x48, 0x3a, 0xff,
	0x97, 0x8f, 0x63, 0xef, 0x49, 0xae, 0x0d, 0xec, 0x93, 0xb3, 0xbe, 0xfd, 0x2c, 0x37, 0x06, 0xfb,
	0xce, 0xb1, 0x8d, 0x33, 0x23, 0xb5, 0x9e, 0x5a, 0x2f, 0x72, 0x63, 0x00, 0x2c, 0xb2, 0x71, 0x76,
	0x99, 0x4f, 0xbc, 0x5a, 0x7a, 0xac, 0x7b, 0xef, 0x01, 0x00, 0x00, 0xff, 0xff, 0xbb, 0xfb, 0x64,
	0xa5, 0x37, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AboutServiceClient is the client API for AboutService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AboutServiceClient interface {
	ListMembers(ctx context.Context, in *ListMembersRequest, opts ...grpc.CallOption) (*ListMembersResponse, error)
	GetDocumentation(ctx context.Context, in *GetDocumentationRequest, opts ...grpc.CallOption) (*GetDocumentationResponse, error)
}

type aboutServiceClient struct {
	cc *grpc.ClientConn
}

func NewAboutServiceClient(cc *grpc.ClientConn) AboutServiceClient {
	return &aboutServiceClient{cc}
}

func (c *aboutServiceClient) ListMembers(ctx context.Context, in *ListMembersRequest, opts ...grpc.CallOption) (*ListMembersResponse, error) {
	out := new(ListMembersResponse)
	err := c.cc.Invoke(ctx, "/cernbox.aboutv1.AboutService/ListMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aboutServiceClient) GetDocumentation(ctx context.Context, in *GetDocumentationRequest, opts ...grpc.CallOption) (*GetDocumentationResponse, error) {
	out := new(GetDocumentationResponse)
	err := c.cc.Invoke(ctx, "/cernbox.aboutv1.AboutService/GetDocumentation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AboutServiceServer is the server API for AboutService service.
type AboutServiceServer interface {
	ListMembers(context.Context, *ListMembersRequest) (*ListMembersResponse, error)
	GetDocumentation(context.Context, *GetDocumentationRequest) (*GetDocumentationResponse, error)
}

func RegisterAboutServiceServer(s *grpc.Server, srv AboutServiceServer) {
	s.RegisterService(&_AboutService_serviceDesc, srv)
}

func _AboutService_ListMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMembersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AboutServiceServer).ListMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cernbox.aboutv1.AboutService/ListMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AboutServiceServer).ListMembers(ctx, req.(*ListMembersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AboutService_GetDocumentation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDocumentationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AboutServiceServer).GetDocumentation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cernbox.aboutv1.AboutService/GetDocumentation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AboutServiceServer).GetDocumentation(ctx, req.(*GetDocumentationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AboutService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cernbox.aboutv1.AboutService",
	HandlerType: (*AboutServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListMembers",
			Handler:    _AboutService_ListMembers_Handler,
		},
		{
			MethodName: "GetDocumentation",
			Handler:    _AboutService_GetDocumentation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cernbox/about/v1/about.proto",
}