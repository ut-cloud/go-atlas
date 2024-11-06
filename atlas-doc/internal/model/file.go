package model

import (
	v1 "atlas-doc/api/doc/v1"
	"github.com/ut-cloud/atlas-toolkit/base"
	"time"
)

// File represents the OSS file table
type File struct {
	UUID          string `gorm:"type:varchar(32);primaryKey;not null" json:"uuid"`     // 主键
	FileName      string `gorm:"type:varchar(255);default:null" json:"file_name"`      // 文件名称
	RefName       string `gorm:"type:varchar(255);default:null" json:"ref_name"`       // 文件名称
	FileExtension string `gorm:"type:varchar(255);default:null" json:"file_extension"` // 文件名称
	KeyURL        string `gorm:"type:varchar(255);default:null" json:"key_url"`        // 地址
	BucketName    string `gorm:"type:varchar(255);default:null" json:"bucket_name"`    // oss桶名
	Endpoint      string `gorm:"type:varchar(255);default:null" json:"endpoint"`       // oss端点
	KeyPrefix     string `gorm:"type:varchar(255);default:null" json:"key_prefix"`     // 文件地址目录
	Size          int32  `gorm:"type:int;default:null" json:"size"`                    // 大小
	FileSize      string `gorm:"type:varchar(16);default:null" json:"file_size"`       // 文件大小
	Type          string `gorm:"type:varchar(255);default:null" json:"type"`           // 文件大小
	base.Entity
}

func (*File) TableName() string {
	return "doc_file"
}

type BizFile struct {
	UUID          string
	FileName      string
	RefName       string
	FileExtension string
	KeyURL        string
	BucketName    string
	Endpoint      string
	KeyPrefix     string
	Size          int32
	FileSize      string
	Type          string
	CreatedAt     time.Time
}

func (file *File) EntityToBiz() *BizFile {
	return &BizFile{
		UUID:          file.UUID,
		FileName:      file.FileName,
		RefName:       file.RefName,
		FileExtension: file.FileExtension,
		KeyURL:        file.KeyURL,
		BucketName:    file.BucketName,
		Endpoint:      file.Endpoint,
		KeyPrefix:     file.KeyPrefix,
		Size:          file.Size,
		FileSize:      file.FileSize,
		Type:          file.Type,
		CreatedAt:     file.CreatedAt,
	}
}

func (file *BizFile) BizToV1() *v1.FileInfoReply {
	return &v1.FileInfoReply{
		Uuid:       file.UUID,
		FileName:   file.FileName,
		KeyUrl:     file.KeyURL,
		BucketName: file.BucketName,
		Endpoint:   file.Endpoint,
		KeyPrefix:  file.KeyPrefix,
		Size:       file.Size,
		FileSize:   file.FileSize,
	}
}
