// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package calcpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	// unary API
	Sum(ctx context.Context, in *CalculatorRequest, opts ...grpc.CallOption) (*CalculatorResponse, error)
	// server streaming
	PrimeNumDecomposition(ctx context.Context, in *PrimeNumDecompRequest, opts ...grpc.CallOption) (CalculatorService_PrimeNumDecompositionClient, error)
	// client streaming
	Average(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_AverageClient, error)
	//bi-directional stream
	Maximum(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_MaximumClient, error)
}

type calculatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorServiceClient(cc grpc.ClientConnInterface) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) Sum(ctx context.Context, in *CalculatorRequest, opts ...grpc.CallOption) (*CalculatorResponse, error) {
	out := new(CalculatorResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) PrimeNumDecomposition(ctx context.Context, in *PrimeNumDecompRequest, opts ...grpc.CallOption) (CalculatorService_PrimeNumDecompositionClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalculatorService_ServiceDesc.Streams[0], "/calculator.CalculatorService/PrimeNumDecomposition", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServicePrimeNumDecompositionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalculatorService_PrimeNumDecompositionClient interface {
	Recv() (*PrimeNumDecompResponse, error)
	grpc.ClientStream
}

type calculatorServicePrimeNumDecompositionClient struct {
	grpc.ClientStream
}

func (x *calculatorServicePrimeNumDecompositionClient) Recv() (*PrimeNumDecompResponse, error) {
	m := new(PrimeNumDecompResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) Average(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_AverageClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalculatorService_ServiceDesc.Streams[1], "/calculator.CalculatorService/Average", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceAverageClient{stream}
	return x, nil
}

type CalculatorService_AverageClient interface {
	Send(*AverageRequest) error
	CloseAndRecv() (*AverageResponse, error)
	grpc.ClientStream
}

type calculatorServiceAverageClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceAverageClient) Send(m *AverageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceAverageClient) CloseAndRecv() (*AverageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AverageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) Maximum(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_MaximumClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalculatorService_ServiceDesc.Streams[2], "/calculator.CalculatorService/Maximum", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceMaximumClient{stream}
	return x, nil
}

type CalculatorService_MaximumClient interface {
	Send(*MaximumRequest) error
	Recv() (*MaximumResponse, error)
	grpc.ClientStream
}

type calculatorServiceMaximumClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceMaximumClient) Send(m *MaximumRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceMaximumClient) Recv() (*MaximumResponse, error) {
	m := new(MaximumResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
// All implementations must embed UnimplementedCalculatorServiceServer
// for forward compatibility
type CalculatorServiceServer interface {
	// unary API
	Sum(context.Context, *CalculatorRequest) (*CalculatorResponse, error)
	// server streaming
	PrimeNumDecomposition(*PrimeNumDecompRequest, CalculatorService_PrimeNumDecompositionServer) error
	// client streaming
	Average(CalculatorService_AverageServer) error
	//bi-directional stream
	Maximum(CalculatorService_MaximumServer) error
	mustEmbedUnimplementedCalculatorServiceServer()
}

// UnimplementedCalculatorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCalculatorServiceServer struct {
}

func (UnimplementedCalculatorServiceServer) Sum(context.Context, *CalculatorRequest) (*CalculatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (UnimplementedCalculatorServiceServer) PrimeNumDecomposition(*PrimeNumDecompRequest, CalculatorService_PrimeNumDecompositionServer) error {
	return status.Errorf(codes.Unimplemented, "method PrimeNumDecomposition not implemented")
}
func (UnimplementedCalculatorServiceServer) Average(CalculatorService_AverageServer) error {
	return status.Errorf(codes.Unimplemented, "method Average not implemented")
}
func (UnimplementedCalculatorServiceServer) Maximum(CalculatorService_MaximumServer) error {
	return status.Errorf(codes.Unimplemented, "method Maximum not implemented")
}
func (UnimplementedCalculatorServiceServer) mustEmbedUnimplementedCalculatorServiceServer() {}

// UnsafeCalculatorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalculatorServiceServer will
// result in compilation errors.
type UnsafeCalculatorServiceServer interface {
	mustEmbedUnimplementedCalculatorServiceServer()
}

func RegisterCalculatorServiceServer(s grpc.ServiceRegistrar, srv CalculatorServiceServer) {
	s.RegisterService(&CalculatorService_ServiceDesc, srv)
}

func _CalculatorService_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculatorRequest)
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
		return srv.(CalculatorServiceServer).Sum(ctx, req.(*CalculatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_PrimeNumDecomposition_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeNumDecompRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServiceServer).PrimeNumDecomposition(m, &calculatorServicePrimeNumDecompositionServer{stream})
}

type CalculatorService_PrimeNumDecompositionServer interface {
	Send(*PrimeNumDecompResponse) error
	grpc.ServerStream
}

type calculatorServicePrimeNumDecompositionServer struct {
	grpc.ServerStream
}

func (x *calculatorServicePrimeNumDecompositionServer) Send(m *PrimeNumDecompResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CalculatorService_Average_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).Average(&calculatorServiceAverageServer{stream})
}

type CalculatorService_AverageServer interface {
	SendAndClose(*AverageResponse) error
	Recv() (*AverageRequest, error)
	grpc.ServerStream
}

type calculatorServiceAverageServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceAverageServer) SendAndClose(m *AverageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceAverageServer) Recv() (*AverageRequest, error) {
	m := new(AverageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalculatorService_Maximum_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).Maximum(&calculatorServiceMaximumServer{stream})
}

type CalculatorService_MaximumServer interface {
	Send(*MaximumResponse) error
	Recv() (*MaximumRequest, error)
	grpc.ServerStream
}

type calculatorServiceMaximumServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceMaximumServer) Send(m *MaximumResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceMaximumServer) Recv() (*MaximumRequest, error) {
	m := new(MaximumRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculatorService_ServiceDesc is the grpc.ServiceDesc for CalculatorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CalculatorService_ServiceDesc = grpc.ServiceDesc{
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
			StreamName:    "PrimeNumDecomposition",
			Handler:       _CalculatorService_PrimeNumDecomposition_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Average",
			Handler:       _CalculatorService_Average_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Maximum",
			Handler:       _CalculatorService_Maximum_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "calculator/calcpb/calc.proto",
}
