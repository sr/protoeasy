// Code generated by protoc-gen-go.
// source: protoeasy.proto
// DO NOT EDIT!

package protoeasy

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "go.pedge.io/google-protobuf"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type GoPluginType int32

const (
	GoPluginType_GO_PLUGIN_TYPE_NONE       GoPluginType = 0
	GoPluginType_GO_PLUGIN_TYPE_GO         GoPluginType = 1
	GoPluginType_GO_PLUGIN_TYPE_GOFAST     GoPluginType = 2
	GoPluginType_GO_PLUGIN_TYPE_GOGO       GoPluginType = 3
	GoPluginType_GO_PLUGIN_TYPE_GOGOFAST   GoPluginType = 4
	GoPluginType_GO_PLUGIN_TYPE_GOGOFASTER GoPluginType = 5
	GoPluginType_GO_PLUGIN_TYPE_GOGOSLICK  GoPluginType = 6
)

var GoPluginType_name = map[int32]string{
	0: "GO_PLUGIN_TYPE_NONE",
	1: "GO_PLUGIN_TYPE_GO",
	2: "GO_PLUGIN_TYPE_GOFAST",
	3: "GO_PLUGIN_TYPE_GOGO",
	4: "GO_PLUGIN_TYPE_GOGOFAST",
	5: "GO_PLUGIN_TYPE_GOGOFASTER",
	6: "GO_PLUGIN_TYPE_GOGOSLICK",
}
var GoPluginType_value = map[string]int32{
	"GO_PLUGIN_TYPE_NONE":       0,
	"GO_PLUGIN_TYPE_GO":         1,
	"GO_PLUGIN_TYPE_GOFAST":     2,
	"GO_PLUGIN_TYPE_GOGO":       3,
	"GO_PLUGIN_TYPE_GOGOFAST":   4,
	"GO_PLUGIN_TYPE_GOGOFASTER": 5,
	"GO_PLUGIN_TYPE_GOGOSLICK":  6,
}

