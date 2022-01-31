// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: api/goods/v1/goods.proto

package v1

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

// GoodsServiceClient is the client API for GoodsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GoodsServiceClient interface {
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginReply, error)
	CreateGoods(ctx context.Context, in *CreateGoodsRequest, opts ...grpc.CallOption) (*CreateGoodsReply, error)
	UpdateGoods(ctx context.Context, in *UpdateGoodsRequest, opts ...grpc.CallOption) (*UpdateGoodsReply, error)
	DeleteGoods(ctx context.Context, in *DeleteGoodsRequest, opts ...grpc.CallOption) (*DeleteGoodsReply, error)
	GetGoods(ctx context.Context, in *GetGoodsRequest, opts ...grpc.CallOption) (*GetGoodsReply, error)
	ListGoods(ctx context.Context, in *ListGoodsRequest, opts ...grpc.CallOption) (*ListGoodsReply, error)
}

type goodsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGoodsServiceClient(cc grpc.ClientConnInterface) GoodsServiceClient {
	return &goodsServiceClient{cc}
}

func (c *goodsServiceClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, "/goods.v1.GoodsService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsServiceClient) CreateGoods(ctx context.Context, in *CreateGoodsRequest, opts ...grpc.CallOption) (*CreateGoodsReply, error) {
	out := new(CreateGoodsReply)
	err := c.cc.Invoke(ctx, "/goods.v1.GoodsService/CreateGoods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsServiceClient) UpdateGoods(ctx context.Context, in *UpdateGoodsRequest, opts ...grpc.CallOption) (*UpdateGoodsReply, error) {
	out := new(UpdateGoodsReply)
	err := c.cc.Invoke(ctx, "/goods.v1.GoodsService/UpdateGoods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsServiceClient) DeleteGoods(ctx context.Context, in *DeleteGoodsRequest, opts ...grpc.CallOption) (*DeleteGoodsReply, error) {
	out := new(DeleteGoodsReply)
	err := c.cc.Invoke(ctx, "/goods.v1.GoodsService/DeleteGoods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsServiceClient) GetGoods(ctx context.Context, in *GetGoodsRequest, opts ...grpc.CallOption) (*GetGoodsReply, error) {
	out := new(GetGoodsReply)
	err := c.cc.Invoke(ctx, "/goods.v1.GoodsService/GetGoods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsServiceClient) ListGoods(ctx context.Context, in *ListGoodsRequest, opts ...grpc.CallOption) (*ListGoodsReply, error) {
	out := new(ListGoodsReply)
	err := c.cc.Invoke(ctx, "/goods.v1.GoodsService/ListGoods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoodsServiceServer is the server API for GoodsService service.
// All implementations must embed UnimplementedGoodsServiceServer
// for forward compatibility
type GoodsServiceServer interface {
	Login(context.Context, *LoginReq) (*LoginReply, error)
	CreateGoods(context.Context, *CreateGoodsRequest) (*CreateGoodsReply, error)
	UpdateGoods(context.Context, *UpdateGoodsRequest) (*UpdateGoodsReply, error)
	DeleteGoods(context.Context, *DeleteGoodsRequest) (*DeleteGoodsReply, error)
	GetGoods(context.Context, *GetGoodsRequest) (*GetGoodsReply, error)
	ListGoods(context.Context, *ListGoodsRequest) (*ListGoodsReply, error)
	mustEmbedUnimplementedGoodsServiceServer()
}

// UnimplementedGoodsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGoodsServiceServer struct {
}

func (UnimplementedGoodsServiceServer) Login(context.Context, *LoginReq) (*LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedGoodsServiceServer) CreateGoods(context.Context, *CreateGoodsRequest) (*CreateGoodsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGoods not implemented")
}
func (UnimplementedGoodsServiceServer) UpdateGoods(context.Context, *UpdateGoodsRequest) (*UpdateGoodsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGoods not implemented")
}
func (UnimplementedGoodsServiceServer) DeleteGoods(context.Context, *DeleteGoodsRequest) (*DeleteGoodsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGoods not implemented")
}
func (UnimplementedGoodsServiceServer) GetGoods(context.Context, *GetGoodsRequest) (*GetGoodsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoods not implemented")
}
func (UnimplementedGoodsServiceServer) ListGoods(context.Context, *ListGoodsRequest) (*ListGoodsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGoods not implemented")
}
func (UnimplementedGoodsServiceServer) mustEmbedUnimplementedGoodsServiceServer() {}

// UnsafeGoodsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GoodsServiceServer will
// result in compilation errors.
type UnsafeGoodsServiceServer interface {
	mustEmbedUnimplementedGoodsServiceServer()
}

func RegisterGoodsServiceServer(s grpc.ServiceRegistrar, srv GoodsServiceServer) {
	s.RegisterService(&GoodsService_ServiceDesc, srv)
}

func _GoodsService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goods.v1.GoodsService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServiceServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoodsService_CreateGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGoodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServiceServer).CreateGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goods.v1.GoodsService/CreateGoods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServiceServer).CreateGoods(ctx, req.(*CreateGoodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoodsService_UpdateGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGoodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServiceServer).UpdateGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goods.v1.GoodsService/UpdateGoods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServiceServer).UpdateGoods(ctx, req.(*UpdateGoodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoodsService_DeleteGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGoodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServiceServer).DeleteGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goods.v1.GoodsService/DeleteGoods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServiceServer).DeleteGoods(ctx, req.(*DeleteGoodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoodsService_GetGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGoodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServiceServer).GetGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goods.v1.GoodsService/GetGoods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServiceServer).GetGoods(ctx, req.(*GetGoodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoodsService_ListGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGoodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServiceServer).ListGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goods.v1.GoodsService/ListGoods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServiceServer).ListGoods(ctx, req.(*ListGoodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GoodsService_ServiceDesc is the grpc.ServiceDesc for GoodsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GoodsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "goods.v1.GoodsService",
	HandlerType: (*GoodsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _GoodsService_Login_Handler,
		},
		{
			MethodName: "CreateGoods",
			Handler:    _GoodsService_CreateGoods_Handler,
		},
		{
			MethodName: "UpdateGoods",
			Handler:    _GoodsService_UpdateGoods_Handler,
		},
		{
			MethodName: "DeleteGoods",
			Handler:    _GoodsService_DeleteGoods_Handler,
		},
		{
			MethodName: "GetGoods",
			Handler:    _GoodsService_GetGoods_Handler,
		},
		{
			MethodName: "ListGoods",
			Handler:    _GoodsService_ListGoods_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/goods/v1/goods.proto",
}
