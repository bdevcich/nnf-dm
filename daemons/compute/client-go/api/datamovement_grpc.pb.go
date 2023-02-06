// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: api/datamovement.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DataMoverClient is the client API for DataMover service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataMoverClient interface {
	// Version sends a request to the data movement daemon and returns a response containing
	// details on the current build version and supported API versions.
	Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*DataMovementVersionResponse, error)
	// Create sends a new data movement request identified by source and destination fields. It returns
	// a response containing a unique identifier which can be used to query the status of the request.
	Create(ctx context.Context, in *DataMovementCreateRequest, opts ...grpc.CallOption) (*DataMovementCreateResponse, error)
	// Status requests the status of a previously submitted data movement request. It accepts a unique
	// identifier that identifies the request and returns a status message.
	Status(ctx context.Context, in *DataMovementStatusRequest, opts ...grpc.CallOption) (*DataMovementStatusResponse, error)
	// Delete will attempt to delete a completed data movement request. It accepts a unique identifier
	// that identifies the request and returns the status of the delete operation.
	Delete(ctx context.Context, in *DataMovementDeleteRequest, opts ...grpc.CallOption) (*DataMovementDeleteResponse, error)
	// List returns all current data movement requests for a given namespace and workflow.
	List(ctx context.Context, in *DataMovementListRequest, opts ...grpc.CallOption) (*DataMovementListResponse, error)
	// Cancel will attempt to stop a data movement request. It accepts a unique identifier
	// that identifies the request and returns the status of the cancel operation.
	Cancel(ctx context.Context, in *DataMovementCancelRequest, opts ...grpc.CallOption) (*DataMovementCancelResponse, error)
}

type dataMoverClient struct {
	cc grpc.ClientConnInterface
}

func NewDataMoverClient(cc grpc.ClientConnInterface) DataMoverClient {
	return &dataMoverClient{cc}
}

func (c *dataMoverClient) Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*DataMovementVersionResponse, error) {
	out := new(DataMovementVersionResponse)
	err := c.cc.Invoke(ctx, "/datamovement.DataMover/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataMoverClient) Create(ctx context.Context, in *DataMovementCreateRequest, opts ...grpc.CallOption) (*DataMovementCreateResponse, error) {
	out := new(DataMovementCreateResponse)
	err := c.cc.Invoke(ctx, "/datamovement.DataMover/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataMoverClient) Status(ctx context.Context, in *DataMovementStatusRequest, opts ...grpc.CallOption) (*DataMovementStatusResponse, error) {
	out := new(DataMovementStatusResponse)
	err := c.cc.Invoke(ctx, "/datamovement.DataMover/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataMoverClient) Delete(ctx context.Context, in *DataMovementDeleteRequest, opts ...grpc.CallOption) (*DataMovementDeleteResponse, error) {
	out := new(DataMovementDeleteResponse)
	err := c.cc.Invoke(ctx, "/datamovement.DataMover/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataMoverClient) List(ctx context.Context, in *DataMovementListRequest, opts ...grpc.CallOption) (*DataMovementListResponse, error) {
	out := new(DataMovementListResponse)
	err := c.cc.Invoke(ctx, "/datamovement.DataMover/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataMoverClient) Cancel(ctx context.Context, in *DataMovementCancelRequest, opts ...grpc.CallOption) (*DataMovementCancelResponse, error) {
	out := new(DataMovementCancelResponse)
	err := c.cc.Invoke(ctx, "/datamovement.DataMover/Cancel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataMoverServer is the server API for DataMover service.
// All implementations must embed UnimplementedDataMoverServer
// for forward compatibility
type DataMoverServer interface {
	// Version sends a request to the data movement daemon and returns a response containing
	// details on the current build version and supported API versions.
	Version(context.Context, *emptypb.Empty) (*DataMovementVersionResponse, error)
	// Create sends a new data movement request identified by source and destination fields. It returns
	// a response containing a unique identifier which can be used to query the status of the request.
	Create(context.Context, *DataMovementCreateRequest) (*DataMovementCreateResponse, error)
	// Status requests the status of a previously submitted data movement request. It accepts a unique
	// identifier that identifies the request and returns a status message.
	Status(context.Context, *DataMovementStatusRequest) (*DataMovementStatusResponse, error)
	// Delete will attempt to delete a completed data movement request. It accepts a unique identifier
	// that identifies the request and returns the status of the delete operation.
	Delete(context.Context, *DataMovementDeleteRequest) (*DataMovementDeleteResponse, error)
	// List returns all current data movement requests for a given namespace and workflow.
	List(context.Context, *DataMovementListRequest) (*DataMovementListResponse, error)
	// Cancel will attempt to stop a data movement request. It accepts a unique identifier
	// that identifies the request and returns the status of the cancel operation.
	Cancel(context.Context, *DataMovementCancelRequest) (*DataMovementCancelResponse, error)
	mustEmbedUnimplementedDataMoverServer()
}

// UnimplementedDataMoverServer must be embedded to have forward compatible implementations.
type UnimplementedDataMoverServer struct {
}

func (UnimplementedDataMoverServer) Version(context.Context, *emptypb.Empty) (*DataMovementVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedDataMoverServer) Create(context.Context, *DataMovementCreateRequest) (*DataMovementCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedDataMoverServer) Status(context.Context, *DataMovementStatusRequest) (*DataMovementStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedDataMoverServer) Delete(context.Context, *DataMovementDeleteRequest) (*DataMovementDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedDataMoverServer) List(context.Context, *DataMovementListRequest) (*DataMovementListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedDataMoverServer) Cancel(context.Context, *DataMovementCancelRequest) (*DataMovementCancelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cancel not implemented")
}
func (UnimplementedDataMoverServer) mustEmbedUnimplementedDataMoverServer() {}

// UnsafeDataMoverServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataMoverServer will
// result in compilation errors.
type UnsafeDataMoverServer interface {
	mustEmbedUnimplementedDataMoverServer()
}

func RegisterDataMoverServer(s grpc.ServiceRegistrar, srv DataMoverServer) {
	s.RegisterService(&DataMover_ServiceDesc, srv)
}

func _DataMover_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataMoverServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datamovement.DataMover/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataMoverServer).Version(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataMover_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataMovementCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataMoverServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datamovement.DataMover/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataMoverServer).Create(ctx, req.(*DataMovementCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataMover_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataMovementStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataMoverServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datamovement.DataMover/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataMoverServer).Status(ctx, req.(*DataMovementStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataMover_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataMovementDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataMoverServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datamovement.DataMover/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataMoverServer).Delete(ctx, req.(*DataMovementDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataMover_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataMovementListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataMoverServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datamovement.DataMover/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataMoverServer).List(ctx, req.(*DataMovementListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataMover_Cancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataMovementCancelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataMoverServer).Cancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datamovement.DataMover/Cancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataMoverServer).Cancel(ctx, req.(*DataMovementCancelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DataMover_ServiceDesc is the grpc.ServiceDesc for DataMover service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataMover_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "datamovement.DataMover",
	HandlerType: (*DataMoverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _DataMover_Version_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _DataMover_Create_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _DataMover_Status_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _DataMover_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _DataMover_List_Handler,
		},
		{
			MethodName: "Cancel",
			Handler:    _DataMover_Cancel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/datamovement.proto",
}
