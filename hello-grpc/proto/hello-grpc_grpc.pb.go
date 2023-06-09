// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: hello-grpc.proto

package proto

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

const (
	HelloWorldlService_HelloWorld_FullMethodName                    = "/hellogrpc.HelloWorldlService/HelloWorld"
	HelloWorldlService_HelloWorldWithServerSteaming_FullMethodName  = "/hellogrpc.HelloWorldlService/HelloWorldWithServerSteaming"
	HelloWorldlService_HelloWorldWithClientStreaming_FullMethodName = "/hellogrpc.HelloWorldlService/HelloWorldWithClientStreaming"
	HelloWorldlService_HelloWorldWithBidirectional_FullMethodName   = "/hellogrpc.HelloWorldlService/HelloWorldWithBidirectional"
	HelloWorldlService_HelloWorldWithErrorMsg_FullMethodName        = "/hellogrpc.HelloWorldlService/HelloWorldWithErrorMsg"
	HelloWorldlService_HelloWorldWithDeadLines_FullMethodName       = "/hellogrpc.HelloWorldlService/HelloWorldWithDeadLines"
)

// HelloWorldlServiceClient is the client API for HelloWorldlService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HelloWorldlServiceClient interface {
	// implement unary GRPC
	HelloWorld(ctx context.Context, in *HelloWorldRequest, opts ...grpc.CallOption) (*HelloWorldResponse, error)
	// implement Server streaming
	HelloWorldWithServerSteaming(ctx context.Context, in *HelloWorldRequest, opts ...grpc.CallOption) (HelloWorldlService_HelloWorldWithServerSteamingClient, error)
	// implement client streaming
	HelloWorldWithClientStreaming(ctx context.Context, opts ...grpc.CallOption) (HelloWorldlService_HelloWorldWithClientStreamingClient, error)
	// implement bidirectional grpc
	HelloWorldWithBidirectional(ctx context.Context, opts ...grpc.CallOption) (HelloWorldlService_HelloWorldWithBidirectionalClient, error)
	// implement grpc with error message
	HelloWorldWithErrorMsg(ctx context.Context, in *HelloWorldRequest, opts ...grpc.CallOption) (*HelloWorldResponse, error)
	// implment helloworld with deadline
	HelloWorldWithDeadLines(ctx context.Context, in *HelloWorldRequest, opts ...grpc.CallOption) (*HelloWorldResponse, error)
}

type helloWorldlServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloWorldlServiceClient(cc grpc.ClientConnInterface) HelloWorldlServiceClient {
	return &helloWorldlServiceClient{cc}
}

