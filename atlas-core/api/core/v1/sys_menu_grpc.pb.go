// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: core/v1/sys_menu.proto

package v1

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
	SysMenu_CreateSysMenu_FullMethodName  = "/api.core.v1.SysMenu/CreateSysMenu"
	SysMenu_UpdateSysMenu_FullMethodName  = "/api.core.v1.SysMenu/UpdateSysMenu"
	SysMenu_DeleteSysMenu_FullMethodName  = "/api.core.v1.SysMenu/DeleteSysMenu"
	SysMenu_GetSysMenu_FullMethodName     = "/api.core.v1.SysMenu/GetSysMenu"
	SysMenu_ListSysMenu_FullMethodName    = "/api.core.v1.SysMenu/ListSysMenu"
	SysMenu_GetSysRoleMenu_FullMethodName = "/api.core.v1.SysMenu/GetSysRoleMenu"
	SysMenu_GetTreeSelect_FullMethodName  = "/api.core.v1.SysMenu/GetTreeSelect"
)

// SysMenuClient is the client API for SysMenu service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SysMenuClient interface {
	CreateSysMenu(ctx context.Context, in *SysMenuRep, opts ...grpc.CallOption) (*EmptyReply, error)
	UpdateSysMenu(ctx context.Context, in *SysMenuRep, opts ...grpc.CallOption) (*EmptyReply, error)
	DeleteSysMenu(ctx context.Context, in *DeleteSysMenuRep, opts ...grpc.CallOption) (*EmptyReply, error)
	GetSysMenu(ctx context.Context, in *GetSysMenuRep, opts ...grpc.CallOption) (*GetSysMenuReply, error)
	ListSysMenu(ctx context.Context, in *ListSysMenuRep, opts ...grpc.CallOption) (*ListSysMenuReply, error)
	GetSysRoleMenu(ctx context.Context, in *GetSysRoleMenuRep, opts ...grpc.CallOption) (*GetSysRoleMenuReply, error)
	GetTreeSelect(ctx context.Context, in *GetTreeSelectRep, opts ...grpc.CallOption) (*GetTreeSelectReply, error)
}

type sysMenuClient struct {
	cc grpc.ClientConnInterface
}

func NewSysMenuClient(cc grpc.ClientConnInterface) SysMenuClient {
	return &sysMenuClient{cc}
}

