// Code generated by protoc-gen-go.
// source: google/iam/v1/iam_policy.proto
// DO NOT EDIT!

/*
Package google_iam_v1 is a generated protocol buffer package.

It is generated from these files:
	google/iam/v1/iam_policy.proto
	google/iam/v1/policy.proto

It has these top-level messages:
	SetIamPolicyRequest
	GetIamPolicyRequest
	TestIamPermissionsRequest
	TestIamPermissionsResponse
	Policy
	Binding
*/
package google_iam_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Request message for `SetIamPolicy` method.
type SetIamPolicyRequest struct {
	// REQUIRED: The resource for which policy is being specified.
	// Resource is usually specified as a path, such as,
	// projects/{project}/zones/{zone}/disks/{disk}.
	Resource string `protobuf:"bytes,1,opt,name=resource" json:"resource,omitempty"`
	// REQUIRED: The complete policy to be applied to the 'resource'. The size of
	// the policy is limited to a few 10s of KB. An empty policy is in general a
	// valid policy but certain services (like Projects) might reject them.
	Policy *Policy `protobuf:"bytes,2,opt,name=policy" json:"policy,omitempty"`
}

func (m *SetIamPolicyRequest) Reset()                    { *m = SetIamPolicyRequest{} }
func (m *SetIamPolicyRequest) String() string            { return proto.CompactTextString(m) }
func (*SetIamPolicyRequest) ProtoMessage()               {}
func (*SetIamPolicyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SetIamPolicyRequest) GetPolicy() *Policy {
	if m != nil {
		return m.Policy
	}
	return nil
}

// Request message for `GetIamPolicy` method.
type GetIamPolicyRequest struct {
	// REQUIRED: The resource for which policy is being requested. Resource
	// is usually specified as a path, such as, projects/{project}.
	Resource string `protobuf:"bytes,1,opt,name=resource" json:"resource,omitempty"`
}

