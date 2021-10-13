// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package taskService

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

// TaskAPIClient is the client API for TaskAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaskAPIClient interface {
	GetTasks(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TaskReply, error)
}

type taskAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskAPIClient(cc grpc.ClientConnInterface) TaskAPIClient {
	return &taskAPIClient{cc}
}

func (c *taskAPIClient) GetTasks(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TaskReply, error) {
	out := new(TaskReply)
	err := c.cc.Invoke(ctx, "/task.TaskAPI/GetTasks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskAPIServer is the server API for TaskAPI service.
// All implementations must embed UnimplementedTaskAPIServer
// for forward compatibility
type TaskAPIServer interface {
	GetTasks(context.Context, *TaskRequest) (*TaskReply, error)
	mustEmbedUnimplementedTaskAPIServer()
}

// UnimplementedTaskAPIServer must be embedded to have forward compatible implementations.
type UnimplementedTaskAPIServer struct {
}

func (UnimplementedTaskAPIServer) GetTasks(context.Context, *TaskRequest) (*TaskReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTasks not implemented")
}
func (UnimplementedTaskAPIServer) mustEmbedUnimplementedTaskAPIServer() {}

// UnsafeTaskAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaskAPIServer will
// result in compilation errors.
type UnsafeTaskAPIServer interface {
	mustEmbedUnimplementedTaskAPIServer()
}

func RegisterTaskAPIServer(s grpc.ServiceRegistrar, srv TaskAPIServer) {
	s.RegisterService(&TaskAPI_ServiceDesc, srv)
}

func _TaskAPI_GetTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskAPIServer).GetTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.TaskAPI/GetTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskAPIServer).GetTasks(ctx, req.(*TaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TaskAPI_ServiceDesc is the grpc.ServiceDesc for TaskAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TaskAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "task.TaskAPI",
	HandlerType: (*TaskAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTasks",
			Handler:    _TaskAPI_GetTasks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "task.proto",
}
