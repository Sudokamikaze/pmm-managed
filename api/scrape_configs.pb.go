// Code generated by protoc-gen-go. DO NOT EDIT.
// source: scrape_configs.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Target health : unknown, down, or up.
type ScrapeTargetHealth_Health int32

const (
	ScrapeTargetHealth_UNKNOWN ScrapeTargetHealth_Health = 0
	ScrapeTargetHealth_DOWN    ScrapeTargetHealth_Health = 1
	ScrapeTargetHealth_UP      ScrapeTargetHealth_Health = 2
)

var ScrapeTargetHealth_Health_name = map[int32]string{
	0: "UNKNOWN",
	1: "DOWN",
	2: "UP",
}
var ScrapeTargetHealth_Health_value = map[string]int32{
	"UNKNOWN": 0,
	"DOWN":    1,
	"UP":      2,
}

func (x ScrapeTargetHealth_Health) String() string {
	return proto.EnumName(ScrapeTargetHealth_Health_name, int32(x))
}
func (ScrapeTargetHealth_Health) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{5, 0} }

type LabelPair struct {
	// Label name
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Label value
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *LabelPair) Reset()                    { *m = LabelPair{} }
func (m *LabelPair) String() string            { return proto.CompactTextString(m) }
func (*LabelPair) ProtoMessage()               {}
func (*LabelPair) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *LabelPair) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LabelPair) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type StaticConfig struct {
	// Hostnames or IPs followed by an optional port number: "1.2.3.4:9090"
	Targets []string `protobuf:"bytes,1,rep,name=targets" json:"targets,omitempty"`
	// Labels assigned to all metrics scraped from the targets
	Labels []*LabelPair `protobuf:"bytes,2,rep,name=labels" json:"labels,omitempty"`
}

func (m *StaticConfig) Reset()                    { *m = StaticConfig{} }
func (m *StaticConfig) String() string            { return proto.CompactTextString(m) }
func (*StaticConfig) ProtoMessage()               {}
func (*StaticConfig) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *StaticConfig) GetTargets() []string {
	if m != nil {
		return m.Targets
	}
	return nil
}

func (m *StaticConfig) GetLabels() []*LabelPair {
	if m != nil {
		return m.Labels
	}
	return nil
}

type BasicAuth struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *BasicAuth) Reset()                    { *m = BasicAuth{} }
func (m *BasicAuth) String() string            { return proto.CompactTextString(m) }
func (*BasicAuth) ProtoMessage()               {}
func (*BasicAuth) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *BasicAuth) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *BasicAuth) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type TLSConfig struct {
	InsecureSkipVerify bool `protobuf:"varint,5,opt,name=insecure_skip_verify,json=insecureSkipVerify" json:"insecure_skip_verify,omitempty"`
}

func (m *TLSConfig) Reset()                    { *m = TLSConfig{} }
func (m *TLSConfig) String() string            { return proto.CompactTextString(m) }
func (*TLSConfig) ProtoMessage()               {}
func (*TLSConfig) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *TLSConfig) GetInsecureSkipVerify() bool {
	if m != nil {
		return m.InsecureSkipVerify
	}
	return false
}

type ScrapeConfig struct {
	// The job name assigned to scraped metrics by default: "example-job" (required)
	JobName string `protobuf:"bytes,1,opt,name=job_name,json=jobName" json:"job_name,omitempty"`
	// How frequently to scrape targets from this job: "10s"
	ScrapeInterval string `protobuf:"bytes,2,opt,name=scrape_interval,json=scrapeInterval" json:"scrape_interval,omitempty"`
	// Per-scrape timeout when scraping this job: "5s"
	ScrapeTimeout string `protobuf:"bytes,3,opt,name=scrape_timeout,json=scrapeTimeout" json:"scrape_timeout,omitempty"`
	// The HTTP resource path on which to fetch metrics from targets: "/metrics"
	MetricsPath string `protobuf:"bytes,4,opt,name=metrics_path,json=metricsPath" json:"metrics_path,omitempty"`
	// Configures the protocol scheme used for requests: "http" or "https"
	Scheme string `protobuf:"bytes,5,opt,name=scheme" json:"scheme,omitempty"`
	// Sets the `Authorization` header on every scrape request with the configured username and password
	BasicAuth *BasicAuth `protobuf:"bytes,6,opt,name=basic_auth,json=basicAuth" json:"basic_auth,omitempty"`
	// Configures the scrape request's TLS settings
	TlsConfig *TLSConfig `protobuf:"bytes,7,opt,name=tls_config,json=tlsConfig" json:"tls_config,omitempty"`
	// List of labeled statically configured targets for this job
	StaticConfigs []*StaticConfig `protobuf:"bytes,8,rep,name=static_configs,json=staticConfigs" json:"static_configs,omitempty"`
}

