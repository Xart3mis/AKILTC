// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: ClientData.proto

package pb

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

// ConsumerClient is the client API for Consumer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConsumerClient interface {
	SubscribeOnScreenText(ctx context.Context, in *ClientDataRequest, opts ...grpc.CallOption) (Consumer_SubscribeOnScreenTextClient, error)
	GetCommand(ctx context.Context, in *ClientDataRequest, opts ...grpc.CallOption) (*ClientExecData, error)
	SetCommandOutput(ctx context.Context, in *ClientExecOutput, opts ...grpc.CallOption) (*Void, error)
	GetFlood(ctx context.Context, in *Void, opts ...grpc.CallOption) (*FloodData, error)
	SetFloodOutput(ctx context.Context, opts ...grpc.CallOption) (Consumer_SetFloodOutputClient, error)
}

type consumerClient struct {
	cc grpc.ClientConnInterface
}

func NewConsumerClient(cc grpc.ClientConnInterface) ConsumerClient {
	return &consumerClient{cc}
}

func (c *consumerClient) SubscribeOnScreenText(ctx context.Context, in *ClientDataRequest, opts ...grpc.CallOption) (Consumer_SubscribeOnScreenTextClient, error) {
	stream, err := c.cc.NewStream(ctx, &Consumer_ServiceDesc.Streams[0], "/pb.Consumer/SubscribeOnScreenText", opts...)
	if err != nil {
		return nil, err
	}
	x := &consumerSubscribeOnScreenTextClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Consumer_SubscribeOnScreenTextClient interface {
	Recv() (*ClientDataOnScreenTextResponse, error)
	grpc.ClientStream
}

type consumerSubscribeOnScreenTextClient struct {
	grpc.ClientStream
}

func (x *consumerSubscribeOnScreenTextClient) Recv() (*ClientDataOnScreenTextResponse, error) {
	m := new(ClientDataOnScreenTextResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *consumerClient) GetCommand(ctx context.Context, in *ClientDataRequest, opts ...grpc.CallOption) (*ClientExecData, error) {
	out := new(ClientExecData)
	err := c.cc.Invoke(ctx, "/pb.Consumer/GetCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consumerClient) SetCommandOutput(ctx context.Context, in *ClientExecOutput, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/pb.Consumer/SetCommandOutput", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consumerClient) GetFlood(ctx context.Context, in *Void, opts ...grpc.CallOption) (*FloodData, error) {
	out := new(FloodData)
	err := c.cc.Invoke(ctx, "/pb.Consumer/GetFlood", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consumerClient) SetFloodOutput(ctx context.Context, opts ...grpc.CallOption) (Consumer_SetFloodOutputClient, error) {
	stream, err := c.cc.NewStream(ctx, &Consumer_ServiceDesc.Streams[1], "/pb.Consumer/SetFloodOutput", opts...)
	if err != nil {
		return nil, err
	}
	x := &consumerSetFloodOutputClient{stream}
	return x, nil
}

type Consumer_SetFloodOutputClient interface {
	Send(*FloodOutput) error
	CloseAndRecv() (*Void, error)
	grpc.ClientStream
}

type consumerSetFloodOutputClient struct {
	grpc.ClientStream
}

func (x *consumerSetFloodOutputClient) Send(m *FloodOutput) error {
	return x.ClientStream.SendMsg(m)
}

func (x *consumerSetFloodOutputClient) CloseAndRecv() (*Void, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Void)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ConsumerServer is the server API for Consumer service.
// All implementations must embed UnimplementedConsumerServer
// for forward compatibility
type ConsumerServer interface {
	SubscribeOnScreenText(*ClientDataRequest, Consumer_SubscribeOnScreenTextServer) error
	GetCommand(context.Context, *ClientDataRequest) (*ClientExecData, error)
	SetCommandOutput(context.Context, *ClientExecOutput) (*Void, error)
	GetFlood(context.Context, *Void) (*FloodData, error)
	SetFloodOutput(Consumer_SetFloodOutputServer) error
	mustEmbedUnimplementedConsumerServer()
}

// UnimplementedConsumerServer must be embedded to have forward compatible implementations.
type UnimplementedConsumerServer struct {
}

func (UnimplementedConsumerServer) SubscribeOnScreenText(*ClientDataRequest, Consumer_SubscribeOnScreenTextServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeOnScreenText not implemented")
}
func (UnimplementedConsumerServer) GetCommand(context.Context, *ClientDataRequest) (*ClientExecData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommand not implemented")
}
func (UnimplementedConsumerServer) SetCommandOutput(context.Context, *ClientExecOutput) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetCommandOutput not implemented")
}
func (UnimplementedConsumerServer) GetFlood(context.Context, *Void) (*FloodData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFlood not implemented")
}
func (UnimplementedConsumerServer) SetFloodOutput(Consumer_SetFloodOutputServer) error {
	return status.Errorf(codes.Unimplemented, "method SetFloodOutput not implemented")
}
func (UnimplementedConsumerServer) mustEmbedUnimplementedConsumerServer() {}

// UnsafeConsumerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConsumerServer will
// result in compilation errors.
type UnsafeConsumerServer interface {
	mustEmbedUnimplementedConsumerServer()
}

func RegisterConsumerServer(s grpc.ServiceRegistrar, srv ConsumerServer) {
	s.RegisterService(&Consumer_ServiceDesc, srv)
}

func _Consumer_SubscribeOnScreenText_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ClientDataRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ConsumerServer).SubscribeOnScreenText(m, &consumerSubscribeOnScreenTextServer{stream})
}

type Consumer_SubscribeOnScreenTextServer interface {
	Send(*ClientDataOnScreenTextResponse) error
	grpc.ServerStream
}

type consumerSubscribeOnScreenTextServer struct {
	grpc.ServerStream
}

func (x *consumerSubscribeOnScreenTextServer) Send(m *ClientDataOnScreenTextResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Consumer_GetCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsumerServer).GetCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Consumer/GetCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsumerServer).GetCommand(ctx, req.(*ClientDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Consumer_SetCommandOutput_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientExecOutput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsumerServer).SetCommandOutput(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Consumer/SetCommandOutput",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsumerServer).SetCommandOutput(ctx, req.(*ClientExecOutput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Consumer_GetFlood_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsumerServer).GetFlood(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Consumer/GetFlood",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsumerServer).GetFlood(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _Consumer_SetFloodOutput_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ConsumerServer).SetFloodOutput(&consumerSetFloodOutputServer{stream})
}

type Consumer_SetFloodOutputServer interface {
	SendAndClose(*Void) error
	Recv() (*FloodOutput, error)
	grpc.ServerStream
}

type consumerSetFloodOutputServer struct {
	grpc.ServerStream
}

func (x *consumerSetFloodOutputServer) SendAndClose(m *Void) error {
	return x.ServerStream.SendMsg(m)
}

func (x *consumerSetFloodOutputServer) Recv() (*FloodOutput, error) {
	m := new(FloodOutput)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Consumer_ServiceDesc is the grpc.ServiceDesc for Consumer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Consumer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Consumer",
	HandlerType: (*ConsumerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCommand",
			Handler:    _Consumer_GetCommand_Handler,
		},
		{
			MethodName: "SetCommandOutput",
			Handler:    _Consumer_SetCommandOutput_Handler,
		},
		{
			MethodName: "GetFlood",
			Handler:    _Consumer_GetFlood_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeOnScreenText",
			Handler:       _Consumer_SubscribeOnScreenText_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SetFloodOutput",
			Handler:       _Consumer_SetFloodOutput_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "ClientData.proto",
}