func (c *helloWorldlServiceClient) HelloWorld(ctx context.Context, in *HelloWorldRequest, opts ...grpc.CallOption) (*HelloWorldResponse, error) {
	out := new(HelloWorldResponse)
	err := c.cc.Invoke(ctx, HelloWorldlService_HelloWorld_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloWorldlServiceClient) HelloWorldWithServerSteaming(ctx context.Context, in *HelloWorldRequest, opts ...grpc.CallOption) (HelloWorldlService_HelloWorldWithServerSteamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &HelloWorldlService_ServiceDesc.Streams[0], HelloWorldlService_HelloWorldWithServerSteaming_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &helloWorldlServiceHelloWorldWithServerSteamingClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type HelloWorldlService_HelloWorldWithServerSteamingClient interface {
	Recv() (*HelloWorldResponse, error)
	grpc.ClientStream
}

type helloWorldlServiceHelloWorldWithServerSteamingClient struct {
	grpc.ClientStream
}

func (x *helloWorldlServiceHelloWorldWithServerSteamingClient) Recv() (*HelloWorldResponse, error) {
	m := new(HelloWorldResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *helloWorldlServiceClient) HelloWorldWithClientStreaming(ctx context.Context, opts ...grpc.CallOption) (HelloWorldlService_HelloWorldWithClientStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &HelloWorldlService_ServiceDesc.Streams[1], HelloWorldlService_HelloWorldWithClientStreaming_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &helloWorldlServiceHelloWorldWithClientStreamingClient{stream}
	return x, nil
}

type HelloWorldlService_HelloWorldWithClientStreamingClient interface {
	Send(*HelloWorldRequest) error
	CloseAndRecv() (*HelloWorldResponse, error)
	grpc.ClientStream
}

type helloWorldlServiceHelloWorldWithClientStreamingClient struct {
	grpc.ClientStream
}

func (x *helloWorldlServiceHelloWorldWithClientStreamingClient) Send(m *HelloWorldRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *helloWorldlServiceHelloWorldWithClientStreamingClient) CloseAndRecv() (*HelloWorldResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(HelloWorldResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *helloWorldlServiceClient) HelloWorldWithBidirectional(ctx context.Context, opts ...grpc.CallOption) (HelloWorldlService_HelloWorldWithBidirectionalClient, error) {
	stream, err := c.cc.NewStream(ctx, &HelloWorldlService_ServiceDesc.Streams[2], HelloWorldlService_HelloWorldWithBidirectional_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &helloWorldlServiceHelloWorldWithBidirectionalClient{stream}
	return x, nil
}

type HelloWorldlService_HelloWorldWithBidirectionalClient interface {
	Send(*HelloWorldRequest) error
	Recv() (*HelloWorldResponse, error)
	grpc.ClientStream
}

type helloWorldlServiceHelloWorldWithBidirectionalClient struct {
	grpc.ClientStream
}

func (x *helloWorldlServiceHelloWorldWithBidirectionalClient) Send(m *HelloWorldRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *helloWorldlServiceHelloWorldWithBidirectionalClient) Recv() (*HelloWorldResponse, error) {
	m := new(HelloWorldResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *helloWorldlServiceClient) HelloWorldWithErrorMsg(ctx context.Context, in *HelloWorldRequest, opts ...grpc.CallOption) (*HelloWorldResponse, error) {
	out := new(HelloWorldResponse)
	err := c.cc.Invoke(ctx, HelloWorldlService_HelloWorldWithErrorMsg_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloWorldlServiceClient) HelloWorldWithDeadLines(ctx context.Context, in *HelloWorldRequest, opts ...grpc.CallOption) (*HelloWorldResponse, error) {
	out := new(HelloWorldResponse)
	err := c.cc.Invoke(ctx, HelloWorldlService_HelloWorldWithDeadLines_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloWorldlServiceServer is the server API for HelloWorldlService service.
// All implementations must embed UnimplementedHelloWorldlServiceServer
// for forward compatibility
type HelloWorldlServiceServer interface {
	// implement unary GRPC
	HelloWorld(context.Context, *HelloWorldRequest) (*HelloWorldResponse, error)
	// implement Server streaming
	HelloWorldWithServerSteaming(*HelloWorldRequest, HelloWorldlService_HelloWorldWithServerSteamingServer) error
	// implement client streaming
	HelloWorldWithClientStreaming(HelloWorldlService_HelloWorldWithClientStreamingServer) error
	// implement bidirectional grpc
	HelloWorldWithBidirectional(HelloWorldlService_HelloWorldWithBidirectionalServer) error
	// implement grpc with error message
	HelloWorldWithErrorMsg(context.Context, *HelloWorldRequest) (*HelloWorldResponse, error)
	// implment helloworld with deadline
	HelloWorldWithDeadLines(context.Context, *HelloWorldRequest) (*HelloWorldResponse, error)
	mustEmbedUnimplementedHelloWorldlServiceServer()
}

// UnimplementedHelloWorldlServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHelloWorldlServiceServer struct {
}

func (UnimplementedHelloWorldlServiceServer) HelloWorld(context.Context, *HelloWorldRequest) (*HelloWorldResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HelloWorld not implemented")
}
func (UnimplementedHelloWorldlServiceServer) HelloWorldWithServerSteaming(*HelloWorldRequest, HelloWorldlService_HelloWorldWithServerSteamingServer) error {
	return status.Errorf(codes.Unimplemented, "method HelloWorldWithServerSteaming not implemented")
}
func (UnimplementedHelloWorldlServiceServer) HelloWorldWithClientStreaming(HelloWorldlService_HelloWorldWithClientStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method HelloWorldWithClientStreaming not implemented")
}
func (UnimplementedHelloWorldlServiceServer) HelloWorldWithBidirectional(HelloWorldlService_HelloWorldWithBidirectionalServer) error {
	return status.Errorf(codes.Unimplemented, "method HelloWorldWithBidirectional not implemented")
}
func (UnimplementedHelloWorldlServiceServer) HelloWorldWithErrorMsg(context.Context, *HelloWorldRequest) (*HelloWorldResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HelloWorldWithErrorMsg not implemented")
}
func (UnimplementedHelloWorldlServiceServer) HelloWorldWithDeadLines(context.Context, *HelloWorldRequest) (*HelloWorldResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HelloWorldWithDeadLines not implemented")
}
func (UnimplementedHelloWorldlServiceServer) mustEmbedUnimplementedHelloWorldlServiceServer() {}

// UnsafeHelloWorldlServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HelloWorldlServiceServer will
// result in compilation errors.
type UnsafeHelloWorldlServiceServer interface {
	mustEmbedUnimplementedHelloWorldlServiceServer()
}

func RegisterHelloWorldlServiceServer(s grpc.ServiceRegistrar, srv HelloWorldlServiceServer) {
	s.RegisterService(&HelloWorldlService_ServiceDesc, srv)
}

func _HelloWorldlService_HelloWorld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloWorldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloWorldlServiceServer).HelloWorld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HelloWorldlService_HelloWorld_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloWorldlServiceServer).HelloWorld(ctx, req.(*HelloWorldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloWorldlService_HelloWorldWithServerSteaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HelloWorldRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HelloWorldlServiceServer).HelloWorldWithServerSteaming(m, &helloWorldlServiceHelloWorldWithServerSteamingServer{stream})
}

type HelloWorldlService_HelloWorldWithServerSteamingServer interface {
	Send(*HelloWorldResponse) error
	grpc.ServerStream
}

type helloWorldlServiceHelloWorldWithServerSteamingServer struct {
	grpc.ServerStream
}

func (x *helloWorldlServiceHelloWorldWithServerSteamingServer) Send(m *HelloWorldResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _HelloWorldlService_HelloWorldWithClientStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HelloWorldlServiceServer).HelloWorldWithClientStreaming(&helloWorldlServiceHelloWorldWithClientStreamingServer{stream})
}

type HelloWorldlService_HelloWorldWithClientStreamingServer interface {
	SendAndClose(*HelloWorldResponse) error
	Recv() (*HelloWorldRequest, error)
	grpc.ServerStream
}

type helloWorldlServiceHelloWorldWithClientStreamingServer struct {
	grpc.ServerStream
}

func (x *helloWorldlServiceHelloWorldWithClientStreamingServer) SendAndClose(m *HelloWorldResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *helloWorldlServiceHelloWorldWithClientStreamingServer) Recv() (*HelloWorldRequest, error) {
	m := new(HelloWorldRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _HelloWorldlService_HelloWorldWithBidirectional_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HelloWorldlServiceServer).HelloWorldWithBidirectional(&helloWorldlServiceHelloWorldWithBidirectionalServer{stream})
}

type HelloWorldlService_HelloWorldWithBidirectionalServer interface {
	Send(*HelloWorldResponse) error
	Recv() (*HelloWorldRequest, error)
	grpc.ServerStream
}

type helloWorldlServiceHelloWorldWithBidirectionalServer struct {
	grpc.ServerStream
}

func (x *helloWorldlServiceHelloWorldWithBidirectionalServer) Send(m *HelloWorldResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *helloWorldlServiceHelloWorldWithBidirectionalServer) Recv() (*HelloWorldRequest, error) {
	m := new(HelloWorldRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _HelloWorldlService_HelloWorldWithErrorMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloWorldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloWorldlServiceServer).HelloWorldWithErrorMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HelloWorldlService_HelloWorldWithErrorMsg_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloWorldlServiceServer).HelloWorldWithErrorMsg(ctx, req.(*HelloWorldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloWorldlService_HelloWorldWithDeadLines_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloWorldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloWorldlServiceServer).HelloWorldWithDeadLines(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HelloWorldlService_HelloWorldWithDeadLines_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloWorldlServiceServer).HelloWorldWithDeadLines(ctx, req.(*HelloWorldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HelloWorldlService_ServiceDesc is the grpc.ServiceDesc for HelloWorldlService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HelloWorldlService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hellogrpc.HelloWorldlService",
	HandlerType: (*HelloWorldlServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HelloWorld",
			Handler:    _HelloWorldlService_HelloWorld_Handler,
		},
		{
			MethodName: "HelloWorldWithErrorMsg",
			Handler:    _HelloWorldlService_HelloWorldWithErrorMsg_Handler,
		},
		{
			MethodName: "HelloWorldWithDeadLines",
			Handler:    _HelloWorldlService_HelloWorldWithDeadLines_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "HelloWorldWithServerSteaming",
			Handler:       _HelloWorldlService_HelloWorldWithServerSteaming_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "HelloWorldWithClientStreaming",
			Handler:       _HelloWorldlService_HelloWorldWithClientStreaming_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "HelloWorldWithBidirectional",
			Handler:       _HelloWorldlService_HelloWorldWithBidirectional_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "hello-grpc.proto",
}
