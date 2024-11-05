// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.1
// - protoc             v5.28.2
// source: doc/v1/file.proto

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

const OperationFileDeleteFile = "/api.doc.v1.File/DeleteFile"
const OperationFileGetFileInfo = "/api.doc.v1.File/GetFileInfo"
const OperationFileListFile = "/api.doc.v1.File/ListFile"
const OperationFileSaveFile = "/api.doc.v1.File/SaveFile"
const OperationFileUpdateFile = "/api.doc.v1.File/UpdateFile"

type FileHTTPServer interface {
	DeleteFile(context.Context, *IdReq) (*EmptyReply, error)
	GetFileInfo(context.Context, *IdReq) (*FileInfoReply, error)
	ListFile(context.Context, *ListFileReq) (*ListFileReply, error)
	SaveFile(context.Context, *EditFileReq) (*EmptyReply, error)
	UpdateFile(context.Context, *EditFileReq) (*EmptyReply, error)
}

func RegisterFileHTTPServer(s *http.Server, srv FileHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/file/info/{id}", _File_GetFileInfo0_HTTP_Handler(srv))
	r.POST("/v1/file/list", _File_ListFile0_HTTP_Handler(srv))
	r.POST("/v1/file/save", _File_SaveFile0_HTTP_Handler(srv))
	r.PUT("/v1/file/update", _File_UpdateFile0_HTTP_Handler(srv))
	r.DELETE("/v1/file/delete/{id}", _File_DeleteFile0_HTTP_Handler(srv))
}

func _File_GetFileInfo0_HTTP_Handler(srv FileHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in IdReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationFileGetFileInfo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetFileInfo(ctx, req.(*IdReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*FileInfoReply)
		return ctx.Result(200, reply)
	}
}

func _File_ListFile0_HTTP_Handler(srv FileHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListFileReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationFileListFile)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListFile(ctx, req.(*ListFileReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListFileReply)
		return ctx.Result(200, reply)
	}
}

func _File_SaveFile0_HTTP_Handler(srv FileHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in EditFileReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationFileSaveFile)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SaveFile(ctx, req.(*EditFileReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*EmptyReply)
		return ctx.Result(200, reply)
	}
}

func _File_UpdateFile0_HTTP_Handler(srv FileHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in EditFileReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationFileUpdateFile)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateFile(ctx, req.(*EditFileReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*EmptyReply)
		return ctx.Result(200, reply)
	}
}

func _File_DeleteFile0_HTTP_Handler(srv FileHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in IdReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationFileDeleteFile)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteFile(ctx, req.(*IdReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*EmptyReply)
		return ctx.Result(200, reply)
	}
}

type FileHTTPClient interface {
	DeleteFile(ctx context.Context, req *IdReq, opts ...http.CallOption) (rsp *EmptyReply, err error)
	GetFileInfo(ctx context.Context, req *IdReq, opts ...http.CallOption) (rsp *FileInfoReply, err error)
	ListFile(ctx context.Context, req *ListFileReq, opts ...http.CallOption) (rsp *ListFileReply, err error)
	SaveFile(ctx context.Context, req *EditFileReq, opts ...http.CallOption) (rsp *EmptyReply, err error)
	UpdateFile(ctx context.Context, req *EditFileReq, opts ...http.CallOption) (rsp *EmptyReply, err error)
}

type FileHTTPClientImpl struct {
	cc *http.Client
}

func NewFileHTTPClient(client *http.Client) FileHTTPClient {
	return &FileHTTPClientImpl{client}
}

func (c *FileHTTPClientImpl) DeleteFile(ctx context.Context, in *IdReq, opts ...http.CallOption) (*EmptyReply, error) {
	var out EmptyReply
	pattern := "/v1/file/delete/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationFileDeleteFile))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *FileHTTPClientImpl) GetFileInfo(ctx context.Context, in *IdReq, opts ...http.CallOption) (*FileInfoReply, error) {
	var out FileInfoReply
	pattern := "/v1/file/info/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationFileGetFileInfo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *FileHTTPClientImpl) ListFile(ctx context.Context, in *ListFileReq, opts ...http.CallOption) (*ListFileReply, error) {
	var out ListFileReply
	pattern := "/v1/file/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationFileListFile))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *FileHTTPClientImpl) SaveFile(ctx context.Context, in *EditFileReq, opts ...http.CallOption) (*EmptyReply, error) {
	var out EmptyReply
	pattern := "/v1/file/save"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationFileSaveFile))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *FileHTTPClientImpl) UpdateFile(ctx context.Context, in *EditFileReq, opts ...http.CallOption) (*EmptyReply, error) {
	var out EmptyReply
	pattern := "/v1/file/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationFileUpdateFile))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
