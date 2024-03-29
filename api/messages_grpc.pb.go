// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: api/messages.proto

package api

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

// MessagesServiceClient is the client API for MessagesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessagesServiceClient interface {
	SendMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*MessagesStatusResponse, error)
	AllMessages(ctx context.Context, in *MessagesZeroRequest, opts ...grpc.CallOption) (*Messages, error)
}

type messagesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMessagesServiceClient(cc grpc.ClientConnInterface) MessagesServiceClient {
	return &messagesServiceClient{cc}
}

func (c *messagesServiceClient) SendMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*MessagesStatusResponse, error) {
	out := new(MessagesStatusResponse)
	err := c.cc.Invoke(ctx, "/api.MessagesService/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagesServiceClient) AllMessages(ctx context.Context, in *MessagesZeroRequest, opts ...grpc.CallOption) (*Messages, error) {
	out := new(Messages)
	err := c.cc.Invoke(ctx, "/api.MessagesService/AllMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessagesServiceServer is the server API for MessagesService service.
// All implementations should embed UnimplementedMessagesServiceServer
// for forward compatibility
type MessagesServiceServer interface {
	SendMessage(context.Context, *Message) (*MessagesStatusResponse, error)
	AllMessages(context.Context, *MessagesZeroRequest) (*Messages, error)
}

// UnimplementedMessagesServiceServer should be embedded to have forward compatible implementations.
type UnimplementedMessagesServiceServer struct {
}

func (UnimplementedMessagesServiceServer) SendMessage(context.Context, *Message) (*MessagesStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedMessagesServiceServer) AllMessages(context.Context, *MessagesZeroRequest) (*Messages, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllMessages not implemented")
}

// UnsafeMessagesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessagesServiceServer will
// result in compilation errors.
type UnsafeMessagesServiceServer interface {
	mustEmbedUnimplementedMessagesServiceServer()
}

func RegisterMessagesServiceServer(s grpc.ServiceRegistrar, srv MessagesServiceServer) {
	s.RegisterService(&MessagesService_ServiceDesc, srv)
}

func _MessagesService_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagesServiceServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.MessagesService/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagesServiceServer).SendMessage(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessagesService_AllMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessagesZeroRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagesServiceServer).AllMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.MessagesService/AllMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagesServiceServer).AllMessages(ctx, req.(*MessagesZeroRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MessagesService_ServiceDesc is the grpc.ServiceDesc for MessagesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessagesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.MessagesService",
	HandlerType: (*MessagesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessage",
			Handler:    _MessagesService_SendMessage_Handler,
		},
		{
			MethodName: "AllMessages",
			Handler:    _MessagesService_AllMessages_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/messages.proto",
}