func (x GoPluginType) String() string {
	return proto.EnumName(GoPluginType_name, int32(x))
}
func (GoPluginType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type CompileOptions struct {
	Grpc                 bool              `protobuf:"varint,1,opt,name=grpc" json:"grpc,omitempty"`
	GrpcGateway          bool              `protobuf:"varint,2,opt,name=grpc_gateway" json:"grpc_gateway,omitempty"`
	NoDefaultIncludes    bool              `protobuf:"varint,3,opt,name=no_default_includes" json:"no_default_includes,omitempty"`
	ExcludePattern       []string          `protobuf:"bytes,4,rep,name=exclude_pattern" json:"exclude_pattern,omitempty"`
	Cpp                  bool              `protobuf:"varint,20,opt,name=cpp" json:"cpp,omitempty"`
	CppRelOut            string            `protobuf:"bytes,21,opt,name=cpp_rel_out" json:"cpp_rel_out,omitempty"`
	Csharp               bool              `protobuf:"varint,30,opt,name=csharp" json:"csharp,omitempty"`
	CsharpRelOut         string            `protobuf:"bytes,31,opt,name=csharp_rel_out" json:"csharp_rel_out,omitempty"`
	Go                   bool              `protobuf:"varint,40,opt,name=go" json:"go,omitempty"`
	GoRelOut             string            `protobuf:"bytes,41,opt,name=go_rel_out" json:"go_rel_out,omitempty"`
	GoImportPath         string            `protobuf:"bytes,42,opt,name=go_import_path" json:"go_import_path,omitempty"`
	GoNoDefaultModifiers bool              `protobuf:"varint,43,opt,name=go_no_default_modifiers" json:"go_no_default_modifiers,omitempty"`
	GoModifiers          map[string]string `protobuf:"bytes,44,rep,name=go_modifiers" json:"go_modifiers,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	GoPluginType         GoPluginType      `protobuf:"varint,45,opt,name=go_plugin_type,enum=protoeasy.GoPluginType" json:"go_plugin_type,omitempty"`
	Objc                 bool              `protobuf:"varint,50,opt,name=objc" json:"objc,omitempty"`
	ObjcRelOut           string            `protobuf:"bytes,51,opt,name=objc_rel_out" json:"objc_rel_out,omitempty"`
	Python               bool              `protobuf:"varint,60,opt,name=python" json:"python,omitempty"`
	PythonRelOut         string            `protobuf:"bytes,61,opt,name=python_rel_out" json:"python_rel_out,omitempty"`
	Ruby                 bool              `protobuf:"varint,70,opt,name=ruby" json:"ruby,omitempty"`
	RubyRelOut           string            `protobuf:"bytes,71,opt,name=ruby_rel_out" json:"ruby_rel_out,omitempty"`
}

func (m *CompileOptions) Reset()                    { *m = CompileOptions{} }
func (m *CompileOptions) String() string            { return proto.CompactTextString(m) }
func (*CompileOptions) ProtoMessage()               {}
func (*CompileOptions) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CompileOptions) GetGoModifiers() map[string]string {
	if m != nil {
		return m.GoModifiers
	}
	return nil
}

type Command struct {
	Arg []string `protobuf:"bytes,1,rep,name=arg" json:"arg,omitempty"`
}

func (m *Command) Reset()                    { *m = Command{} }
func (m *Command) String() string            { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()               {}
func (*Command) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type CompileInfo struct {
	Command         []*Command                `protobuf:"bytes,1,rep,name=command" json:"command,omitempty"`
	InputSizeBytes  uint64                    `protobuf:"varint,2,opt,name=input_size_bytes" json:"input_size_bytes,omitempty"`
	OutputSizeBytes uint64                    `protobuf:"varint,3,opt,name=output_size_bytes" json:"output_size_bytes,omitempty"`
	Duration        *google_protobuf.Duration `protobuf:"bytes,4,opt,name=duration" json:"duration,omitempty"`
}

func (m *CompileInfo) Reset()                    { *m = CompileInfo{} }
func (m *CompileInfo) String() string            { return proto.CompactTextString(m) }
func (*CompileInfo) ProtoMessage()               {}
func (*CompileInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CompileInfo) GetCommand() []*Command {
	if m != nil {
		return m.Command
	}
	return nil
}

func (m *CompileInfo) GetDuration() *google_protobuf.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

type CompileRequest struct {
	Tar            []byte          `protobuf:"bytes,1,opt,name=tar,proto3" json:"tar,omitempty"`
	CompileOptions *CompileOptions `protobuf:"bytes,2,opt,name=compile_options" json:"compile_options,omitempty"`
}

func (m *CompileRequest) Reset()                    { *m = CompileRequest{} }
func (m *CompileRequest) String() string            { return proto.CompactTextString(m) }
func (*CompileRequest) ProtoMessage()               {}
func (*CompileRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CompileRequest) GetCompileOptions() *CompileOptions {
	if m != nil {
		return m.CompileOptions
	}
	return nil
}

type CompileResponse struct {
	Tar         []byte       `protobuf:"bytes,1,opt,name=tar,proto3" json:"tar,omitempty"`
	CompileInfo *CompileInfo `protobuf:"bytes,2,opt,name=compile_info" json:"compile_info,omitempty"`
}

func (m *CompileResponse) Reset()                    { *m = CompileResponse{} }
func (m *CompileResponse) String() string            { return proto.CompactTextString(m) }
func (*CompileResponse) ProtoMessage()               {}
func (*CompileResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CompileResponse) GetCompileInfo() *CompileInfo {
	if m != nil {
		return m.CompileInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*CompileOptions)(nil), "protoeasy.CompileOptions")
	proto.RegisterType((*Command)(nil), "protoeasy.Command")
	proto.RegisterType((*CompileInfo)(nil), "protoeasy.CompileInfo")
	proto.RegisterType((*CompileRequest)(nil), "protoeasy.CompileRequest")
	proto.RegisterType((*CompileResponse)(nil), "protoeasy.CompileResponse")
	proto.RegisterEnum("protoeasy.GoPluginType", GoPluginType_name, GoPluginType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for API service

type APIClient interface {
	Compile(ctx context.Context, in *CompileRequest, opts ...grpc.CallOption) (*CompileResponse, error)
}

type aPIClient struct {
	cc *grpc.ClientConn
}

func NewAPIClient(cc *grpc.ClientConn) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) Compile(ctx context.Context, in *CompileRequest, opts ...grpc.CallOption) (*CompileResponse, error) {
	out := new(CompileResponse)
	err := grpc.Invoke(ctx, "/protoeasy.API/Compile", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for API service

type APIServer interface {
	Compile(context.Context, *CompileRequest) (*CompileResponse, error)
}

func RegisterAPIServer(s *grpc.Server, srv APIServer) {
	s.RegisterService(&_API_serviceDesc, srv)
}

func _API_Compile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(CompileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).Compile(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _API_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protoeasy.API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Compile",
			Handler:    _API_Compile_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 652 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x53, 0x5f, 0x53, 0xd3, 0x4e,
	0x14, 0xa5, 0x7f, 0xf8, 0xd3, 0xdb, 0xfe, 0xda, 0xb0, 0xfc, 0xe9, 0xd2, 0x9f, 0x02, 0x53, 0x5f,
	0x10, 0xb0, 0xcc, 0x84, 0x17, 0xc7, 0xd1, 0x19, 0x11, 0x4b, 0xa6, 0x23, 0xd2, 0x0a, 0xf8, 0xe0,
	0xd3, 0x4e, 0x9a, 0x6e, 0x43, 0x34, 0xcd, 0xae, 0xc9, 0x46, 0x8d, 0x9f, 0xc3, 0x8f, 0xe0, 0xc7,
	0xf1, 0x43, 0xb9, 0xbb, 0x69, 0x22, 0x85, 0xf2, 0xd6, 0x7b, 0xce, 0xc9, 0xd9, 0x7b, 0xcf, 0xbd,
	0x85, 0x06, 0x0f, 0x99, 0x60, 0xd4, 0x8e, 0x92, 0x8e, 0xfe, 0x85, 0x2a, 0x39, 0xd0, 0xda, 0x76,
	0x19, 0x73, 0x7d, 0x7a, 0xa4, 0x91, 0x61, 0x3c, 0x3e, 0x1a, 0xc5, 0xa1, 0x2d, 0x3c, 0x16, 0xa4,
	0xd2, 0xf6, 0xef, 0x32, 0xd4, 0x4f, 0xd9, 0x84, 0x7b, 0x3e, 0xed, 0x73, 0x85, 0x47, 0xa8, 0x06,
	0x65, 0x37, 0xe4, 0x0e, 0x2e, 0xec, 0x16, 0xf6, 0x56, 0xd0, 0x3a, 0xd4, 0x54, 0x45, 0x5c, 0x5b,
	0xd0, 0xef, 0x76, 0x82, 0x8b, 0x1a, 0xfd, 0x1f, 0xd6, 0x02, 0x46, 0x46, 0x74, 0x6c, 0xc7, 0xbe,
	0x20, 0x5e, 0xe0, 0xf8, 0xf1, 0x88, 0x46, 0xb8, 0xa4, 0xc9, 0x26, 0x34, 0xe8, 0x0f, 0x8d, 0x10,
	0x6e, 0x0b, 0x41, 0xc3, 0x00, 0x97, 0x77, 0x4b, 0x7b, 0x15, 0x54, 0x85, 0x92, 0xc3, 0x39, 0x5e,
	0xd7, 0xaa, 0x35, 0xa8, 0xca, 0x82, 0x84, 0xd4, 0x27, 0x2c, 0x16, 0x78, 0x43, 0x82, 0x15, 0x54,
	0x87, 0x25, 0x27, 0xba, 0xb1, 0x43, 0x8e, 0xb7, 0xb5, 0x68, 0x13, 0xea, 0x69, 0x9d, 0xeb, 0x76,
	0xb4, 0x0e, 0xa0, 0xe8, 0x32, 0xbc, 0xa7, 0x35, 0xb2, 0x70, 0x59, 0xce, 0x3f, 0xd5, 0xbc, 0xfc,
	0x4e, 0x62, 0xde, 0x84, 0xb3, 0x50, 0xa8, 0x26, 0x6e, 0xf0, 0xbe, 0xc6, 0x77, 0xa0, 0x29, 0xf1,
	0x5b, 0xad, 0x4f, 0xd8, 0xc8, 0x1b, 0x7b, 0x34, 0x8c, 0xf0, 0x81, 0x36, 0x7b, 0x2d, 0xc7, 0x65,
	0xb7, 0xd0, 0x43, 0xd9, 0x78, 0xd5, 0xdc, 0xef, 0xfc, 0x8b, 0x78, 0x36, 0xad, 0x8e, 0xc5, 0xde,
	0x67, 0xe2, 0x6e, 0x20, 0xc2, 0x04, 0x1d, 0xe9, 0xa7, 0xb9, 0x1f, 0xbb, 0x5e, 0x40, 0x44, 0xc2,
	0x29, 0x7e, 0x26, 0x9d, 0xeb, 0x66, 0xf3, 0x96, 0x87, 0xc5, 0x06, 0x9a, 0xbf, 0x96, 0xb4, 0xca,
	0x9b, 0x0d, 0x3f, 0x3b, 0xd8, 0xcc, 0xf2, 0x56, 0x55, 0x3e, 0xcf, 0x71, 0x96, 0x0b, 0x4f, 0xc4,
	0x0d, 0x0b, 0xf0, 0xcb, 0x2c, 0x97, 0xb4, 0xce, 0x75, 0xaf, 0xb4, 0x4e, 0x7a, 0x85, 0xf1, 0x30,
	0xc1, 0x67, 0x99, 0x97, 0xaa, 0x72, 0x8d, 0xa5, 0x34, 0x2d, 0x13, 0x8c, 0x7b, 0x4d, 0xcb, 0xcd,
	0x7c, 0xa1, 0x89, 0x5e, 0x79, 0x05, 0xfd, 0x07, 0x8b, 0xdf, 0x6c, 0x3f, 0xa6, 0x7a, 0xd7, 0x95,
	0x17, 0xc5, 0xe7, 0x85, 0xf6, 0x26, 0x2c, 0xcb, 0xb9, 0x27, 0x76, 0x30, 0x52, 0x52, 0x3b, 0x74,
	0xa5, 0x54, 0x6e, 0xb4, 0xfd, 0xab, 0x00, 0xd5, 0x69, 0x20, 0xbd, 0x60, 0xcc, 0xd0, 0x13, 0x58,
	0x76, 0x52, 0x9d, 0x16, 0x54, 0x4d, 0x34, 0x9b, 0x9c, 0x76, 0xc0, 0x60, 0x78, 0x01, 0x8f, 0x05,
	0x89, 0xbc, 0x9f, 0x94, 0x0c, 0x13, 0x21, 0x2f, 0x47, 0x3d, 0x55, 0x46, 0x5b, 0xb0, 0x2a, 0xfb,
	0xbc, 0x43, 0x95, 0x34, 0x75, 0x00, 0x2b, 0xd9, 0xe9, 0xca, 0x6b, 0x2a, 0x48, 0xeb, 0xad, 0x4e,
	0x7a, 0xdb, 0x9d, 0xec, 0xb6, 0x3b, 0x6f, 0xa7, 0x82, 0xf6, 0x87, 0xfc, 0xa8, 0x2f, 0xe9, 0xd7,
	0x98, 0x46, 0x42, 0x75, 0x2d, 0xec, 0x50, 0x0f, 0x58, 0x43, 0x26, 0x34, 0x9c, 0x94, 0x26, 0x2c,
	0x5d, 0xa3, 0x7e, 0x5f, 0x59, 0x3e, 0xb4, 0xe7, 0xf6, 0x39, 0x34, 0x72, 0xcb, 0x88, 0x4b, 0x84,
	0xce, 0x7a, 0x1e, 0x42, 0x2d, 0xf3, 0xf4, 0x64, 0x12, 0x53, 0xc3, 0xcd, 0xfb, 0x86, 0x2a, 0xa7,
	0xfd, 0x3f, 0x05, 0xa8, 0xcd, 0x1c, 0x41, 0x13, 0xd6, 0xac, 0x3e, 0x19, 0x9c, 0x7f, 0xb4, 0x7a,
	0x17, 0xe4, 0xfa, 0xd3, 0xa0, 0x4b, 0x2e, 0xfa, 0x17, 0x5d, 0x63, 0x01, 0x6d, 0xc0, 0xea, 0x1d,
	0xc2, 0xea, 0x1b, 0x05, 0x99, 0xd4, 0xc6, 0x3d, 0xf8, 0xec, 0xe4, 0xea, 0xda, 0x28, 0xce, 0xb1,
	0xb2, 0xfa, 0xf2, 0x9b, 0x92, 0xfc, 0xd3, 0x36, 0xe7, 0x10, 0xfa, 0xab, 0x32, 0x7a, 0x0c, 0x5b,
	0x0f, 0x90, 0xdd, 0x4b, 0x63, 0x11, 0x3d, 0x02, 0x3c, 0x87, 0xbe, 0x3a, 0xef, 0x9d, 0xbe, 0x33,
	0x96, 0xcc, 0x1e, 0x94, 0x4e, 0x06, 0x3d, 0xf4, 0x46, 0x5f, 0x89, 0x1a, 0x12, 0xcd, 0x49, 0x72,
	0xba, 0x8a, 0x56, 0x6b, 0x1e, 0x95, 0x46, 0xda, 0x5e, 0x18, 0x2e, 0x69, 0xf2, 0xf8, 0x6f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x9c, 0xe2, 0xa1, 0xa1, 0xd5, 0x04, 0x00, 0x00,
}