func (m *ScrapeConfig) Reset()                    { *m = ScrapeConfig{} }
func (m *ScrapeConfig) String() string            { return proto.CompactTextString(m) }
func (*ScrapeConfig) ProtoMessage()               {}
func (*ScrapeConfig) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

func (m *ScrapeConfig) GetJobName() string {
	if m != nil {
		return m.JobName
	}
	return ""
}

func (m *ScrapeConfig) GetScrapeInterval() string {
	if m != nil {
		return m.ScrapeInterval
	}
	return ""
}

func (m *ScrapeConfig) GetScrapeTimeout() string {
	if m != nil {
		return m.ScrapeTimeout
	}
	return ""
}

func (m *ScrapeConfig) GetMetricsPath() string {
	if m != nil {
		return m.MetricsPath
	}
	return ""
}

func (m *ScrapeConfig) GetScheme() string {
	if m != nil {
		return m.Scheme
	}
	return ""
}

func (m *ScrapeConfig) GetBasicAuth() *BasicAuth {
	if m != nil {
		return m.BasicAuth
	}
	return nil
}

func (m *ScrapeConfig) GetTlsConfig() *TLSConfig {
	if m != nil {
		return m.TlsConfig
	}
	return nil
}

func (m *ScrapeConfig) GetStaticConfigs() []*StaticConfig {
	if m != nil {
		return m.StaticConfigs
	}
	return nil
}

// ScrapeTargetHealth represents Prometheus scrape target health: unknown, down, or up.
type ScrapeTargetHealth struct {
	// Original scrape job name
	JobName string `protobuf:"bytes,1,opt,name=job_name,json=jobName" json:"job_name,omitempty"`
	// "job" label value, may be different from job_name due to relabeling
	Job string `protobuf:"bytes,2,opt,name=job" json:"job,omitempty"`
	// Original target
	Target string `protobuf:"bytes,3,opt,name=target" json:"target,omitempty"`
	// "instance" label value, may be different from target due to relabeling
	Instance string                    `protobuf:"bytes,4,opt,name=instance" json:"instance,omitempty"`
	Health   ScrapeTargetHealth_Health `protobuf:"varint,5,opt,name=health,enum=api.ScrapeTargetHealth_Health" json:"health,omitempty"`
}

func (m *ScrapeTargetHealth) Reset()                    { *m = ScrapeTargetHealth{} }
func (m *ScrapeTargetHealth) String() string            { return proto.CompactTextString(m) }
func (*ScrapeTargetHealth) ProtoMessage()               {}
func (*ScrapeTargetHealth) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{5} }

func (m *ScrapeTargetHealth) GetJobName() string {
	if m != nil {
		return m.JobName
	}
	return ""
}

func (m *ScrapeTargetHealth) GetJob() string {
	if m != nil {
		return m.Job
	}
	return ""
}

func (m *ScrapeTargetHealth) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *ScrapeTargetHealth) GetInstance() string {
	if m != nil {
		return m.Instance
	}
	return ""
}

func (m *ScrapeTargetHealth) GetHealth() ScrapeTargetHealth_Health {
	if m != nil {
		return m.Health
	}
	return ScrapeTargetHealth_UNKNOWN
}