func (m *GetIamPolicyRequest) Reset()                    { *m = GetIamPolicyRequest{} }
func (m *GetIamPolicyRequest) String() string            { return proto.CompactTextString(m) }
func (*GetIamPolicyRequest) ProtoMessage()               {}
func (*GetIamPolicyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// Request message for `TestIamPermissions` method.
type TestIamPermissionsRequest struct {
	// REQUIRED: The resource for which policy detail is being requested.
	// Resource is usually specified as a path, such as, projects/{project}.
	Resource string `protobuf:"bytes,1,opt,name=resource" json:"resource,omitempty"`
	// The set of permissions to check for the 'resource'. Permissions with
	// wildcards (such as '*' or 'storage.*') are not allowed.
	Permissions []string `protobuf:"bytes,2,rep,name=permissions" json:"permissions,omitempty"`
}

func (m *TestIamPermissionsRequest) Reset()                    { *m = TestIamPermissionsRequest{} }
func (m *TestIamPermissionsRequest) String() string            { return proto.CompactTextString(m) }
func (*TestIamPermissionsRequest) ProtoMessage()               {}
func (*TestIamPermissionsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// Response message for `TestIamPermissions` method.
type TestIamPermissionsResponse struct {
	// A subset of `TestPermissionsRequest.permissions` that the caller is
	// allowed.
	Permissions []string `protobuf:"bytes,1,rep,name=permissions" json:"permissions,omitempty"`
}

func (m *TestIamPermissionsResponse) Reset()                    { *m = TestIamPermissionsResponse{} }
func (m *TestIamPermissionsResponse) String() string            { return proto.CompactTextString(m) }
func (*TestIamPermissionsResponse) ProtoMessage()               {}
func (*TestIamPermissionsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*SetIamPolicyRequest)(nil), "google.iam.v1.SetIamPolicyRequest")
	proto.RegisterType((*GetIamPolicyRequest)(nil), "google.iam.v1.GetIamPolicyRequest")
	proto.RegisterType((*TestIamPermissionsRequest)(nil), "google.iam.v1.TestIamPermissionsRequest")
	proto.RegisterType((*TestIamPermissionsResponse)(nil), "google.iam.v1.TestIamPermissionsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for IAMPolicy service

type IAMPolicyClient interface {
	// Sets the access control policy on the specified resource. Replaces any
	// existing policy.
	SetIamPolicy(ctx context.Context, in *SetIamPolicyRequest, opts ...grpc.CallOption) (*Policy, error)
	// Gets the access control policy for a resource. Is empty if the
	// policy or the resource does not exist.
	GetIamPolicy(ctx context.Context, in *GetIamPolicyRequest, opts ...grpc.CallOption) (*Policy, error)
	// Returns permissions that a caller has on the specified resource.
	TestIamPermissions(ctx context.Context, in *TestIamPermissionsRequest, opts ...grpc.CallOption) (*TestIamPermissionsResponse, error)
}

type iAMPolicyClient struct {
	cc *grpc.ClientConn
}

func NewIAMPolicyClient(cc *grpc.ClientConn) IAMPolicyClient {
	return &iAMPolicyClient{cc}
}

func (c *iAMPolicyClient) SetIamPolicy(ctx context.Context, in *SetIamPolicyRequest, opts ...grpc.CallOption) (*Policy, error) {
	out := new(Policy)
	err := grpc.Invoke(ctx, "/google.iam.v1.IAMPolicy/SetIamPolicy", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMPolicyClient) GetIamPolicy(ctx context.Context, in *GetIamPolicyRequest, opts ...grpc.CallOption) (*Policy, error) {
	out := new(Policy)
	err := grpc.Invoke(ctx, "/google.iam.v1.IAMPolicy/GetIamPolicy", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMPolicyClient) TestIamPermissions(ctx context.Context, in *TestIamPermissionsRequest, opts ...grpc.CallOption) (*TestIamPermissionsResponse, error) {
	out := new(TestIamPermissionsResponse)
	err := grpc.Invoke(ctx, "/google.iam.v1.IAMPolicy/TestIamPermissions", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for IAMPolicy service

type IAMPolicyServer interface {
	// Sets the access control policy on the specified resource. Replaces any
	// existing policy.
	SetIamPolicy(context.Context, *SetIamPolicyRequest) (*Policy, error)
	// Gets the access control policy for a resource. Is empty if the
	// policy or the resource does not exist.
	GetIamPolicy(context.Context, *GetIamPolicyRequest) (*Policy, error)
	// Returns permissions that a caller has on the specified resource.
	TestIamPermissions(context.Context, *TestIamPermissionsRequest) (*TestIamPermissionsResponse, error)
}

func RegisterIAMPolicyServer(s *grpc.Server, srv IAMPolicyServer) {
	s.RegisterService(&_IAMPolicy_serviceDesc, srv)
}

func _IAMPolicy_SetIamPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(SetIamPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(IAMPolicyServer).SetIamPolicy(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _IAMPolicy_GetIamPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(GetIamPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(IAMPolicyServer).GetIamPolicy(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _IAMPolicy_TestIamPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(TestIamPermissionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(IAMPolicyServer).TestIamPermissions(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _IAMPolicy_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.iam.v1.IAMPolicy",
	HandlerType: (*IAMPolicyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetIamPolicy",
			Handler:    _IAMPolicy_SetIamPolicy_Handler,
		},
		{
			MethodName: "GetIamPolicy",
			Handler:    _IAMPolicy_GetIamPolicy_Handler,
		},
		{
			MethodName: "TestIamPermissions",
			Handler:    _IAMPolicy_TestIamPermissions_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0xcf, 0x4c, 0xcc, 0xd5, 0x2f, 0x33, 0x04, 0x51, 0xf1, 0x05, 0xf9, 0x39, 0x99,
	0xc9, 0x95, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xbc, 0x10, 0x79, 0x3d, 0xa0, 0x84, 0x5e,
	0x99, 0xa1, 0x94, 0x14, 0xaa, 0x72, 0x64, 0xa5, 0x4a, 0x7e, 0x5c, 0xc2, 0xc1, 0xa9, 0x25, 0x9e,
	0x89, 0xb9, 0x01, 0x60, 0xd1, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x01, 0x2e, 0x8e,
	0xa2, 0xd4, 0xe2, 0xfc, 0xd2, 0xa2, 0xe4, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x4e, 0x21, 0x55,
	0x2e, 0x36, 0x88, 0x46, 0x09, 0x26, 0x20, 0x9f, 0xdb, 0x48, 0x54, 0x0f, 0xc5, 0x12, 0x3d, 0x88,
	0x7e, 0x25, 0x75, 0x2e, 0x61, 0x77, 0x62, 0xcc, 0x53, 0x72, 0xe2, 0x92, 0x0c, 0x01, 0xca, 0x80,
	0x54, 0xa6, 0x16, 0xe5, 0x66, 0x16, 0x17, 0x67, 0xe6, 0xe7, 0x15, 0xe3, 0xb6, 0x5e, 0x98, 0x8b,
	0xbb, 0x00, 0xa1, 0x0e, 0xe8, 0x06, 0x66, 0xa0, 0x19, 0x86, 0x5c, 0x52, 0xd8, 0xcc, 0x28, 0x2e,
	0x00, 0x52, 0xa9, 0xe8, 0x5a, 0x18, 0x41, 0x5a, 0x8c, 0x7a, 0x98, 0xb8, 0x38, 0x3d, 0x1d, 0x7d,
	0x21, 0xae, 0x13, 0xf2, 0xe4, 0xe2, 0x41, 0xf6, 0xbd, 0x90, 0x12, 0x9a, 0xa7, 0xb0, 0x04, 0x8d,
	0x14, 0x76, 0x8f, 0x83, 0x8c, 0x72, 0xc7, 0x67, 0x94, 0x3b, 0xf1, 0x46, 0x65, 0x72, 0x09, 0x61,
	0x7a, 0x4b, 0x48, 0x03, 0x4d, 0x31, 0xce, 0xd0, 0x93, 0xd2, 0x24, 0x42, 0x25, 0x24, 0x8c, 0x9c,
	0x54, 0xb9, 0x04, 0x93, 0xf3, 0x73, 0x51, 0xd5, 0x3b, 0xf1, 0xc1, 0x1d, 0x1a, 0x00, 0x4a, 0x23,
	0x01, 0x8c, 0x49, 0x6c, 0xe0, 0xc4, 0x62, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x7f, 0xe3, 0xb4,
	0xab, 0x79, 0x02, 0x00, 0x00,
}
