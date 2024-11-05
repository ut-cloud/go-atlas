package service

import (
	v1 "atlas-doc/api/doc/v1"
	"atlas-doc/internal/biz"
	"context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	"io"
)

type FileService struct {
	uc *biz.FileUseCase
	v1.UnimplementedFileServer
}

func NewFileService(uc *biz.FileUseCase) *FileService {
	return &FileService{uc: uc}
}

// Upload
//
//	@Description: 文件上传 http调用
//	@receiver s
//	@return http.HandlerFunc
func (s *FileService) Upload() http.HandlerFunc {
	return func(ctx http.Context) error {
		var in v1.UploadFileReq
		file, header, err := ctx.Request().FormFile("file")
		if err != nil {
			return err
		}
		in.File, err = io.ReadAll(file)
		if err != nil {
			return err
		}
		in.FileSize = int32(header.Size)
		in.Name = header.Filename

		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return s.UploadFile(ctx, req.(*v1.UploadFileReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.UploadFileReply)
		return ctx.Result(200, reply)
	}
}

// UploadFile
//
//	@Description: 文件上传 grpc调用
//	@receiver s
//	@param ctx
//	@param req
//	@return *v1.UploadFileReply
//	@return error
func (s *FileService) UploadFile(ctx context.Context, req *v1.UploadFileReq) (*v1.UploadFileReply, error) {
	return s.uc.UploadFile(ctx, req)
}
