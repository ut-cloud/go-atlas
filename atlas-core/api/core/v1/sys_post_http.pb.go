// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.1
// - protoc             v5.28.2
// source: core/v1/sys_post.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationSysPostCreateSysPost = "/api.core.v1.SysPost/CreateSysPost"
const OperationSysPostDeleteSysPost = "/api.core.v1.SysPost/DeleteSysPost"
const OperationSysPostGetSysPost = "/api.core.v1.SysPost/GetSysPost"
const OperationSysPostListSysPost = "/api.core.v1.SysPost/ListSysPost"
const OperationSysPostUpdateSysPost = "/api.core.v1.SysPost/UpdateSysPost"

type SysPostHTTPServer interface {
	CreateSysPost(context.Context, *SysPostRep) (*EmptyReply, error)
	DeleteSysPost(context.Context, *DeleteSysPostRep) (*EmptyReply, error)
	GetSysPost(context.Context, *GetSysPostRep) (*GetSysPostReply, error)
	ListSysPost(context.Context, *ListSysPostRep) (*ListSysPostReply, error)
	UpdateSysPost(context.Context, *SysPostRep) (*EmptyReply, error)
}

func RegisterSysPostHTTPServer(s *http.Server, srv SysPostHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/post/save", _SysPost_CreateSysPost0_HTTP_Handler(srv))
	r.PUT("/v1/post/update", _SysPost_UpdateSysPost0_HTTP_Handler(srv))
	r.DELETE("/v1/post/delete/{id}", _SysPost_DeleteSysPost0_HTTP_Handler(srv))
	r.GET("/v1/post/info/{id}", _SysPost_GetSysPost0_HTTP_Handler(srv))
	r.POST("/v1/post/list", _SysPost_ListSysPost0_HTTP_Handler(srv))
}

func _SysPost_CreateSysPost0_HTTP_Handler(srv SysPostHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SysPostRep
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSysPostCreateSysPost)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateSysPost(ctx, req.(*SysPostRep))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*EmptyReply)
		return ctx.Result(200, reply)
	}
}

func _SysPost_UpdateSysPost0_HTTP_Handler(srv SysPostHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SysPostRep
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSysPostUpdateSysPost)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateSysPost(ctx, req.(*SysPostRep))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*EmptyReply)
		return ctx.Result(200, reply)
	}
}

func _SysPost_DeleteSysPost0_HTTP_Handler(srv SysPostHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteSysPostRep
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSysPostDeleteSysPost)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteSysPost(ctx, req.(*DeleteSysPostRep))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*EmptyReply)
		return ctx.Result(200, reply)
	}
}

func _SysPost_GetSysPost0_HTTP_Handler(srv SysPostHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetSysPostRep
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSysPostGetSysPost)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetSysPost(ctx, req.(*GetSysPostRep))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetSysPostReply)
		return ctx.Result(200, reply)
	}
}

func _SysPost_ListSysPost0_HTTP_Handler(srv SysPostHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListSysPostRep
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSysPostListSysPost)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListSysPost(ctx, req.(*ListSysPostRep))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListSysPostReply)
		return ctx.Result(200, reply)
	}
}

type SysPostHTTPClient interface {
	CreateSysPost(ctx context.Context, req *SysPostRep, opts ...http.CallOption) (rsp *EmptyReply, err error)
	DeleteSysPost(ctx context.Context, req *DeleteSysPostRep, opts ...http.CallOption) (rsp *EmptyReply, err error)
	GetSysPost(ctx context.Context, req *GetSysPostRep, opts ...http.CallOption) (rsp *GetSysPostReply, err error)
	ListSysPost(ctx context.Context, req *ListSysPostRep, opts ...http.CallOption) (rsp *ListSysPostReply, err error)
	UpdateSysPost(ctx context.Context, req *SysPostRep, opts ...http.CallOption) (rsp *EmptyReply, err error)
}

type SysPostHTTPClientImpl struct {
	cc *http.Client
}

func NewSysPostHTTPClient(client *http.Client) SysPostHTTPClient {
	return &SysPostHTTPClientImpl{client}
}

func (c *SysPostHTTPClientImpl) CreateSysPost(ctx context.Context, in *SysPostRep, opts ...http.CallOption) (*EmptyReply, error) {
	var out EmptyReply
	pattern := "/v1/post/save"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSysPostCreateSysPost))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SysPostHTTPClientImpl) DeleteSysPost(ctx context.Context, in *DeleteSysPostRep, opts ...http.CallOption) (*EmptyReply, error) {
	var out EmptyReply
	pattern := "/v1/post/delete/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSysPostDeleteSysPost))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SysPostHTTPClientImpl) GetSysPost(ctx context.Context, in *GetSysPostRep, opts ...http.CallOption) (*GetSysPostReply, error) {
	var out GetSysPostReply
	pattern := "/v1/post/info/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSysPostGetSysPost))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SysPostHTTPClientImpl) ListSysPost(ctx context.Context, in *ListSysPostRep, opts ...http.CallOption) (*ListSysPostReply, error) {
	var out ListSysPostReply
	pattern := "/v1/post/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSysPostListSysPost))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SysPostHTTPClientImpl) UpdateSysPost(ctx context.Context, in *SysPostRep, opts ...http.CallOption) (*EmptyReply, error) {
	var out EmptyReply
	pattern := "/v1/post/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSysPostUpdateSysPost))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
