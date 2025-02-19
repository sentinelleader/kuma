// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0
// source: mesh/v1alpha1/mux.proto

package v1alpha1

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

// MultiplexServiceClient is the client API for MultiplexService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MultiplexServiceClient interface {
	StreamMessage(ctx context.Context, opts ...grpc.CallOption) (MultiplexService_StreamMessageClient, error)
}

type multiplexServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMultiplexServiceClient(cc grpc.ClientConnInterface) MultiplexServiceClient {
	return &multiplexServiceClient{cc}
}

func (c *multiplexServiceClient) StreamMessage(ctx context.Context, opts ...grpc.CallOption) (MultiplexService_StreamMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &MultiplexService_ServiceDesc.Streams[0], "/kuma.mesh.v1alpha1.MultiplexService/StreamMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &multiplexServiceStreamMessageClient{stream}
	return x, nil
}

type MultiplexService_StreamMessageClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type multiplexServiceStreamMessageClient struct {
	grpc.ClientStream
}

func (x *multiplexServiceStreamMessageClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *multiplexServiceStreamMessageClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MultiplexServiceServer is the server API for MultiplexService service.
// All implementations must embed UnimplementedMultiplexServiceServer
// for forward compatibility
type MultiplexServiceServer interface {
	StreamMessage(MultiplexService_StreamMessageServer) error
	mustEmbedUnimplementedMultiplexServiceServer()
}

// UnimplementedMultiplexServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMultiplexServiceServer struct {
}

func (UnimplementedMultiplexServiceServer) StreamMessage(MultiplexService_StreamMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamMessage not implemented")
}
func (UnimplementedMultiplexServiceServer) mustEmbedUnimplementedMultiplexServiceServer() {}

// UnsafeMultiplexServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MultiplexServiceServer will
// result in compilation errors.
type UnsafeMultiplexServiceServer interface {
	mustEmbedUnimplementedMultiplexServiceServer()
}

func RegisterMultiplexServiceServer(s grpc.ServiceRegistrar, srv MultiplexServiceServer) {
	s.RegisterService(&MultiplexService_ServiceDesc, srv)
}

func _MultiplexService_StreamMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MultiplexServiceServer).StreamMessage(&multiplexServiceStreamMessageServer{stream})
}

type MultiplexService_StreamMessageServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type multiplexServiceStreamMessageServer struct {
	grpc.ServerStream
}

func (x *multiplexServiceStreamMessageServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *multiplexServiceStreamMessageServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MultiplexService_ServiceDesc is the grpc.ServiceDesc for MultiplexService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MultiplexService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kuma.mesh.v1alpha1.MultiplexService",
	HandlerType: (*MultiplexServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamMessage",
			Handler:       _MultiplexService_StreamMessage_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "mesh/v1alpha1/mux.proto",
}
