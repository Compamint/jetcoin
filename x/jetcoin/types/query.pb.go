// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: jetcoin/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("jetcoin/query.proto", fileDescriptor_8a8b4b0ad0235980) }

var fileDescriptor_8a8b4b0ad0235980 = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8e, 0xb1, 0x0e, 0x82, 0x30,
	0x10, 0x40, 0x61, 0x50, 0x13, 0x46, 0x9d, 0x24, 0xa6, 0x1f, 0x60, 0x0c, 0x17, 0xf4, 0x0f, 0x1c,
	0xdc, 0x5d, 0xdd, 0xda, 0xa6, 0xa9, 0x35, 0xb6, 0x57, 0xe9, 0x61, 0xe4, 0x2f, 0xfc, 0x2c, 0x47,
	0x46, 0x47, 0x03, 0x3f, 0x62, 0x04, 0x61, 0xba, 0x4b, 0xee, 0xdd, 0xcb, 0x4b, 0x16, 0x17, 0x45,
	0x12, 0x8d, 0x83, 0x5b, 0xa9, 0x8a, 0x2a, 0xf3, 0x05, 0x12, 0xce, 0x97, 0x12, 0xad, 0xe7, 0xd6,
	0x38, 0xca, 0xfe, 0xe7, 0x61, 0xa6, 0x2b, 0x8d, 0xa8, 0xaf, 0x0a, 0xb8, 0x37, 0xc0, 0x9d, 0x43,
	0xe2, 0x64, 0xd0, 0x85, 0xfe, 0x31, 0x5d, 0x4b, 0x0c, 0x16, 0x03, 0x08, 0x1e, 0x54, 0x6f, 0x84,
	0x7b, 0x2e, 0x14, 0xf1, 0x1c, 0x3c, 0xd7, 0xc6, 0x75, 0x70, 0xcf, 0x6e, 0x67, 0xc9, 0xe4, 0xf8,
	0x23, 0xf6, 0x87, 0x57, 0xc3, 0xe2, 0xba, 0x61, 0xf1, 0xa7, 0x61, 0xf1, 0xb3, 0x65, 0x51, 0xdd,
	0xb2, 0xe8, 0xdd, 0xb2, 0xe8, 0xb4, 0xd1, 0x86, 0xce, 0xa5, 0xc8, 0x24, 0x5a, 0x18, 0x93, 0x60,
	0x28, 0x7e, 0x8c, 0x1b, 0x55, 0x5e, 0x05, 0x31, 0xed, 0xbc, 0xbb, 0x6f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xdf, 0xf7, 0x24, 0x3a, 0xd3, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

// QueryServer is the server API for Query service.
type QueryServer interface {
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "compamint.jetcoin.jetcoin.Query",
	HandlerType: (*QueryServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "jetcoin/query.proto",
}
