package service

import (
	"atlas-doc/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/transport/http"
	"io"

	"atlas-doc/api/doc/v1"
)

type FileService struct {
	v1.UnimplementedFileServer
	uc *biz.FileUseCase
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

func (s *FileService) UploadFile(ctx context.Context, req *v1.UploadFileReq) (*v1.UploadFileReply, error) {
	return s.uc.UploadFile(ctx, req)
}

func (s *FileService) Download(ctx context.Context, req *v1.IdReq) (*v1.DownloadFileReply, error) {
	return s.uc.Download(ctx, req)
}

func (s *FileService) GetFileInfo(ctx context.Context, req *v1.IdReq) (*v1.FileInfoReply, error) {
	return s.uc.GetFileInfo(ctx, req)
}
func (s *FileService) ListFile(ctx context.Context, req *v1.ListFileReq) (*v1.ListFileReply, error) {
	return s.uc.ListFile(ctx, req)
}
func (s *FileService) SaveFile(ctx context.Context, req *v1.EditFileReq) (*v1.EmptyReply, error) {
	return s.uc.SaveFile(ctx, req)
}
func (s *FileService) UpdateFile(ctx context.Context, req *v1.EditFileReq) (*v1.EmptyReply, error) {
	return s.uc.UpdateFile(ctx, req)
}
func (s *FileService) DeleteFile(ctx context.Context, req *v1.IdReq) (*v1.EmptyReply, error) {
	return s.uc.DeleteFile(ctx, req)
}