type ScrapeConfigsListRequest struct {
}

func (m *ScrapeConfigsListRequest) Reset()                    { *m = ScrapeConfigsListRequest{} }
func (m *ScrapeConfigsListRequest) String() string            { return proto.CompactTextString(m) }
func (*ScrapeConfigsListRequest) ProtoMessage()               {}
func (*ScrapeConfigsListRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{6} }

type ScrapeConfigsListResponse struct {
	ScrapeConfigs []*ScrapeConfig `protobuf:"bytes,1,rep,name=scrape_configs,json=scrapeConfigs" json:"scrape_configs,omitempty"`
	// Scrape targets health for all managed scrape jobs
	ScrapeTargetsHealth []*ScrapeTargetHealth `protobuf:"bytes,2,rep,name=scrape_targets_health,json=scrapeTargetsHealth" json:"scrape_targets_health,omitempty"`
}

func (m *ScrapeConfigsListResponse) Reset()                    { *m = ScrapeConfigsListResponse{} }
func (m *ScrapeConfigsListResponse) String() string            { return proto.CompactTextString(m) }
func (*ScrapeConfigsListResponse) ProtoMessage()               {}
func (*ScrapeConfigsListResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{7} }

func (m *ScrapeConfigsListResponse) GetScrapeConfigs() []*ScrapeConfig {
	if m != nil {
		return m.ScrapeConfigs
	}
	return nil
}

func (m *ScrapeConfigsListResponse) GetScrapeTargetsHealth() []*ScrapeTargetHealth {
	if m != nil {
		return m.ScrapeTargetsHealth
	}
	return nil
}

type ScrapeConfigsGetRequest struct {
	JobName string `protobuf:"bytes,1,opt,name=job_name,json=jobName" json:"job_name,omitempty"`
}

func (m *ScrapeConfigsGetRequest) Reset()                    { *m = ScrapeConfigsGetRequest{} }
func (m *ScrapeConfigsGetRequest) String() string            { return proto.CompactTextString(m) }
func (*ScrapeConfigsGetRequest) ProtoMessage()               {}
func (*ScrapeConfigsGetRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{8} }

func (m *ScrapeConfigsGetRequest) GetJobName() string {
	if m != nil {
		return m.JobName
	}
	return ""
}

type ScrapeConfigsGetResponse struct {
	ScrapeConfig *ScrapeConfig `protobuf:"bytes,1,opt,name=scrape_config,json=scrapeConfig" json:"scrape_config,omitempty"`
	// Scrape targets health for this scrape job
	ScrapeTargetsHealth []*ScrapeTargetHealth `protobuf:"bytes,2,rep,name=scrape_targets_health,json=scrapeTargetsHealth" json:"scrape_targets_health,omitempty"`
}

func (m *ScrapeConfigsGetResponse) Reset()                    { *m = ScrapeConfigsGetResponse{} }
func (m *ScrapeConfigsGetResponse) String() string            { return proto.CompactTextString(m) }
func (*ScrapeConfigsGetResponse) ProtoMessage()               {}
func (*ScrapeConfigsGetResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{9} }

func (m *ScrapeConfigsGetResponse) GetScrapeConfig() *ScrapeConfig {
	if m != nil {
		return m.ScrapeConfig
	}
	return nil
}

func (m *ScrapeConfigsGetResponse) GetScrapeTargetsHealth() []*ScrapeTargetHealth {
	if m != nil {
		return m.ScrapeTargetsHealth
	}
	return nil
}

type ScrapeConfigsCreateRequest struct {
	ScrapeConfig *ScrapeConfig `protobuf:"bytes,1,opt,name=scrape_config,json=scrapeConfig" json:"scrape_config,omitempty"`
	// Check that added targets can be scraped from PMM Server
	CheckReachability bool `protobuf:"varint,2,opt,name=check_reachability,json=checkReachability" json:"check_reachability,omitempty"`
}

