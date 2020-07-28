// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth/external/auth_api.proto

package externalv1beta1

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/mwitkow/go-proto-validators"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type SignUpRequest struct {
	// User email.
	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	// User password.
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignUpRequest) Reset()         { *m = SignUpRequest{} }
func (m *SignUpRequest) String() string { return proto.CompactTextString(m) }
func (*SignUpRequest) ProtoMessage()    {}
func (*SignUpRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e0315b5819ba925, []int{0}
}

func (m *SignUpRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignUpRequest.Unmarshal(m, b)
}
func (m *SignUpRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignUpRequest.Marshal(b, m, deterministic)
}
func (m *SignUpRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignUpRequest.Merge(m, src)
}
func (m *SignUpRequest) XXX_Size() int {
	return xxx_messageInfo_SignUpRequest.Size(m)
}
func (m *SignUpRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignUpRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignUpRequest proto.InternalMessageInfo

func (m *SignUpRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *SignUpRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type SignUpResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignUpResponse) Reset()         { *m = SignUpResponse{} }
func (m *SignUpResponse) String() string { return proto.CompactTextString(m) }
func (*SignUpResponse) ProtoMessage()    {}
func (*SignUpResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e0315b5819ba925, []int{1}
}

func (m *SignUpResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignUpResponse.Unmarshal(m, b)
}
func (m *SignUpResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignUpResponse.Marshal(b, m, deterministic)
}
func (m *SignUpResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignUpResponse.Merge(m, src)
}
func (m *SignUpResponse) XXX_Size() int {
	return xxx_messageInfo_SignUpResponse.Size(m)
}
func (m *SignUpResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SignUpResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SignUpResponse proto.InternalMessageInfo

type SignInRequest struct {
	// User email.
	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	// User password.
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignInRequest) Reset()         { *m = SignInRequest{} }
func (m *SignInRequest) String() string { return proto.CompactTextString(m) }
func (*SignInRequest) ProtoMessage()    {}
func (*SignInRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e0315b5819ba925, []int{2}
}

func (m *SignInRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignInRequest.Unmarshal(m, b)
}
func (m *SignInRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignInRequest.Marshal(b, m, deterministic)
}
func (m *SignInRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignInRequest.Merge(m, src)
}
func (m *SignInRequest) XXX_Size() int {
	return xxx_messageInfo_SignInRequest.Size(m)
}
func (m *SignInRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignInRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignInRequest proto.InternalMessageInfo

func (m *SignInRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *SignInRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type SignInResponse struct {
	// Session ID.
	SessionId string `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	// Session expiration time.
	ExpireTime           *timestamp.Timestamp `protobuf:"bytes,2,opt,name=expire_time,json=expireTime,proto3" json:"expire_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SignInResponse) Reset()         { *m = SignInResponse{} }
func (m *SignInResponse) String() string { return proto.CompactTextString(m) }
func (*SignInResponse) ProtoMessage()    {}
func (*SignInResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e0315b5819ba925, []int{3}
}

func (m *SignInResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignInResponse.Unmarshal(m, b)
}
func (m *SignInResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignInResponse.Marshal(b, m, deterministic)
}
func (m *SignInResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignInResponse.Merge(m, src)
}
func (m *SignInResponse) XXX_Size() int {
	return xxx_messageInfo_SignInResponse.Size(m)
}
func (m *SignInResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SignInResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SignInResponse proto.InternalMessageInfo

func (m *SignInResponse) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *SignInResponse) GetExpireTime() *timestamp.Timestamp {
	if m != nil {
		return m.ExpireTime
	}
	return nil
}

type RefreshSessionRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RefreshSessionRequest) Reset()         { *m = RefreshSessionRequest{} }
func (m *RefreshSessionRequest) String() string { return proto.CompactTextString(m) }
func (*RefreshSessionRequest) ProtoMessage()    {}
func (*RefreshSessionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e0315b5819ba925, []int{4}
}

func (m *RefreshSessionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RefreshSessionRequest.Unmarshal(m, b)
}
func (m *RefreshSessionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RefreshSessionRequest.Marshal(b, m, deterministic)
}
func (m *RefreshSessionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefreshSessionRequest.Merge(m, src)
}
func (m *RefreshSessionRequest) XXX_Size() int {
	return xxx_messageInfo_RefreshSessionRequest.Size(m)
}
func (m *RefreshSessionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RefreshSessionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RefreshSessionRequest proto.InternalMessageInfo

type RefreshSessionResponse struct {
	// Session expiration time.
	ExpireTime           *timestamp.Timestamp `protobuf:"bytes,1,opt,name=expire_time,json=expireTime,proto3" json:"expire_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *RefreshSessionResponse) Reset()         { *m = RefreshSessionResponse{} }
func (m *RefreshSessionResponse) String() string { return proto.CompactTextString(m) }
func (*RefreshSessionResponse) ProtoMessage()    {}
func (*RefreshSessionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e0315b5819ba925, []int{5}
}

func (m *RefreshSessionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RefreshSessionResponse.Unmarshal(m, b)
}
func (m *RefreshSessionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RefreshSessionResponse.Marshal(b, m, deterministic)
}
func (m *RefreshSessionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefreshSessionResponse.Merge(m, src)
}
func (m *RefreshSessionResponse) XXX_Size() int {
	return xxx_messageInfo_RefreshSessionResponse.Size(m)
}
func (m *RefreshSessionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RefreshSessionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RefreshSessionResponse proto.InternalMessageInfo

func (m *RefreshSessionResponse) GetExpireTime() *timestamp.Timestamp {
	if m != nil {
		return m.ExpireTime
	}
	return nil
}

func init() {
	proto.RegisterType((*SignUpRequest)(nil), "percona.platform.auth.external.v1beta1.SignUpRequest")
	proto.RegisterType((*SignUpResponse)(nil), "percona.platform.auth.external.v1beta1.SignUpResponse")
	proto.RegisterType((*SignInRequest)(nil), "percona.platform.auth.external.v1beta1.SignInRequest")
	proto.RegisterType((*SignInResponse)(nil), "percona.platform.auth.external.v1beta1.SignInResponse")
	proto.RegisterType((*RefreshSessionRequest)(nil), "percona.platform.auth.external.v1beta1.RefreshSessionRequest")
	proto.RegisterType((*RefreshSessionResponse)(nil), "percona.platform.auth.external.v1beta1.RefreshSessionResponse")
}

func init() {
	proto.RegisterFile("auth/external/auth_api.proto", fileDescriptor_7e0315b5819ba925)
}

var fileDescriptor_7e0315b5819ba925 = []byte{
	// 392 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x41, 0x6b, 0xdb, 0x30,
	0x1c, 0xc5, 0x71, 0xc6, 0xb2, 0x45, 0x61, 0xd9, 0x26, 0xd8, 0x16, 0x4c, 0x46, 0x8a, 0x0f, 0xa5,
	0x97, 0xc8, 0x24, 0xa5, 0xb9, 0x94, 0x16, 0x92, 0x9b, 0x6f, 0xad, 0xd3, 0x40, 0xe9, 0x25, 0xc8,
	0xb1, 0x62, 0x8b, 0xda, 0x96, 0x2a, 0xc9, 0x71, 0xbe, 0x44, 0xe9, 0x47, 0x2c, 0xf4, 0x93, 0x14,
	0x5b, 0x76, 0x68, 0x42, 0x0f, 0x71, 0xe9, 0xcd, 0xfa, 0xfb, 0xe9, 0xfd, 0xa4, 0xf7, 0x04, 0x7a,
	0x38, 0x55, 0xa1, 0x4d, 0x36, 0x8a, 0x88, 0x04, 0x47, 0x76, 0xbe, 0x5a, 0x60, 0x4e, 0x11, 0x17,
	0x4c, 0x31, 0x78, 0xcc, 0x89, 0x58, 0xb2, 0x04, 0x23, 0x1e, 0x61, 0xb5, 0x62, 0x22, 0x46, 0xb9,
	0x00, 0x55, 0x72, 0xb4, 0x1e, 0x7a, 0x44, 0xe1, 0xa1, 0x39, 0x0e, 0xa8, 0x0a, 0x53, 0x0f, 0x2d,
	0x59, 0x6c, 0xc7, 0x19, 0x55, 0xf7, 0x2c, 0xb3, 0x03, 0x36, 0x28, 0x4c, 0x06, 0x6b, 0x1c, 0x51,
	0x1f, 0x2b, 0x26, 0xa4, 0xbd, 0xfd, 0xd4, 0xfe, 0x66, 0x3f, 0x60, 0x2c, 0x88, 0x88, 0x5d, 0xac,
	0xbc, 0x74, 0x65, 0x2b, 0x1a, 0x13, 0xa9, 0x70, 0xcc, 0xb5, 0xc0, 0xba, 0x06, 0x3f, 0x66, 0x34,
	0x48, 0xe6, 0xdc, 0x25, 0x0f, 0x29, 0x91, 0x0a, 0xf6, 0xc0, 0x57, 0x12, 0x63, 0x1a, 0x75, 0x8d,
	0x23, 0xe3, 0xa4, 0x35, 0x6d, 0xbe, 0x3c, 0xf7, 0x1b, 0xb7, 0x86, 0xab, 0x87, 0xd0, 0x02, 0xdf,
	0x39, 0x96, 0x32, 0x63, 0xc2, 0xef, 0x36, 0x76, 0x04, 0xdb, 0xb9, 0xf5, 0x0b, 0x74, 0x2a, 0x4b,
	0xc9, 0x59, 0x22, 0x49, 0x05, 0x71, 0x92, 0xcf, 0x83, 0x44, 0x1a, 0x92, 0x5b, 0x6a, 0x08, 0xfc,
	0x0f, 0x80, 0x24, 0x52, 0x52, 0x96, 0x2c, 0xa8, 0xaf, 0x8d, 0xdd, 0x56, 0x39, 0x71, 0x7c, 0x78,
	0x0e, 0xda, 0x64, 0xc3, 0xa9, 0x20, 0x8b, 0x3c, 0x82, 0xc2, 0xb7, 0x3d, 0x32, 0x91, 0xce, 0x07,
	0x55, 0xf9, 0xa0, 0x9b, 0x2a, 0x1f, 0x17, 0x68, 0x79, 0x3e, 0xb0, 0xfe, 0x81, 0x3f, 0x2e, 0x59,
	0x09, 0x22, 0xc3, 0x99, 0x36, 0x2c, 0x2f, 0x62, 0xcd, 0xc1, 0xdf, 0xfd, 0x1f, 0xe5, 0x71, 0xf6,
	0x78, 0x46, 0x1d, 0xde, 0xe8, 0xf1, 0x0b, 0xf8, 0x36, 0x49, 0x55, 0x38, 0xb9, 0x72, 0x60, 0x06,
	0x9a, 0x3a, 0x4e, 0x78, 0x86, 0x0e, 0x7b, 0x2d, 0x68, 0xa7, 0x51, 0x73, 0x5c, 0x77, 0x5b, 0x79,
	0x83, 0x12, 0xec, 0x24, 0xf5, 0xc0, 0xdb, 0x96, 0xeb, 0x81, 0xdf, 0x34, 0xf9, 0x64, 0x80, 0xce,
	0x6e, 0xaa, 0xf0, 0xe2, 0x50, 0xab, 0x77, 0x6b, 0x32, 0x2f, 0x3f, 0xba, 0x5d, 0x9f, 0x68, 0xfa,
	0xfb, 0xee, 0x67, 0x25, 0x2d, 0x95, 0x5e, 0xb3, 0xa8, 0xf0, 0xf4, 0x35, 0x00, 0x00, 0xff, 0xff,
	0x65, 0x9f, 0x8d, 0x1c, 0xe0, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthAPIClient is the client API for AuthAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthAPIClient interface {
	// SignUp creates new user with given email and password.
	SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error)
	// SignIn checks user authentication, creates session and returns session ID.
	SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error)
	// RefreshSession refreshes session timeout.
	RefreshSession(ctx context.Context, in *RefreshSessionRequest, opts ...grpc.CallOption) (*RefreshSessionResponse, error)
}

type authAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthAPIClient(cc grpc.ClientConnInterface) AuthAPIClient {
	return &authAPIClient{cc}
}

func (c *authAPIClient) SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error) {
	out := new(SignUpResponse)
	err := c.cc.Invoke(ctx, "/percona.platform.auth.external.v1beta1.AuthAPI/SignUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authAPIClient) SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error) {
	out := new(SignInResponse)
	err := c.cc.Invoke(ctx, "/percona.platform.auth.external.v1beta1.AuthAPI/SignIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authAPIClient) RefreshSession(ctx context.Context, in *RefreshSessionRequest, opts ...grpc.CallOption) (*RefreshSessionResponse, error) {
	out := new(RefreshSessionResponse)
	err := c.cc.Invoke(ctx, "/percona.platform.auth.external.v1beta1.AuthAPI/RefreshSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthAPIServer is the server API for AuthAPI service.
type AuthAPIServer interface {
	// SignUp creates new user with given email and password.
	SignUp(context.Context, *SignUpRequest) (*SignUpResponse, error)
	// SignIn checks user authentication, creates session and returns session ID.
	SignIn(context.Context, *SignInRequest) (*SignInResponse, error)
	// RefreshSession refreshes session timeout.
	RefreshSession(context.Context, *RefreshSessionRequest) (*RefreshSessionResponse, error)
}

// UnimplementedAuthAPIServer can be embedded to have forward compatible implementations.
type UnimplementedAuthAPIServer struct {
}

func (*UnimplementedAuthAPIServer) SignUp(ctx context.Context, req *SignUpRequest) (*SignUpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}
func (*UnimplementedAuthAPIServer) SignIn(ctx context.Context, req *SignInRequest) (*SignInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}
func (*UnimplementedAuthAPIServer) RefreshSession(ctx context.Context, req *RefreshSessionRequest) (*RefreshSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshSession not implemented")
}

func RegisterAuthAPIServer(s *grpc.Server, srv AuthAPIServer) {
	s.RegisterService(&_AuthAPI_serviceDesc, srv)
}

func _AuthAPI_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthAPIServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/percona.platform.auth.external.v1beta1.AuthAPI/SignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthAPIServer).SignUp(ctx, req.(*SignUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthAPI_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthAPIServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/percona.platform.auth.external.v1beta1.AuthAPI/SignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthAPIServer).SignIn(ctx, req.(*SignInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthAPI_RefreshSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthAPIServer).RefreshSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/percona.platform.auth.external.v1beta1.AuthAPI/RefreshSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthAPIServer).RefreshSession(ctx, req.(*RefreshSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "percona.platform.auth.external.v1beta1.AuthAPI",
	HandlerType: (*AuthAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignUp",
			Handler:    _AuthAPI_SignUp_Handler,
		},
		{
			MethodName: "SignIn",
			Handler:    _AuthAPI_SignIn_Handler,
		},
		{
			MethodName: "RefreshSession",
			Handler:    _AuthAPI_RefreshSession_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/external/auth_api.proto",
}
