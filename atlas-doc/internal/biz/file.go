package biz

import (
	v1 "atlas-doc/api/doc/v1"
	"atlas-doc/internal/conf"
	"bytes"
	"context"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/ut-cloud/atlas-toolkit/utils"
	"strings"
)

type FileRepo interface {
}

type FileUseCase struct {
	cnf  *conf.Oss
	repo FileRepo
	log  *log.Helper
}

func NewFileUseCase(repo FileRepo, cnf *conf.Oss, logger log.Logger) *FileUseCase {
	return &FileUseCase{repo: repo, cnf: cnf, log: log.NewHelper(logger)}
}

func (uc *FileUseCase) UploadFile(ctx context.Context, req *v1.UploadFileReq) (*v1.UploadFileReply, error) {
	uuid := utils.GetID()
	uc.AliyunOssUpload(uuid, "doc", req.GetName(), req.File)
	return &v1.UploadFileReply{}, nil
}

// AliyunOssUpload
//
//	@Description: 阿里云文件上传
//	@receiver s
//	@param keyPrefix 上传的文件路径
//	@param fileName 文件名
//	@param data 文件流
func (uc *FileUseCase) AliyunOssUpload(uuid, keyPrefix, fileName string, data []byte) {
	if !strings.HasSuffix(keyPrefix, "/") {
		keyPrefix = keyPrefix + "/"
	}
	provider, err := NewDefaultCredentialsProvider(uc.cnf.AccessKey, uc.cnf.AccessKeySecret, "")
	if err != nil {
		log.Fatalf("NewDefaultCredentialsProvider err: %v", err)
	}
	clientOptions := []oss.ClientOption{oss.SetCredentialsProvider(&provider)}
	clientOptions = append(clientOptions, oss.Region(strings.TrimPrefix(strings.Split(uc.cnf.Endpoint, ".")[0], "oss-")))
	clientOptions = append(clientOptions, oss.AuthVersion(oss.AuthV4))
	client, err := oss.New("https://"+uc.cnf.Endpoint, "", "", clientOptions...)
	if err != nil {
		log.Fatalf("Failed to create OSS client: %v", err)
	}
	bucket, err := client.Bucket(uc.cnf.BucketName)
	if err != nil {
		log.Fatalf("Failed to get bucket: %v", err)
	}

	objectKey := keyPrefix + uuid + "/" + fileName
	//localFilePath := "/Users/ut/code/go/github/go-atlas/atlas-doc/Makefile"
	err = bucket.PutObject(objectKey, bytes.NewReader(data))
	//err = bucket.PutObjectFromFile(objectKey, localFilePath)
	if err != nil {
		log.Fatalf("Failed to put object from file: %v", err)
	}
	uc.log.Info("File uploaded successfully.")
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
