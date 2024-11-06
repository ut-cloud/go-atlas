package biz

import (
	v1 "atlas-doc/api/doc/v1"
	"atlas-doc/internal/conf"
	"atlas-doc/internal/constants"
	"atlas-doc/internal/model"
	"bytes"
	"context"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/ut-cloud/atlas-toolkit/utils"
	"path/filepath"
	"strings"
	"time"
)

type FileRepo interface {
	GetInfo(ctx context.Context, id string) (*model.BizFile, error)
	Save(ctx context.Context, file *model.File) error
	Update(ctx context.Context, file *model.File) error
	Delete(ctx context.Context, fileId string) error
	GetPageSet(ctx context.Context, pageNum int64, pageSize int64, query model.File) ([]*model.BizFile, int64, error)
}

type FileUseCase struct {
	cnf  *conf.File
	repo FileRepo
	log  *log.Helper
}

func NewFileUseCase(repo FileRepo, cnf *conf.File, logger log.Logger) *FileUseCase {
	return &FileUseCase{repo: repo, cnf: cnf, log: log.NewHelper(logger)}
}

func (uc *FileUseCase) UploadFile(ctx context.Context, req *v1.UploadFileReq) (*v1.UploadFileReply, error) {
	fileExtension := filepath.Ext(req.GetName())
	id := utils.GetID()
	refName := id + req.GetName()
	keyPrefix := constants.FilePrefix + time.Now().Format("20060102")
	size := req.GetFileSize()
	var sizeStr string
	// 根据大小选择转换单位
	if size >= 1024*1024*1024 {
		sizeStr = fmt.Sprintf("%.2f GB", float64(size)/(1024*1024*1024)) // 转换为 GB
	} else if size >= 1024*1024 {
		sizeStr = fmt.Sprintf("%.2f MB", float64(size)/(1024*1024)) // 转换为 MB
	} else if size >= 1024 {
		sizeStr = fmt.Sprintf("%.2f KB", float64(size)/1024) // 转换为 KB
	} else {
		sizeStr = fmt.Sprintf("%d B", size) // 小于 1 KB 显示为字节
	}
	var keyUrl string
	uploadType := uc.cnf.Type
	if uploadType == "aliyun" {
		var err error
		keyUrl, err = uc.AliyunOssUpload(keyPrefix, refName, req.File)
		if err != nil {
			return nil, err
		}
	}
	if uploadType == "local" {

	}
	file := &model.File{
		UUID:          id,
		FileName:      req.GetName(),
		RefName:       refName,
		FileExtension: fileExtension,
		KeyURL:        keyUrl,
		KeyPrefix:     keyPrefix,
		Endpoint:      uc.cnf.Oss.Endpoint,
		Size:          req.GetFileSize(),
		FileSize:      sizeStr,
		Type:          uploadType,
	}
	file.InsertEntity(ctx)
	if err := uc.repo.Save(ctx, file); err != nil {
		return nil, err
	}
	return &v1.UploadFileReply{FileId: id}, nil
}

// AliyunOssUpload
//
//	@Description: 阿里云文件上传
//	@receiver s
//	@param keyPrefix 上传的文件路径
//	@param fileName 文件名
//	@param data 文件流
func (uc *FileUseCase) AliyunOssUpload(keyPrefix, refName string, data []byte) (string, error) {
	err, bucket, objectKey := uc.AliyunClient(keyPrefix, refName)
	err = bucket.PutObject(objectKey, bytes.NewReader(data))
	if err != nil {
		uc.log.Errorf("Failed to put object from file: %v", err)
		return "", err
	}
	uc.log.Info("File uploaded successfully.")
	return objectKey, nil
}

func (uc *FileUseCase) AliyunOssDownload(keyPrefix, refName string) (string, error) {
	err, bucket, objectKey := uc.AliyunClient(keyPrefix, refName)
	// 生成用于下载的签名URL，并指定签名URL的有效时间为1小时。
	signedURL, err := bucket.SignURL(objectKey, oss.HTTPGet, 60*60)
	if err != nil {
		uc.log.Errorf("Failed to sign object from file: %v", err)
	}
	uc.log.Info("Sign Url:%s", signedURL)
	return signedURL, nil
}

