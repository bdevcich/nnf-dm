// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// RsyncDataMoverClient is the client API for RsyncDataMover service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RsyncDataMoverClient interface {
	// Create sends a new data movement request identified by source and destination fields. It returns
	// a response containing a unique identifier which can be used to query the status of the request.
	Create(ctx context.Context, in *RsyncDataMovementCreateRequest, opts ...grpc.CallOption) (*RsyncDataMovementCreateResponse, error)
	// Status requests the status of a previously submitted data movement request. It accepts a unique
	// identifier that identifies the request and returns a status message.
	Status(ctx context.Context, in *RsyncDataMovementStatusRequest, opts ...grpc.CallOption) (*RsyncDataMovementStatusResponse, error)
	// Delete will attempt to delete a completed data movement request. It accepts a unique identifer
	// that identifies the request and returns the status of the delete operation.
	Delete(ctx context.Context, in *RsyncDataMovementDeleteRequest, opts ...grpc.CallOption) (*RsyncDataMovementDeleteResponse, error)
}

type rsyncDataMoverClient struct {
	cc grpc.ClientConnInterface
}

func NewRsyncDataMoverClient(cc grpc.ClientConnInterface) RsyncDataMoverClient {
	return &rsyncDataMoverClient{cc}
}

func (c *rsyncDataMoverClient) Create(ctx context.Context, in *RsyncDataMovementCreateRequest, opts ...grpc.CallOption) (*RsyncDataMovementCreateResponse, error) {
	out := new(RsyncDataMovementCreateResponse)
	err := c.cc.Invoke(ctx, "/datamovement.RsyncDataMover/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rsyncDataMoverClient) Status(ctx context.Context, in *RsyncDataMovementStatusRequest, opts ...grpc.CallOption) (*RsyncDataMovementStatusResponse, error) {
	out := new(RsyncDataMovementStatusResponse)
	err := c.cc.Invoke(ctx, "/datamovement.RsyncDataMover/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rsyncDataMoverClient) Delete(ctx context.Context, in *RsyncDataMovementDeleteRequest, opts ...grpc.CallOption) (*RsyncDataMovementDeleteResponse, error) {
	out := new(RsyncDataMovementDeleteResponse)
	err := c.cc.Invoke(ctx, "/datamovement.RsyncDataMover/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RsyncDataMoverServer is the server API for RsyncDataMover service.
// All implementations must embed UnimplementedRsyncDataMoverServer
// for forward compatibility
type RsyncDataMoverServer interface {
	// Create sends a new data movement request identified by source and destination fields. It returns
	// a response containing a unique identifier which can be used to query the status of the request.
	Create(context.Context, *RsyncDataMovementCreateRequest) (*RsyncDataMovementCreateResponse, error)
	// Status requests the status of a previously submitted data movement request. It accepts a unique
	// identifier that identifies the request and returns a status message.
	Status(context.Context, *RsyncDataMovementStatusRequest) (*RsyncDataMovementStatusResponse, error)
	// Delete will attempt to delete a completed data movement request. It accepts a unique identifer
	// that identifies the request and returns the status of the delete operation.
	Delete(context.Context, *RsyncDataMovementDeleteRequest) (*RsyncDataMovementDeleteResponse, error)
	mustEmbedUnimplementedRsyncDataMoverServer()
}

// UnimplementedRsyncDataMoverServer must be embedded to have forward compatible implementations.
type UnimplementedRsyncDataMoverServer struct {
}

func (UnimplementedRsyncDataMoverServer) Create(context.Context, *RsyncDataMovementCreateRequest) (*RsyncDataMovementCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedRsyncDataMoverServer) Status(context.Context, *RsyncDataMovementStatusRequest) (*RsyncDataMovementStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedRsyncDataMoverServer) Delete(context.Context, *RsyncDataMovementDeleteRequest) (*RsyncDataMovementDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedRsyncDataMoverServer) mustEmbedUnimplementedRsyncDataMoverServer() {}

// UnsafeRsyncDataMoverServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RsyncDataMoverServer will
// result in compilation errors.
type UnsafeRsyncDataMoverServer interface {
	mustEmbedUnimplementedRsyncDataMoverServer()
}

func RegisterRsyncDataMoverServer(s grpc.ServiceRegistrar, srv RsyncDataMoverServer) {
	s.RegisterService(&RsyncDataMover_ServiceDesc, srv)
}

func _RsyncDataMover_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RsyncDataMovementCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RsyncDataMoverServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datamovement.RsyncDataMover/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RsyncDataMoverServer).Create(ctx, req.(*RsyncDataMovementCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RsyncDataMover_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RsyncDataMovementStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RsyncDataMoverServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datamovement.RsyncDataMover/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RsyncDataMoverServer).Status(ctx, req.(*RsyncDataMovementStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RsyncDataMover_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RsyncDataMovementDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RsyncDataMoverServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datamovement.RsyncDataMover/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RsyncDataMoverServer).Delete(ctx, req.(*RsyncDataMovementDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RsyncDataMover_ServiceDesc is the grpc.ServiceDesc for RsyncDataMover service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RsyncDataMover_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "datamovement.RsyncDataMover",
	HandlerType: (*RsyncDataMoverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _RsyncDataMover_Create_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _RsyncDataMover_Status_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _RsyncDataMover_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/rsyncdatamovement.proto",
}