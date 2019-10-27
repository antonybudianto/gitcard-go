// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/github.proto

package protos

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type GithubRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GithubRequest) Reset()         { *m = GithubRequest{} }
func (m *GithubRequest) String() string { return proto.CompactTextString(m) }
func (*GithubRequest) ProtoMessage()    {}
func (*GithubRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_da726d3ccdb16248, []int{0}
}

func (m *GithubRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GithubRequest.Unmarshal(m, b)
}
func (m *GithubRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GithubRequest.Marshal(b, m, deterministic)
}
func (m *GithubRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GithubRequest.Merge(m, src)
}
func (m *GithubRequest) XXX_Size() int {
	return xxx_messageInfo_GithubRequest.Size(m)
}
func (m *GithubRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GithubRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GithubRequest proto.InternalMessageInfo

func (m *GithubRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type GithubResponse struct {
	Username             string           `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Starcount            int32            `protobuf:"varint,2,opt,name=starcount,proto3" json:"starcount,omitempty"`
	Repocount            int32            `protobuf:"varint,3,opt,name=repocount,proto3" json:"repocount,omitempty"`
	Forkcount            int32            `protobuf:"varint,4,opt,name=forkcount,proto3" json:"forkcount,omitempty"`
	Langmap              map[string]int32 `protobuf:"bytes,5,rep,name=langmap,proto3" json:"langmap,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GithubResponse) Reset()         { *m = GithubResponse{} }
func (m *GithubResponse) String() string { return proto.CompactTextString(m) }
func (*GithubResponse) ProtoMessage()    {}
func (*GithubResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_da726d3ccdb16248, []int{1}
}

func (m *GithubResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GithubResponse.Unmarshal(m, b)
}
func (m *GithubResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GithubResponse.Marshal(b, m, deterministic)
}
func (m *GithubResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GithubResponse.Merge(m, src)
}
func (m *GithubResponse) XXX_Size() int {
	return xxx_messageInfo_GithubResponse.Size(m)
}
func (m *GithubResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GithubResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GithubResponse proto.InternalMessageInfo

func (m *GithubResponse) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *GithubResponse) GetStarcount() int32 {
	if m != nil {
		return m.Starcount
	}
	return 0
}

func (m *GithubResponse) GetRepocount() int32 {
	if m != nil {
		return m.Repocount
	}
	return 0
}

func (m *GithubResponse) GetForkcount() int32 {
	if m != nil {
		return m.Forkcount
	}
	return 0
}

func (m *GithubResponse) GetLangmap() map[string]int32 {
	if m != nil {
		return m.Langmap
	}
	return nil
}

func init() {
	proto.RegisterType((*GithubRequest)(nil), "protos.GithubRequest")
	proto.RegisterType((*GithubResponse)(nil), "protos.GithubResponse")
	proto.RegisterMapType((map[string]int32)(nil), "protos.GithubResponse.LangmapEntry")
}

func init() { proto.RegisterFile("protos/github.proto", fileDescriptor_da726d3ccdb16248) }

var fileDescriptor_da726d3ccdb16248 = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0x2f, 0xd6, 0x4f, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x03, 0xf3, 0x84, 0xd8, 0x20, 0x82,
	0x4a, 0xda, 0x5c, 0xbc, 0xee, 0x60, 0xf1, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x29,
	0x2e, 0x8e, 0xd2, 0xe2, 0xd4, 0xa2, 0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce,
	0x20, 0x38, 0x5f, 0xa9, 0x91, 0x89, 0x8b, 0x0f, 0xa6, 0xba, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x15,
	0x9f, 0x72, 0x21, 0x19, 0x2e, 0xce, 0xe2, 0x92, 0xc4, 0xa2, 0xe4, 0xfc, 0xd2, 0xbc, 0x12, 0x09,
	0x26, 0x05, 0x46, 0x0d, 0xd6, 0x20, 0x84, 0x00, 0x48, 0xb6, 0x28, 0xb5, 0x20, 0x1f, 0x22, 0xcb,
	0x0c, 0x91, 0x85, 0x0b, 0x80, 0x64, 0xd3, 0xf2, 0x8b, 0xb2, 0x21, 0xb2, 0x2c, 0x10, 0x59, 0xb8,
	0x80, 0x90, 0x2d, 0x17, 0x7b, 0x4e, 0x62, 0x5e, 0x7a, 0x6e, 0x62, 0x81, 0x04, 0xab, 0x02, 0xb3,
	0x06, 0xb7, 0x91, 0x32, 0xc4, 0x5b, 0xc5, 0x7a, 0xa8, 0xce, 0xd3, 0xf3, 0x81, 0xa8, 0x72, 0xcd,
	0x2b, 0x29, 0xaa, 0x0c, 0x82, 0xe9, 0x91, 0xb2, 0xe2, 0xe2, 0x41, 0x96, 0x10, 0x12, 0xe0, 0x62,
	0xce, 0x4e, 0xad, 0x84, 0xba, 0x1f, 0xc4, 0x14, 0x12, 0xe1, 0x62, 0x2d, 0x4b, 0xcc, 0x29, 0x4d,
	0x85, 0x3a, 0x1b, 0xc2, 0xb1, 0x62, 0xb2, 0x60, 0x34, 0x0a, 0x86, 0x05, 0x58, 0x70, 0x6a, 0x51,
	0x59, 0x66, 0x72, 0xaa, 0x90, 0x13, 0x17, 0xbf, 0x5b, 0x6a, 0x49, 0x72, 0x86, 0x53, 0x65, 0x28,
	0xcc, 0xe3, 0xa2, 0xe8, 0xae, 0x01, 0x07, 0xad, 0x94, 0x18, 0x76, 0x47, 0x2a, 0x31, 0x24, 0x41,
	0x62, 0xc3, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x64, 0x60, 0x27, 0x72, 0xab, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GithubServiceClient is the client API for GithubService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GithubServiceClient interface {
	FetchByUsername(ctx context.Context, in *GithubRequest, opts ...grpc.CallOption) (*GithubResponse, error)
}

type githubServiceClient struct {
	cc *grpc.ClientConn
}

func NewGithubServiceClient(cc *grpc.ClientConn) GithubServiceClient {
	return &githubServiceClient{cc}
}

func (c *githubServiceClient) FetchByUsername(ctx context.Context, in *GithubRequest, opts ...grpc.CallOption) (*GithubResponse, error) {
	out := new(GithubResponse)
	err := c.cc.Invoke(ctx, "/protos.GithubService/FetchByUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GithubServiceServer is the server API for GithubService service.
type GithubServiceServer interface {
	FetchByUsername(context.Context, *GithubRequest) (*GithubResponse, error)
}

// UnimplementedGithubServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGithubServiceServer struct {
}

func (*UnimplementedGithubServiceServer) FetchByUsername(ctx context.Context, req *GithubRequest) (*GithubResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchByUsername not implemented")
}

func RegisterGithubServiceServer(s *grpc.Server, srv GithubServiceServer) {
	s.RegisterService(&_GithubService_serviceDesc, srv)
}

func _GithubService_FetchByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GithubRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GithubServiceServer).FetchByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.GithubService/FetchByUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GithubServiceServer).FetchByUsername(ctx, req.(*GithubRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GithubService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.GithubService",
	HandlerType: (*GithubServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchByUsername",
			Handler:    _GithubService_FetchByUsername_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/github.proto",
}