func (m *ScrapeConfigsCreateRequest) Reset()                    { *m = ScrapeConfigsCreateRequest{} }
func (m *ScrapeConfigsCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*ScrapeConfigsCreateRequest) ProtoMessage()               {}
func (*ScrapeConfigsCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{10} }

func (m *ScrapeConfigsCreateRequest) GetScrapeConfig() *ScrapeConfig {
	if m != nil {
		return m.ScrapeConfig
	}
	return nil
}

func (m *ScrapeConfigsCreateRequest) GetCheckReachability() bool {
	if m != nil {
		return m.CheckReachability
	}
	return false
}

type ScrapeConfigsCreateResponse struct {
}

func (m *ScrapeConfigsCreateResponse) Reset()                    { *m = ScrapeConfigsCreateResponse{} }
func (m *ScrapeConfigsCreateResponse) String() string            { return proto.CompactTextString(m) }
func (*ScrapeConfigsCreateResponse) ProtoMessage()               {}
func (*ScrapeConfigsCreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{11} }

type ScrapeConfigsDeleteRequest struct {
	JobName string `protobuf:"bytes,1,opt,name=job_name,json=jobName" json:"job_name,omitempty"`
}

func (m *ScrapeConfigsDeleteRequest) Reset()                    { *m = ScrapeConfigsDeleteRequest{} }
func (m *ScrapeConfigsDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*ScrapeConfigsDeleteRequest) ProtoMessage()               {}
func (*ScrapeConfigsDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{12} }

func (m *ScrapeConfigsDeleteRequest) GetJobName() string {
	if m != nil {
		return m.JobName
	}
	return ""
}

type ScrapeConfigsDeleteResponse struct {
}

func (m *ScrapeConfigsDeleteResponse) Reset()                    { *m = ScrapeConfigsDeleteResponse{} }
func (m *ScrapeConfigsDeleteResponse) String() string            { return proto.CompactTextString(m) }
func (*ScrapeConfigsDeleteResponse) ProtoMessage()               {}
func (*ScrapeConfigsDeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{13} }

