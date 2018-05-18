// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

// The request message containing the process name.
type ProcessesRequest struct {
	Process              string   `protobuf:"bytes,1,opt,name=process" json:"process,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProcessesRequest) Reset()         { *m = ProcessesRequest{} }
func (m *ProcessesRequest) String() string { return proto.CompactTextString(m) }
func (*ProcessesRequest) ProtoMessage()    {}
func (*ProcessesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_d4851022f46f517e, []int{0}
}
func (m *ProcessesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcessesRequest.Unmarshal(m, b)
}
func (m *ProcessesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcessesRequest.Marshal(b, m, deterministic)
}
func (dst *ProcessesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessesRequest.Merge(dst, src)
}
func (m *ProcessesRequest) XXX_Size() int {
	return xxx_messageInfo_ProcessesRequest.Size(m)
}
func (m *ProcessesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessesRequest proto.InternalMessageInfo

func (m *ProcessesRequest) GetProcess() string {
	if m != nil {
		return m.Process
	}
	return ""
}

// The response message containing the requested logs.
type ProcessesReply struct {
	Content              []byte   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProcessesReply) Reset()         { *m = ProcessesReply{} }
func (m *ProcessesReply) String() string { return proto.CompactTextString(m) }
func (*ProcessesReply) ProtoMessage()    {}
func (*ProcessesReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_d4851022f46f517e, []int{1}
}
func (m *ProcessesReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcessesReply.Unmarshal(m, b)
}
func (m *ProcessesReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcessesReply.Marshal(b, m, deterministic)
}
func (dst *ProcessesReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessesReply.Merge(dst, src)
}
func (m *ProcessesReply) XXX_Size() int {
	return xxx_messageInfo_ProcessesReply.Size(m)
}
func (m *ProcessesReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessesReply.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessesReply proto.InternalMessageInfo

func (m *ProcessesReply) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

// The request message containing the process name.
type LogsRequest struct {
	Process              string   `protobuf:"bytes,1,opt,name=process" json:"process,omitempty"`
	Container            bool     `protobuf:"varint,2,opt,name=container" json:"container,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogsRequest) Reset()         { *m = LogsRequest{} }
func (m *LogsRequest) String() string { return proto.CompactTextString(m) }
func (*LogsRequest) ProtoMessage()    {}
func (*LogsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_d4851022f46f517e, []int{2}
}
func (m *LogsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogsRequest.Unmarshal(m, b)
}
func (m *LogsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogsRequest.Marshal(b, m, deterministic)
}
func (dst *LogsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogsRequest.Merge(dst, src)
}
func (m *LogsRequest) XXX_Size() int {
	return xxx_messageInfo_LogsRequest.Size(m)
}
func (m *LogsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LogsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LogsRequest proto.InternalMessageInfo

func (m *LogsRequest) GetProcess() string {
	if m != nil {
		return m.Process
	}
	return ""
}

func (m *LogsRequest) GetContainer() bool {
	if m != nil {
		return m.Container
	}
	return false
}

// The response message containing the requested logs.
type Data struct {
	Bytes                []byte   `protobuf:"bytes,1,opt,name=bytes,proto3" json:"bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Data) Reset()         { *m = Data{} }
func (m *Data) String() string { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()    {}
func (*Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_d4851022f46f517e, []int{3}
}
func (m *Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Data.Unmarshal(m, b)
}
func (m *Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Data.Marshal(b, m, deterministic)
}
func (dst *Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data.Merge(dst, src)
}
func (m *Data) XXX_Size() int {
	return xxx_messageInfo_Data.Size(m)
}
func (m *Data) XXX_DiscardUnknown() {
	xxx_messageInfo_Data.DiscardUnknown(m)
}

var xxx_messageInfo_Data proto.InternalMessageInfo

func (m *Data) GetBytes() []byte {
	if m != nil {
		return m.Bytes
	}
	return nil
}

func init() {
	proto.RegisterType((*ProcessesRequest)(nil), "proto.ProcessesRequest")
	proto.RegisterType((*ProcessesReply)(nil), "proto.ProcessesReply")
	proto.RegisterType((*LogsRequest)(nil), "proto.LogsRequest")
	proto.RegisterType((*Data)(nil), "proto.Data")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OSDClient is the client API for OSD service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OSDClient interface {
	Kubeconfig(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Data, error)
	Processes(ctx context.Context, in *ProcessesRequest, opts ...grpc.CallOption) (*ProcessesReply, error)
	Logs(ctx context.Context, in *LogsRequest, opts ...grpc.CallOption) (OSD_LogsClient, error)
	Dmesg(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Data, error)
	Version(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Data, error)
}

type oSDClient struct {
	cc *grpc.ClientConn
}

func NewOSDClient(cc *grpc.ClientConn) OSDClient {
	return &oSDClient{cc}
}

func (c *oSDClient) Kubeconfig(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Data, error) {
	out := new(Data)
	err := c.cc.Invoke(ctx, "/proto.OSD/Kubeconfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oSDClient) Processes(ctx context.Context, in *ProcessesRequest, opts ...grpc.CallOption) (*ProcessesReply, error) {
	out := new(ProcessesReply)
	err := c.cc.Invoke(ctx, "/proto.OSD/Processes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oSDClient) Logs(ctx context.Context, in *LogsRequest, opts ...grpc.CallOption) (OSD_LogsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_OSD_serviceDesc.Streams[0], "/proto.OSD/Logs", opts...)
	if err != nil {
		return nil, err
	}
	x := &oSDLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OSD_LogsClient interface {
	Recv() (*Data, error)
	grpc.ClientStream
}

type oSDLogsClient struct {
	grpc.ClientStream
}

func (x *oSDLogsClient) Recv() (*Data, error) {
	m := new(Data)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *oSDClient) Dmesg(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Data, error) {
	out := new(Data)
	err := c.cc.Invoke(ctx, "/proto.OSD/Dmesg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oSDClient) Version(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Data, error) {
	out := new(Data)
	err := c.cc.Invoke(ctx, "/proto.OSD/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for OSD service

type OSDServer interface {
	Kubeconfig(context.Context, *empty.Empty) (*Data, error)
	Processes(context.Context, *ProcessesRequest) (*ProcessesReply, error)
	Logs(*LogsRequest, OSD_LogsServer) error
	Dmesg(context.Context, *empty.Empty) (*Data, error)
	Version(context.Context, *empty.Empty) (*Data, error)
}

func RegisterOSDServer(s *grpc.Server, srv OSDServer) {
	s.RegisterService(&_OSD_serviceDesc, srv)
}

func _OSD_Kubeconfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OSDServer).Kubeconfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.OSD/Kubeconfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OSDServer).Kubeconfig(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _OSD_Processes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OSDServer).Processes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.OSD/Processes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OSDServer).Processes(ctx, req.(*ProcessesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OSD_Logs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(LogsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OSDServer).Logs(m, &oSDLogsServer{stream})
}

type OSD_LogsServer interface {
	Send(*Data) error
	grpc.ServerStream
}

type oSDLogsServer struct {
	grpc.ServerStream
}

func (x *oSDLogsServer) Send(m *Data) error {
	return x.ServerStream.SendMsg(m)
}

func _OSD_Dmesg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OSDServer).Dmesg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.OSD/Dmesg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OSDServer).Dmesg(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _OSD_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OSDServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.OSD/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OSDServer).Version(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _OSD_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.OSD",
	HandlerType: (*OSDServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Kubeconfig",
			Handler:    _OSD_Kubeconfig_Handler,
		},
		{
			MethodName: "Processes",
			Handler:    _OSD_Processes_Handler,
		},
		{
			MethodName: "Dmesg",
			Handler:    _OSD_Dmesg_Handler,
		},
		{
			MethodName: "Version",
			Handler:    _OSD_Version_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Logs",
			Handler:       _OSD_Logs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_api_d4851022f46f517e) }

var fileDescriptor_api_d4851022f46f517e = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x9b, 0xda, 0x58, 0x33, 0x15, 0x91, 0xc1, 0x3f, 0x21, 0xf6, 0x50, 0xf6, 0x54, 0x54,
	0xb6, 0xc5, 0x9e, 0xbd, 0xa5, 0x27, 0x05, 0x25, 0x82, 0xf7, 0x24, 0x4c, 0x43, 0x20, 0xcd, 0xae,
	0xd9, 0xcd, 0x21, 0xdf, 0xc0, 0x8f, 0x2d, 0x9b, 0x4d, 0xb4, 0xd4, 0x43, 0x7b, 0x5a, 0xde, 0x9b,
	0x37, 0x33, 0xfb, 0x1b, 0xf0, 0x62, 0x99, 0x73, 0x59, 0x09, 0x2d, 0xd0, 0x6d, 0x9f, 0xe0, 0x2e,
	0x13, 0x22, 0x2b, 0x68, 0xd1, 0xaa, 0xa4, 0xde, 0x2c, 0x68, 0x2b, 0x75, 0x63, 0x33, 0xec, 0x11,
	0x2e, 0xdf, 0x2b, 0x91, 0x92, 0x52, 0xa4, 0x22, 0xfa, 0xaa, 0x49, 0x69, 0xf4, 0x61, 0x2c, 0xad,
	0xe7, 0x3b, 0x33, 0x67, 0xee, 0x45, 0xbd, 0x64, 0xf7, 0x70, 0xb1, 0x93, 0x96, 0x45, 0x63, 0xb2,
	0xa9, 0x28, 0x35, 0x95, 0xba, 0xcd, 0x9e, 0x47, 0xbd, 0x64, 0x6b, 0x98, 0xbc, 0x8a, 0xec, 0xf0,
	0x50, 0x9c, 0x82, 0x67, 0x7a, 0xe2, 0xbc, 0xa4, 0xca, 0x1f, 0xce, 0x9c, 0xf9, 0x59, 0xf4, 0x67,
	0xb0, 0x29, 0x8c, 0xc2, 0x58, 0xc7, 0x78, 0x05, 0x6e, 0xd2, 0x68, 0x52, 0xdd, 0x1a, 0x2b, 0x9e,
	0xbe, 0x87, 0x70, 0xf2, 0xf6, 0x11, 0xe2, 0x0a, 0xe0, 0xa5, 0x4e, 0x28, 0x15, 0xe5, 0x26, 0xcf,
	0xf0, 0x86, 0x5b, 0x64, 0xde, 0x23, 0xf3, 0xb5, 0x41, 0x0e, 0x26, 0xd6, 0xe0, 0x66, 0x20, 0x1b,
	0xe0, 0x33, 0x78, 0xbf, 0x34, 0x78, 0xdb, 0xd5, 0xf6, 0xaf, 0x11, 0x5c, 0xff, 0x2f, 0xc8, 0xa2,
	0x61, 0x03, 0x7c, 0x80, 0x91, 0x01, 0x44, 0xec, 0x02, 0x3b, 0xb4, 0x7b, 0x9b, 0x96, 0x0e, 0x72,
	0x70, 0xc3, 0x2d, 0xa9, 0xa3, 0xff, 0xb6, 0x84, 0xf1, 0x27, 0x55, 0x2a, 0x17, 0xe5, 0x91, 0x1d,
	0xc9, 0x69, 0xab, 0x56, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xdb, 0xe3, 0x46, 0x52, 0x01, 0x02,
	0x00, 0x00,
}