func (uc *FileUseCase) AliyunClient(keyPrefix string, refName string) (error, *oss.Bucket, string) {
	if !strings.HasSuffix(keyPrefix, "/") {
		keyPrefix = keyPrefix + "/"
	}
	provider, err := NewDefaultCredentialsProvider(uc.cnf.Oss.AccessKey, uc.cnf.Oss.AccessKeySecret, "")
	if err != nil {
		uc.log.Errorf("NewDefaultCredentialsProvider failed: %v", err)
		return err, nil, ""
	}
	clientOptions := []oss.ClientOption{oss.SetCredentialsProvider(&provider)}
	clientOptions = append(clientOptions, oss.Region(strings.TrimPrefix(strings.Split(uc.cnf.Oss.Endpoint, ".")[0], "oss-")))
	clientOptions = append(clientOptions, oss.AuthVersion(oss.AuthV4))
	client, err := oss.New("https://"+uc.cnf.Oss.Endpoint, "", "", clientOptions...)
	if err != nil {
		uc.log.Errorf("oss new client failed: %v", err)
		return err, nil, ""
	}
	bucket, err := client.Bucket(uc.cnf.Oss.BucketName)
	if err != nil {
		uc.log.Errorf("Failed to get bucket: %v", err)
		return err, nil, ""
	}
	objectKey := keyPrefix + refName
	return err, bucket, objectKey
}

func (uc *FileUseCase) Download(ctx context.Context, req *v1.IdReq) (*v1.DownloadFileReply, error) {
	fileInfo, err := uc.repo.GetInfo(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	if fileInfo.Type == constants.Aliyun {
		if fileUrl, err := uc.AliyunOssDownload(fileInfo.KeyPrefix, fileInfo.RefName); err != nil {
			return nil, err
		} else {
			return &v1.DownloadFileReply{
				Url: fileUrl,
			}, nil
		}
	}
	return &v1.DownloadFileReply{}, nil
}

func (uc *FileUseCase) GetFileInfo(ctx context.Context, req *v1.IdReq) (*v1.FileInfoReply, error) {
	return &v1.FileInfoReply{}, nil
}

func (uc *FileUseCase) ListFile(ctx context.Context, req *v1.ListFileReq) (*v1.ListFileReply, error) {
	return &v1.ListFileReply{}, nil
}

func (uc *FileUseCase) SaveFile(ctx context.Context, req *v1.EditFileReq) (*v1.EmptyReply, error) {
	return &v1.EmptyReply{}, nil
}

func (uc *FileUseCase) UpdateFile(ctx context.Context, req *v1.EditFileReq) (*v1.EmptyReply, error) {
	return &v1.EmptyReply{}, nil
}

func (uc *FileUseCase) DeleteFile(ctx context.Context, req *v1.IdReq) (*v1.EmptyReply, error) {
	return &v1.EmptyReply{}, nil
}

type defaultCredentials struct {
	config *oss.Config
}

func (defCre *defaultCredentials) GetAccessKeyID() string {
	return defCre.config.AccessKeyID
}

func (defCre *defaultCredentials) GetAccessKeySecret() string {
	return defCre.config.AccessKeySecret
}

func (defCre *defaultCredentials) GetSecurityToken() string {
	return defCre.config.SecurityToken
}

type DefaultCredentialsProvider struct {
	config *oss.Config
}

func (defBuild *DefaultCredentialsProvider) GetCredentials() oss.Credentials {
	return &defaultCredentials{config: defBuild.config}
}
func NewDefaultCredentialsProvider(accessID, accessKey, token string) (DefaultCredentialsProvider, error) {
	var provider DefaultCredentialsProvider
	if accessID == "" {
		return provider, fmt.Errorf("access key id is empty")
	}
	if accessKey == "" {
		return provider, fmt.Errorf("access key secret is empty")
	}
	config := &oss.Config{
		AccessKeyID:     accessID,
		AccessKeySecret: accessKey,
		SecurityToken:   token,
	}
	return DefaultCredentialsProvider{
		config,
	}, nil
}