func init() {
	proto.RegisterType((*LabelPair)(nil), "api.LabelPair")
	proto.RegisterType((*StaticConfig)(nil), "api.StaticConfig")
	proto.RegisterType((*BasicAuth)(nil), "api.BasicAuth")
	proto.RegisterType((*TLSConfig)(nil), "api.TLSConfig")
	proto.RegisterType((*ScrapeConfig)(nil), "api.ScrapeConfig")
	proto.RegisterType((*ScrapeTargetHealth)(nil), "api.ScrapeTargetHealth")
	proto.RegisterType((*ScrapeConfigsListRequest)(nil), "api.ScrapeConfigsListRequest")
	proto.RegisterType((*ScrapeConfigsListResponse)(nil), "api.ScrapeConfigsListResponse")
	proto.RegisterType((*ScrapeConfigsGetRequest)(nil), "api.ScrapeConfigsGetRequest")
	proto.RegisterType((*ScrapeConfigsGetResponse)(nil), "api.ScrapeConfigsGetResponse")
	proto.RegisterType((*ScrapeConfigsCreateRequest)(nil), "api.ScrapeConfigsCreateRequest")
	proto.RegisterType((*ScrapeConfigsCreateResponse)(nil), "api.ScrapeConfigsCreateResponse")
	proto.RegisterType((*ScrapeConfigsDeleteRequest)(nil), "api.ScrapeConfigsDeleteRequest")
	proto.RegisterType((*ScrapeConfigsDeleteResponse)(nil), "api.ScrapeConfigsDeleteResponse")
	proto.RegisterEnum("api.ScrapeTargetHealth_Health", ScrapeTargetHealth_Health_name, ScrapeTargetHealth_Health_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ScrapeConfigs service

type ScrapeConfigsClient interface {
	// List returns all scrape configs.
	List(ctx context.Context, in *ScrapeConfigsListRequest, opts ...grpc.CallOption) (*ScrapeConfigsListResponse, error)
	// Get returns a scrape config by job name.
	// Errors: NotFound(5) if no such scrape config is present.
	Get(ctx context.Context, in *ScrapeConfigsGetRequest, opts ...grpc.CallOption) (*ScrapeConfigsGetResponse, error)
	// Create creates a new scrape config.
	// Errors: InvalidArgument(3) if some argument is not valid,
	// AlreadyExists(6) if scrape config with that job name is already present,
	// FailedPrecondition(9) if reachability check was requested and some scrape target can't be reached.
	Create(ctx context.Context, in *ScrapeConfigsCreateRequest, opts ...grpc.CallOption) (*ScrapeConfigsCreateResponse, error)
	// Delete removes existing scrape config by job name.
	// Errors: NotFound(5) if no such scrape config is present.
	Delete(ctx context.Context, in *ScrapeConfigsDeleteRequest, opts ...grpc.CallOption) (*ScrapeConfigsDeleteResponse, error)
}

type scrapeConfigsClient struct {
	cc *grpc.ClientConn
}

func NewScrapeConfigsClient(cc *grpc.ClientConn) ScrapeConfigsClient {
	return &scrapeConfigsClient{cc}
}

func (c *scrapeConfigsClient) List(ctx context.Context, in *ScrapeConfigsListRequest, opts ...grpc.CallOption) (*ScrapeConfigsListResponse, error) {
	out := new(ScrapeConfigsListResponse)
	err := grpc.Invoke(ctx, "/api.ScrapeConfigs/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scrapeConfigsClient) Get(ctx context.Context, in *ScrapeConfigsGetRequest, opts ...grpc.CallOption) (*ScrapeConfigsGetResponse, error) {
	out := new(ScrapeConfigsGetResponse)
	err := grpc.Invoke(ctx, "/api.ScrapeConfigs/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scrapeConfigsClient) Create(ctx context.Context, in *ScrapeConfigsCreateRequest, opts ...grpc.CallOption) (*ScrapeConfigsCreateResponse, error) {
	out := new(ScrapeConfigsCreateResponse)
	err := grpc.Invoke(ctx, "/api.ScrapeConfigs/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scrapeConfigsClient) Delete(ctx context.Context, in *ScrapeConfigsDeleteRequest, opts ...grpc.CallOption) (*ScrapeConfigsDeleteResponse, error) {
	out := new(ScrapeConfigsDeleteResponse)
	err := grpc.Invoke(ctx, "/api.ScrapeConfigs/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ScrapeConfigs service

type ScrapeConfigsServer interface {
	// List returns all scrape configs.
	List(context.Context, *ScrapeConfigsListRequest) (*ScrapeConfigsListResponse, error)
	// Get returns a scrape config by job name.
	// Errors: NotFound(5) if no such scrape config is present.
	Get(context.Context, *ScrapeConfigsGetRequest) (*ScrapeConfigsGetResponse, error)
	// Create creates a new scrape config.
	// Errors: InvalidArgument(3) if some argument is not valid,
	// AlreadyExists(6) if scrape config with that job name is already present,
	// FailedPrecondition(9) if reachability check was requested and some scrape target can't be reached.
	Create(context.Context, *ScrapeConfigsCreateRequest) (*ScrapeConfigsCreateResponse, error)
	// Delete removes existing scrape config by job name.
	// Errors: NotFound(5) if no such scrape config is present.
	Delete(context.Context, *ScrapeConfigsDeleteRequest) (*ScrapeConfigsDeleteResponse, error)
}

func RegisterScrapeConfigsServer(s *grpc.Server, srv ScrapeConfigsServer) {
	s.RegisterService(&_ScrapeConfigs_serviceDesc, srv)
}

func _ScrapeConfigs_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScrapeConfigsListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScrapeConfigsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ScrapeConfigs/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScrapeConfigsServer).List(ctx, req.(*ScrapeConfigsListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScrapeConfigs_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScrapeConfigsGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScrapeConfigsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ScrapeConfigs/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScrapeConfigsServer).Get(ctx, req.(*ScrapeConfigsGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScrapeConfigs_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScrapeConfigsCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScrapeConfigsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ScrapeConfigs/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScrapeConfigsServer).Create(ctx, req.(*ScrapeConfigsCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScrapeConfigs_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScrapeConfigsDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScrapeConfigsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ScrapeConfigs/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScrapeConfigsServer).Delete(ctx, req.(*ScrapeConfigsDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ScrapeConfigs_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.ScrapeConfigs",
	HandlerType: (*ScrapeConfigsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _ScrapeConfigs_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ScrapeConfigs_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _ScrapeConfigs_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ScrapeConfigs_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scrape_configs.proto",
}

func init() { proto.RegisterFile("scrape_configs.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 815 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xdd, 0x8e, 0xdb, 0x44,
	0x14, 0xc6, 0x49, 0xea, 0xc4, 0x67, 0x93, 0xb0, 0x1d, 0x02, 0xeb, 0x9a, 0xa6, 0x0d, 0x96, 0x4a,
	0xa3, 0x4a, 0x49, 0xaa, 0x05, 0x0a, 0x42, 0xe2, 0x02, 0xb6, 0x52, 0x41, 0x5d, 0x85, 0xc8, 0xd9,
	0x85, 0x4b, 0x6b, 0x6c, 0x66, 0xe3, 0xc9, 0x3a, 0xb6, 0xf1, 0x4c, 0x82, 0x56, 0x88, 0x1b, 0x78,
	0x04, 0x1e, 0x00, 0x89, 0xf7, 0xe0, 0x0d, 0xb8, 0x43, 0xe2, 0x09, 0x78, 0x10, 0xe4, 0x99, 0xb1,
	0xd7, 0xde, 0xfc, 0x5c, 0x20, 0xae, 0xe2, 0x73, 0xce, 0xcc, 0x39, 0xdf, 0xf7, 0x9d, 0x6f, 0x14,
	0xe8, 0x31, 0x3f, 0xc5, 0x09, 0x71, 0xfd, 0x38, 0xba, 0xa2, 0x0b, 0x36, 0x4e, 0xd2, 0x98, 0xc7,
	0xa8, 0x8e, 0x13, 0x6a, 0x3d, 0x5c, 0xc4, 0xf1, 0x22, 0x24, 0x13, 0x9c, 0xd0, 0x09, 0x8e, 0xa2,
	0x98, 0x63, 0x4e, 0xe3, 0x48, 0x1d, 0xb1, 0x3f, 0x02, 0xe3, 0x1c, 0x7b, 0x24, 0x9c, 0x61, 0x9a,
	0x22, 0x04, 0x8d, 0x08, 0xaf, 0x88, 0xa9, 0x0d, 0xb4, 0xa1, 0xe1, 0x88, 0x6f, 0xd4, 0x83, 0x7b,
	0x1b, 0x1c, 0xae, 0x89, 0x59, 0x13, 0x49, 0x19, 0xd8, 0x33, 0x68, 0xcf, 0xb3, 0x46, 0xfe, 0x99,
	0x18, 0x88, 0x4c, 0x68, 0x72, 0x9c, 0x2e, 0x08, 0x67, 0xa6, 0x36, 0xa8, 0x0f, 0x0d, 0x27, 0x0f,
	0xd1, 0xfb, 0xa0, 0x87, 0xd9, 0x00, 0x66, 0xd6, 0x06, 0xf5, 0xe1, 0xd1, 0x69, 0x77, 0x8c, 0x13,
	0x3a, 0x2e, 0x66, 0x3a, 0xaa, 0x6a, 0x9f, 0x81, 0xf1, 0x05, 0x66, 0xd4, 0xff, 0x7c, 0xcd, 0x03,
	0x64, 0x41, 0x6b, 0xcd, 0x48, 0x5a, 0x02, 0x53, 0xc4, 0x59, 0x2d, 0xc1, 0x8c, 0xfd, 0x10, 0xa7,
	0xdf, 0x29, 0x4c, 0x45, 0x6c, 0x7f, 0x06, 0xc6, 0xc5, 0xf9, 0x5c, 0x61, 0x7a, 0x0e, 0x3d, 0x1a,
	0x31, 0xe2, 0xaf, 0x53, 0xe2, 0xb2, 0x6b, 0x9a, 0xb8, 0x1b, 0x92, 0xd2, 0xab, 0x1b, 0xf3, 0xde,
	0x40, 0x1b, 0xb6, 0x1c, 0x94, 0xd7, 0xe6, 0xd7, 0x34, 0xf9, 0x46, 0x54, 0xec, 0x3f, 0x6b, 0xd0,
	0x9e, 0x0b, 0x21, 0x55, 0x8b, 0x07, 0xd0, 0x5a, 0xc6, 0x9e, 0x5b, 0xc2, 0xd1, 0x5c, 0xc6, 0xde,
	0x34, 0x83, 0xf1, 0x14, 0xde, 0x54, 0x9a, 0xd3, 0x88, 0x93, 0x74, 0x83, 0x43, 0x85, 0xa6, 0x2b,
	0xd3, 0x5f, 0xa9, 0x2c, 0x7a, 0x02, 0x2a, 0xe3, 0x72, 0xba, 0x22, 0xf1, 0x9a, 0x9b, 0x75, 0x71,
	0xae, 0x23, 0xb3, 0x17, 0x32, 0x89, 0xde, 0x83, 0xf6, 0x8a, 0xf0, 0x94, 0xfa, 0xcc, 0x4d, 0x30,
	0x0f, 0xcc, 0x86, 0x38, 0x74, 0xa4, 0x72, 0x33, 0xcc, 0x03, 0xf4, 0x0e, 0xe8, 0xcc, 0x0f, 0xc8,
	0x8a, 0x08, 0x0a, 0x86, 0xa3, 0x22, 0x34, 0x02, 0xf0, 0x32, 0xe9, 0x5c, 0xbc, 0xe6, 0x81, 0xa9,
	0x0f, 0xb4, 0x42, 0xe6, 0x42, 0x51, 0xc7, 0xf0, 0x0a, 0x71, 0x47, 0x00, 0x3c, 0x64, 0xca, 0x2a,
	0x66, 0xb3, 0x74, 0xbc, 0xd0, 0xce, 0x31, 0x78, 0xc8, 0x94, 0x06, 0x9f, 0x40, 0x97, 0x89, 0x55,
	0xe7, 0xe6, 0x32, 0x5b, 0x62, 0x91, 0xf7, 0xc5, 0x95, 0xb2, 0x0b, 0x9c, 0x0e, 0x2b, 0x45, 0xcc,
	0xfe, 0x5b, 0x03, 0x24, 0xe5, 0xbc, 0x10, 0x66, 0xf8, 0x92, 0xe0, 0x90, 0x07, 0x87, 0x44, 0x3d,
	0x86, 0xfa, 0x32, 0xf6, 0x94, 0x90, 0xd9, 0x67, 0xc6, 0x59, 0x3a, 0x49, 0xa9, 0xa6, 0xa2, 0xcc,
	0x05, 0x34, 0x62, 0x1c, 0x47, 0x3e, 0x51, 0x52, 0x15, 0x31, 0x7a, 0x01, 0x7a, 0x20, 0x46, 0x09,
	0x9d, 0xba, 0xa7, 0x8f, 0x24, 0xd2, 0x2d, 0x24, 0x63, 0xf9, 0xe3, 0xa8, 0xd3, 0xf6, 0x53, 0xd0,
	0x15, 0xc4, 0x23, 0x68, 0x5e, 0x4e, 0x5f, 0x4f, 0xbf, 0xfe, 0x76, 0x7a, 0xfc, 0x06, 0x6a, 0x41,
	0xe3, 0x65, 0xf6, 0xa5, 0x21, 0x1d, 0x6a, 0x97, 0xb3, 0xe3, 0x9a, 0x6d, 0x81, 0x59, 0xb6, 0x09,
	0x3b, 0xa7, 0x8c, 0x3b, 0xe4, 0xfb, 0x35, 0x61, 0xdc, 0xfe, 0x5d, 0x83, 0x07, 0x3b, 0x8a, 0x2c,
	0x89, 0x23, 0x46, 0x84, 0x98, 0x95, 0x97, 0x2a, 0x9e, 0x4b, 0x21, 0x66, 0xe9, 0x5e, 0xee, 0x0f,
	0xd5, 0x05, 0xbd, 0x86, 0xb7, 0x73, 0x1b, 0xc9, 0x97, 0xe5, 0x2a, 0x8e, 0xf2, 0x59, 0x9d, 0xec,
	0xe1, 0xe8, 0xbc, 0xc5, 0x4a, 0x39, 0x26, 0x93, 0xf6, 0x87, 0x70, 0x52, 0xc1, 0xf8, 0x8a, 0xe4,
	0xf8, 0x0f, 0x6c, 0xc7, 0xfe, 0x4d, 0xbb, 0xc3, 0x5b, 0x5c, 0x53, 0xcc, 0x5e, 0x40, 0xa7, 0xc2,
	0x4c, 0x5c, 0xde, 0x49, 0xac, 0x5d, 0x26, 0xf6, 0xff, 0xf2, 0xfa, 0x45, 0x03, 0xab, 0x82, 0xf0,
	0x2c, 0x25, 0x98, 0x93, 0x9c, 0xdb, 0x7f, 0xc5, 0x38, 0x02, 0xe4, 0x07, 0xc4, 0xbf, 0x76, 0x53,
	0x82, 0xfd, 0x00, 0x7b, 0x34, 0xa4, 0xfc, 0x46, 0xb8, 0xb4, 0xe5, 0xdc, 0x17, 0x15, 0xa7, 0x54,
	0xb0, 0xfb, 0xf0, 0xee, 0x4e, 0x10, 0x52, 0x29, 0xfb, 0xe3, 0x3b, 0x18, 0x5f, 0x92, 0x90, 0xdc,
	0x62, 0x3c, 0xa0, 0xff, 0xdd, 0xbe, 0xf9, 0x45, 0xd9, 0xf7, 0xf4, 0x8f, 0x3a, 0x74, 0x2a, 0x75,
	0x84, 0xa1, 0x91, 0xb9, 0x0f, 0xf5, 0xb7, 0x08, 0x96, 0x2d, 0x6b, 0x3d, 0xda, 0x57, 0x56, 0x80,
	0xad, 0x9f, 0xff, 0xfa, 0xe7, 0xd7, 0x5a, 0x0f, 0xa1, 0xc9, 0xe6, 0xf9, 0x44, 0x0a, 0x33, 0x52,
	0xf6, 0x45, 0x14, 0xea, 0xaf, 0x08, 0x47, 0x0f, 0xb7, 0x5b, 0xdc, 0x7a, 0xca, 0xea, 0xef, 0xa9,
	0xaa, 0xfe, 0x4f, 0x44, 0xff, 0xc7, 0xa8, 0xbf, 0xdd, 0x7f, 0xf2, 0x63, 0x2e, 0xc6, 0x4f, 0x68,
	0x09, 0xba, 0x54, 0x12, 0x3d, 0xde, 0xee, 0x57, 0x59, 0xb4, 0x35, 0xd8, 0x7f, 0x40, 0xcd, 0xec,
	0x8b, 0x99, 0x27, 0xf6, 0x0e, 0x4e, 0x9f, 0x6a, 0xcf, 0x50, 0x0a, 0xba, 0x54, 0x77, 0xd7, 0xac,
	0xca, 0xc2, 0x76, 0xcd, 0xaa, 0x2e, 0x26, 0xe7, 0xf7, 0xec, 0x30, 0x3f, 0x4f, 0x17, 0xff, 0xc8,
	0x1f, 0xfc, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x21, 0xe7, 0x35, 0x27, 0xcc, 0x07, 0x00, 0x00,
}
