// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calculator/calculatorpb/calculator.proto

package calculatorpb

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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Sum struct {
	FirstNumber          int32    `protobuf:"varint,1,opt,name=first_number,json=firstNumber,proto3" json:"first_number,omitempty"`
	SecondNumber         int32    `protobuf:"varint,2,opt,name=second_number,json=secondNumber,proto3" json:"second_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Sum) Reset()         { *m = Sum{} }
func (m *Sum) String() string { return proto.CompactTextString(m) }
func (*Sum) ProtoMessage()    {}
func (*Sum) Descriptor() ([]byte, []int) {
	return fileDescriptor_calculator_be6d262a2f835227, []int{0}
}
func (m *Sum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Sum.Unmarshal(m, b)
}
func (m *Sum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Sum.Marshal(b, m, deterministic)
}
func (dst *Sum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sum.Merge(dst, src)
}
func (m *Sum) XXX_Size() int {
	return xxx_messageInfo_Sum.Size(m)
}
func (m *Sum) XXX_DiscardUnknown() {
	xxx_messageInfo_Sum.DiscardUnknown(m)
}

var xxx_messageInfo_Sum proto.InternalMessageInfo

func (m *Sum) GetFirstNumber() int32 {
	if m != nil {
		return m.FirstNumber
	}
	return 0
}

func (m *Sum) GetSecondNumber() int32 {
	if m != nil {
		return m.SecondNumber
	}
	return 0
}

type SumRequest struct {
	Sum                  *Sum     `protobuf:"bytes,1,opt,name=sum,proto3" json:"sum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumRequest) Reset()         { *m = SumRequest{} }
func (m *SumRequest) String() string { return proto.CompactTextString(m) }
func (*SumRequest) ProtoMessage()    {}
func (*SumRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_calculator_be6d262a2f835227, []int{1}
}
func (m *SumRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumRequest.Unmarshal(m, b)
}
func (m *SumRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumRequest.Marshal(b, m, deterministic)
}
func (dst *SumRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumRequest.Merge(dst, src)
}
func (m *SumRequest) XXX_Size() int {
	return xxx_messageInfo_SumRequest.Size(m)
}
func (m *SumRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SumRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SumRequest proto.InternalMessageInfo

func (m *SumRequest) GetSum() *Sum {
	if m != nil {
		return m.Sum
	}
	return nil
}

type SumResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumResponse) Reset()         { *m = SumResponse{} }
func (m *SumResponse) String() string { return proto.CompactTextString(m) }
func (*SumResponse) ProtoMessage()    {}
func (*SumResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_calculator_be6d262a2f835227, []int{2}
}
func (m *SumResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumResponse.Unmarshal(m, b)
}
func (m *SumResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumResponse.Marshal(b, m, deterministic)
}
func (dst *SumResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumResponse.Merge(dst, src)
}
func (m *SumResponse) XXX_Size() int {
	return xxx_messageInfo_SumResponse.Size(m)
}
func (m *SumResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SumResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SumResponse proto.InternalMessageInfo

func (m *SumResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type PrimeNumberRequest struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrimeNumberRequest) Reset()         { *m = PrimeNumberRequest{} }
func (m *PrimeNumberRequest) String() string { return proto.CompactTextString(m) }
func (*PrimeNumberRequest) ProtoMessage()    {}
func (*PrimeNumberRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_calculator_be6d262a2f835227, []int{3}
}
func (m *PrimeNumberRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimeNumberRequest.Unmarshal(m, b)
}
func (m *PrimeNumberRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimeNumberRequest.Marshal(b, m, deterministic)
}
func (dst *PrimeNumberRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimeNumberRequest.Merge(dst, src)
}
func (m *PrimeNumberRequest) XXX_Size() int {
	return xxx_messageInfo_PrimeNumberRequest.Size(m)
}
func (m *PrimeNumberRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimeNumberRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PrimeNumberRequest proto.InternalMessageInfo

func (m *PrimeNumberRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type PrimeNumberResponse struct {
	Prime                int32    `protobuf:"varint,1,opt,name=prime,proto3" json:"prime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrimeNumberResponse) Reset()         { *m = PrimeNumberResponse{} }
func (m *PrimeNumberResponse) String() string { return proto.CompactTextString(m) }
func (*PrimeNumberResponse) ProtoMessage()    {}
func (*PrimeNumberResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_calculator_be6d262a2f835227, []int{4}
}
func (m *PrimeNumberResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimeNumberResponse.Unmarshal(m, b)
}
func (m *PrimeNumberResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimeNumberResponse.Marshal(b, m, deterministic)
}
func (dst *PrimeNumberResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimeNumberResponse.Merge(dst, src)
}
func (m *PrimeNumberResponse) XXX_Size() int {
	return xxx_messageInfo_PrimeNumberResponse.Size(m)
}
func (m *PrimeNumberResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimeNumberResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PrimeNumberResponse proto.InternalMessageInfo

func (m *PrimeNumberResponse) GetPrime() int32 {
	if m != nil {
		return m.Prime
	}
	return 0
}

type ComputeAverageRequest struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ComputeAverageRequest) Reset()         { *m = ComputeAverageRequest{} }
func (m *ComputeAverageRequest) String() string { return proto.CompactTextString(m) }
func (*ComputeAverageRequest) ProtoMessage()    {}
func (*ComputeAverageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_calculator_be6d262a2f835227, []int{5}
}
func (m *ComputeAverageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComputeAverageRequest.Unmarshal(m, b)
}
func (m *ComputeAverageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComputeAverageRequest.Marshal(b, m, deterministic)
}
func (dst *ComputeAverageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComputeAverageRequest.Merge(dst, src)
}
func (m *ComputeAverageRequest) XXX_Size() int {
	return xxx_messageInfo_ComputeAverageRequest.Size(m)
}
func (m *ComputeAverageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ComputeAverageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ComputeAverageRequest proto.InternalMessageInfo

func (m *ComputeAverageRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type ComputeAverageResponse struct {
	Average              float32  `protobuf:"fixed32,1,opt,name=average,proto3" json:"average,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ComputeAverageResponse) Reset()         { *m = ComputeAverageResponse{} }
func (m *ComputeAverageResponse) String() string { return proto.CompactTextString(m) }
func (*ComputeAverageResponse) ProtoMessage()    {}
func (*ComputeAverageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_calculator_be6d262a2f835227, []int{6}
}
func (m *ComputeAverageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComputeAverageResponse.Unmarshal(m, b)
}
func (m *ComputeAverageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComputeAverageResponse.Marshal(b, m, deterministic)
}
func (dst *ComputeAverageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComputeAverageResponse.Merge(dst, src)
}
func (m *ComputeAverageResponse) XXX_Size() int {
	return xxx_messageInfo_ComputeAverageResponse.Size(m)
}
func (m *ComputeAverageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ComputeAverageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ComputeAverageResponse proto.InternalMessageInfo

func (m *ComputeAverageResponse) GetAverage() float32 {
	if m != nil {
		return m.Average
	}
	return 0
}

func init() {
	proto.RegisterType((*Sum)(nil), "calculator.Sum")
	proto.RegisterType((*SumRequest)(nil), "calculator.SumRequest")
	proto.RegisterType((*SumResponse)(nil), "calculator.SumResponse")
	proto.RegisterType((*PrimeNumberRequest)(nil), "calculator.PrimeNumberRequest")
	proto.RegisterType((*PrimeNumberResponse)(nil), "calculator.PrimeNumberResponse")
	proto.RegisterType((*ComputeAverageRequest)(nil), "calculator.ComputeAverageRequest")
	proto.RegisterType((*ComputeAverageResponse)(nil), "calculator.ComputeAverageResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
	// server streaming
	PrimeStream(ctx context.Context, in *PrimeNumberRequest, opts ...grpc.CallOption) (CalculatorService_PrimeStreamClient, error)
	// client streaming
	ComputeAverage(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_ComputeAverageClient, error)
}

type calculatorServiceClient struct {
	cc *grpc.ClientConn
}

func NewCalculatorServiceClient(cc *grpc.ClientConn) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) PrimeStream(ctx context.Context, in *PrimeNumberRequest, opts ...grpc.CallOption) (CalculatorService_PrimeStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[0], "/calculator.CalculatorService/PrimeStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServicePrimeStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalculatorService_PrimeStreamClient interface {
	Recv() (*PrimeNumberResponse, error)
	grpc.ClientStream
}

type calculatorServicePrimeStreamClient struct {
	grpc.ClientStream
}

func (x *calculatorServicePrimeStreamClient) Recv() (*PrimeNumberResponse, error) {
	m := new(PrimeNumberResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) ComputeAverage(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_ComputeAverageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[1], "/calculator.CalculatorService/ComputeAverage", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceComputeAverageClient{stream}
	return x, nil
}

type CalculatorService_ComputeAverageClient interface {
	Send(*ComputeAverageRequest) error
	CloseAndRecv() (*ComputeAverageResponse, error)
	grpc.ClientStream
}

type calculatorServiceComputeAverageClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceComputeAverageClient) Send(m *ComputeAverageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceComputeAverageClient) CloseAndRecv() (*ComputeAverageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ComputeAverageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
type CalculatorServiceServer interface {
	Sum(context.Context, *SumRequest) (*SumResponse, error)
	// server streaming
	PrimeStream(*PrimeNumberRequest, CalculatorService_PrimeStreamServer) error
	// client streaming
	ComputeAverage(CalculatorService_ComputeAverageServer) error
}

func RegisterCalculatorServiceServer(s *grpc.Server, srv CalculatorServiceServer) {
	s.RegisterService(&_CalculatorService_serviceDesc, srv)
}

func _CalculatorService_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_PrimeStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeNumberRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServiceServer).PrimeStream(m, &calculatorServicePrimeStreamServer{stream})
}

type CalculatorService_PrimeStreamServer interface {
	Send(*PrimeNumberResponse) error
	grpc.ServerStream
}

type calculatorServicePrimeStreamServer struct {
	grpc.ServerStream
}

func (x *calculatorServicePrimeStreamServer) Send(m *PrimeNumberResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CalculatorService_ComputeAverage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).ComputeAverage(&calculatorServiceComputeAverageServer{stream})
}

type CalculatorService_ComputeAverageServer interface {
	SendAndClose(*ComputeAverageResponse) error
	Recv() (*ComputeAverageRequest, error)
	grpc.ServerStream
}

type calculatorServiceComputeAverageServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceComputeAverageServer) SendAndClose(m *ComputeAverageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceComputeAverageServer) Recv() (*ComputeAverageRequest, error) {
	m := new(ComputeAverageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _CalculatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _CalculatorService_Sum_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PrimeStream",
			Handler:       _CalculatorService_PrimeStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ComputeAverage",
			Handler:       _CalculatorService_ComputeAverage_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "calculator/calculatorpb/calculator.proto",
}

func init() {
	proto.RegisterFile("calculator/calculatorpb/calculator.proto", fileDescriptor_calculator_be6d262a2f835227)
}

var fileDescriptor_calculator_be6d262a2f835227 = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcd, 0x4f, 0xfa, 0x40,
	0x10, 0xfd, 0x15, 0x02, 0xbf, 0x64, 0x8a, 0x18, 0x57, 0xad, 0x84, 0x83, 0xca, 0x1a, 0x13, 0x12,
	0x0d, 0x18, 0xbc, 0x78, 0x55, 0xce, 0x1a, 0x42, 0x4f, 0x7a, 0x31, 0x6d, 0x1d, 0x4d, 0x93, 0x6e,
	0xb7, 0xee, 0x07, 0x7f, 0xbc, 0x27, 0xd3, 0xdd, 0xad, 0xb4, 0xe0, 0xc7, 0x6d, 0xdf, 0xdb, 0x37,
	0x6f, 0x66, 0x5e, 0x06, 0xc6, 0x49, 0x94, 0x25, 0x3a, 0x8b, 0x14, 0x17, 0xd3, 0xf5, 0xb3, 0x88,
	0x6b, 0x60, 0x52, 0x08, 0xae, 0x38, 0x81, 0x35, 0x43, 0xef, 0xa1, 0x1d, 0x6a, 0x46, 0x46, 0xd0,
	0x7b, 0x4d, 0x85, 0x54, 0xcf, 0xb9, 0x66, 0x31, 0x8a, 0x81, 0x77, 0xea, 0x8d, 0x3b, 0x4b, 0xdf,
	0x70, 0x0f, 0x86, 0x22, 0x67, 0xb0, 0x23, 0x31, 0xe1, 0xf9, 0x4b, 0xa5, 0x69, 0x19, 0x4d, 0xcf,
	0x92, 0x56, 0x44, 0xa7, 0x00, 0xa1, 0x66, 0x4b, 0x7c, 0xd7, 0x28, 0x15, 0x19, 0x41, 0x5b, 0x6a,
	0x66, 0xcc, 0xfc, 0xd9, 0xee, 0xa4, 0x36, 0x48, 0x29, 0x2a, 0xff, 0xe8, 0x39, 0xf8, 0xa6, 0x40,
	0x16, 0x3c, 0x97, 0x48, 0x02, 0xe8, 0x0a, 0x94, 0x3a, 0x53, 0x6e, 0x02, 0x87, 0xe8, 0x25, 0x90,
	0x85, 0x48, 0x19, 0xda, 0x36, 0x95, 0x7f, 0x00, 0xdd, 0xc6, 0xbc, 0x0e, 0xd1, 0x0b, 0xd8, 0x6f,
	0xa8, 0x9d, 0xf9, 0x01, 0x74, 0x8a, 0x92, 0x76, 0x6a, 0x0b, 0xe8, 0x14, 0x0e, 0xe7, 0x9c, 0x15,
	0x5a, 0xe1, 0xed, 0x0a, 0x45, 0xf4, 0x86, 0x7f, 0xb9, 0xcf, 0x20, 0xd8, 0x2c, 0x70, 0x0d, 0x06,
	0xf0, 0x3f, 0xb2, 0x94, 0x29, 0x69, 0x2d, 0x2b, 0x38, 0xfb, 0xf0, 0x60, 0x6f, 0xfe, 0xb5, 0x7e,
	0x88, 0x62, 0x95, 0x26, 0x48, 0x6e, 0x6c, 0xf8, 0xc1, 0x66, 0x32, 0x76, 0x80, 0xe1, 0xd1, 0x16,
	0x6f, 0xfb, 0xd0, 0x7f, 0x64, 0x01, 0xbe, 0xd9, 0x30, 0x54, 0x02, 0x23, 0x46, 0x8e, 0xeb, 0xca,
	0xed, 0xa0, 0x86, 0x27, 0x3f, 0xfe, 0x5b, 0xc7, 0x2b, 0x8f, 0x3c, 0x42, 0xbf, 0xb9, 0x15, 0x19,
	0xd5, 0x8b, 0xbe, 0x8d, 0x68, 0x48, 0x7f, 0x93, 0x58, 0xeb, 0xb1, 0x77, 0xd7, 0x7f, 0xea, 0xd5,
	0x0f, 0x32, 0xee, 0x9a, 0x33, 0xbc, 0xfe, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x89, 0x03, 0x0d, 0x76,
	0xb2, 0x02, 0x00, 0x00,
}
