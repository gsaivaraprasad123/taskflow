// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.1
// source: api.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	CoordinatorService_UpdateTaskStatus_FullMethodName = "/grpc.CoordinatorService/UpdateTaskStatus"
)

// CoordinatorServiceClient is the client API for CoordinatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CoordinatorServiceClient interface {
	UpdateTaskStatus(ctx context.Context, in *UpdateTaskStatusRequest, opts ...grpc.CallOption) (*UpdateTaskStatusResponse, error)
}

type coordinatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCoordinatorServiceClient(cc grpc.ClientConnInterface) CoordinatorServiceClient {
	return &coordinatorServiceClient{cc}
}

func (c *coordinatorServiceClient) UpdateTaskStatus(ctx context.Context, in *UpdateTaskStatusRequest, opts ...grpc.CallOption) (*UpdateTaskStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateTaskStatusResponse)
	err := c.cc.Invoke(ctx, CoordinatorService_UpdateTaskStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoordinatorServiceServer is the server API for CoordinatorService service.
// All implementations must embed UnimplementedCoordinatorServiceServer
// for forward compatibility.
type CoordinatorServiceServer interface {
	UpdateTaskStatus(context.Context, *UpdateTaskStatusRequest) (*UpdateTaskStatusResponse, error)
	mustEmbedUnimplementedCoordinatorServiceServer()
}

// UnimplementedCoordinatorServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCoordinatorServiceServer struct{}

func (UnimplementedCoordinatorServiceServer) UpdateTaskStatus(context.Context, *UpdateTaskStatusRequest) (*UpdateTaskStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTaskStatus not implemented")
}
func (UnimplementedCoordinatorServiceServer) mustEmbedUnimplementedCoordinatorServiceServer() {}
func (UnimplementedCoordinatorServiceServer) testEmbeddedByValue()                            {}

// UnsafeCoordinatorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CoordinatorServiceServer will
// result in compilation errors.
type UnsafeCoordinatorServiceServer interface {
	mustEmbedUnimplementedCoordinatorServiceServer()
}

func RegisterCoordinatorServiceServer(s grpc.ServiceRegistrar, srv CoordinatorServiceServer) {
	// If the following call pancis, it indicates UnimplementedCoordinatorServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CoordinatorService_ServiceDesc, srv)
}

func _CoordinatorService_UpdateTaskStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTaskStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServiceServer).UpdateTaskStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CoordinatorService_UpdateTaskStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServiceServer).UpdateTaskStatus(ctx, req.(*UpdateTaskStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CoordinatorService_ServiceDesc is the grpc.ServiceDesc for CoordinatorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CoordinatorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.CoordinatorService",
	HandlerType: (*CoordinatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateTaskStatus",
			Handler:    _CoordinatorService_UpdateTaskStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

const (
	QueueService_EnqueueTask_FullMethodName = "/grpc.QueueService/EnqueueTask"
)

// QueueServiceClient is the client API for QueueService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueueServiceClient interface {
	EnqueueTask(ctx context.Context, in *EnqueueTaskRequest, opts ...grpc.CallOption) (*EnqueueTaskResponse, error)
}

type queueServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewQueueServiceClient(cc grpc.ClientConnInterface) QueueServiceClient {
	return &queueServiceClient{cc}
}

func (c *queueServiceClient) EnqueueTask(ctx context.Context, in *EnqueueTaskRequest, opts ...grpc.CallOption) (*EnqueueTaskResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EnqueueTaskResponse)
	err := c.cc.Invoke(ctx, QueueService_EnqueueTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueueServiceServer is the server API for QueueService service.
// All implementations must embed UnimplementedQueueServiceServer
// for forward compatibility.
type QueueServiceServer interface {
	EnqueueTask(context.Context, *EnqueueTaskRequest) (*EnqueueTaskResponse, error)
	mustEmbedUnimplementedQueueServiceServer()
}

// UnimplementedQueueServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedQueueServiceServer struct{}

func (UnimplementedQueueServiceServer) EnqueueTask(context.Context, *EnqueueTaskRequest) (*EnqueueTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnqueueTask not implemented")
}
func (UnimplementedQueueServiceServer) mustEmbedUnimplementedQueueServiceServer() {}
func (UnimplementedQueueServiceServer) testEmbeddedByValue()                      {}

// UnsafeQueueServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueueServiceServer will
// result in compilation errors.
type UnsafeQueueServiceServer interface {
	mustEmbedUnimplementedQueueServiceServer()
}

func RegisterQueueServiceServer(s grpc.ServiceRegistrar, srv QueueServiceServer) {
	// If the following call pancis, it indicates UnimplementedQueueServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&QueueService_ServiceDesc, srv)
}

func _QueueService_EnqueueTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnqueueTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueServiceServer).EnqueueTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QueueService_EnqueueTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueServiceServer).EnqueueTask(ctx, req.(*EnqueueTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// QueueService_ServiceDesc is the grpc.ServiceDesc for QueueService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QueueService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.QueueService",
	HandlerType: (*QueueServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EnqueueTask",
			Handler:    _QueueService_EnqueueTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