func (c *sysMenuClient) CreateSysMenu(ctx context.Context, in *SysMenuRep, opts ...grpc.CallOption) (*EmptyReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EmptyReply)
	err := c.cc.Invoke(ctx, SysMenu_CreateSysMenu_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysMenuClient) UpdateSysMenu(ctx context.Context, in *SysMenuRep, opts ...grpc.CallOption) (*EmptyReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EmptyReply)
	err := c.cc.Invoke(ctx, SysMenu_UpdateSysMenu_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysMenuClient) DeleteSysMenu(ctx context.Context, in *DeleteSysMenuRep, opts ...grpc.CallOption) (*EmptyReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EmptyReply)
	err := c.cc.Invoke(ctx, SysMenu_DeleteSysMenu_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysMenuClient) GetSysMenu(ctx context.Context, in *GetSysMenuRep, opts ...grpc.CallOption) (*GetSysMenuReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSysMenuReply)
	err := c.cc.Invoke(ctx, SysMenu_GetSysMenu_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysMenuClient) ListSysMenu(ctx context.Context, in *ListSysMenuRep, opts ...grpc.CallOption) (*ListSysMenuReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListSysMenuReply)
	err := c.cc.Invoke(ctx, SysMenu_ListSysMenu_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysMenuClient) GetSysRoleMenu(ctx context.Context, in *GetSysRoleMenuRep, opts ...grpc.CallOption) (*GetSysRoleMenuReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSysRoleMenuReply)
	err := c.cc.Invoke(ctx, SysMenu_GetSysRoleMenu_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysMenuClient) GetTreeSelect(ctx context.Context, in *GetTreeSelectRep, opts ...grpc.CallOption) (*GetTreeSelectReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTreeSelectReply)
	err := c.cc.Invoke(ctx, SysMenu_GetTreeSelect_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SysMenuServer is the server API for SysMenu service.
// All implementations must embed UnimplementedSysMenuServer
// for forward compatibility.
type SysMenuServer interface {
	CreateSysMenu(context.Context, *SysMenuRep) (*EmptyReply, error)
	UpdateSysMenu(context.Context, *SysMenuRep) (*EmptyReply, error)
	DeleteSysMenu(context.Context, *DeleteSysMenuRep) (*EmptyReply, error)
	GetSysMenu(context.Context, *GetSysMenuRep) (*GetSysMenuReply, error)
	ListSysMenu(context.Context, *ListSysMenuRep) (*ListSysMenuReply, error)
	GetSysRoleMenu(context.Context, *GetSysRoleMenuRep) (*GetSysRoleMenuReply, error)
	GetTreeSelect(context.Context, *GetTreeSelectRep) (*GetTreeSelectReply, error)
	mustEmbedUnimplementedSysMenuServer()
}

// UnimplementedSysMenuServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSysMenuServer struct{}

func (UnimplementedSysMenuServer) CreateSysMenu(context.Context, *SysMenuRep) (*EmptyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSysMenu not implemented")
}
func (UnimplementedSysMenuServer) UpdateSysMenu(context.Context, *SysMenuRep) (*EmptyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSysMenu not implemented")
}
func (UnimplementedSysMenuServer) DeleteSysMenu(context.Context, *DeleteSysMenuRep) (*EmptyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSysMenu not implemented")
}
func (UnimplementedSysMenuServer) GetSysMenu(context.Context, *GetSysMenuRep) (*GetSysMenuReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSysMenu not implemented")
}
func (UnimplementedSysMenuServer) ListSysMenu(context.Context, *ListSysMenuRep) (*ListSysMenuReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSysMenu not implemented")
}
func (UnimplementedSysMenuServer) GetSysRoleMenu(context.Context, *GetSysRoleMenuRep) (*GetSysRoleMenuReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSysRoleMenu not implemented")
}
func (UnimplementedSysMenuServer) GetTreeSelect(context.Context, *GetTreeSelectRep) (*GetTreeSelectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTreeSelect not implemented")
}
func (UnimplementedSysMenuServer) mustEmbedUnimplementedSysMenuServer() {}
func (UnimplementedSysMenuServer) testEmbeddedByValue()                 {}

// UnsafeSysMenuServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SysMenuServer will
// result in compilation errors.
type UnsafeSysMenuServer interface {
	mustEmbedUnimplementedSysMenuServer()
}

func RegisterSysMenuServer(s grpc.ServiceRegistrar, srv SysMenuServer) {
	// If the following call pancis, it indicates UnimplementedSysMenuServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SysMenu_ServiceDesc, srv)
}

func _SysMenu_CreateSysMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysMenuRep)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysMenuServer).CreateSysMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysMenu_CreateSysMenu_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysMenuServer).CreateSysMenu(ctx, req.(*SysMenuRep))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysMenu_UpdateSysMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysMenuRep)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysMenuServer).UpdateSysMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysMenu_UpdateSysMenu_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysMenuServer).UpdateSysMenu(ctx, req.(*SysMenuRep))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysMenu_DeleteSysMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSysMenuRep)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysMenuServer).DeleteSysMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysMenu_DeleteSysMenu_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysMenuServer).DeleteSysMenu(ctx, req.(*DeleteSysMenuRep))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysMenu_GetSysMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSysMenuRep)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysMenuServer).GetSysMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysMenu_GetSysMenu_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysMenuServer).GetSysMenu(ctx, req.(*GetSysMenuRep))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysMenu_ListSysMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSysMenuRep)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysMenuServer).ListSysMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysMenu_ListSysMenu_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysMenuServer).ListSysMenu(ctx, req.(*ListSysMenuRep))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysMenu_GetSysRoleMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSysRoleMenuRep)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysMenuServer).GetSysRoleMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysMenu_GetSysRoleMenu_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysMenuServer).GetSysRoleMenu(ctx, req.(*GetSysRoleMenuRep))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysMenu_GetTreeSelect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTreeSelectRep)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysMenuServer).GetTreeSelect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysMenu_GetTreeSelect_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysMenuServer).GetTreeSelect(ctx, req.(*GetTreeSelectRep))
	}
	return interceptor(ctx, in, info, handler)
}

// SysMenu_ServiceDesc is the grpc.ServiceDesc for SysMenu service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SysMenu_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.core.v1.SysMenu",
	HandlerType: (*SysMenuServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSysMenu",
			Handler:    _SysMenu_CreateSysMenu_Handler,
		},
		{
			MethodName: "UpdateSysMenu",
			Handler:    _SysMenu_UpdateSysMenu_Handler,
		},
		{
			MethodName: "DeleteSysMenu",
			Handler:    _SysMenu_DeleteSysMenu_Handler,
		},
		{
			MethodName: "GetSysMenu",
			Handler:    _SysMenu_GetSysMenu_Handler,
		},
		{
			MethodName: "ListSysMenu",
			Handler:    _SysMenu_ListSysMenu_Handler,
		},
		{
			MethodName: "GetSysRoleMenu",
			Handler:    _SysMenu_GetSysRoleMenu_Handler,
		},
		{
			MethodName: "GetTreeSelect",
			Handler:    _SysMenu_GetTreeSelect_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "core/v1/sys_menu.proto",
}
