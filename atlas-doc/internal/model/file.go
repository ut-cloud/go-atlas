package model

// OssFile represents the OSS file table
type OssFile struct {
	UUID       string  `gorm:"type:varchar(32);primaryKey;not null" json:"uuid"`  // 主键
	FileName   string  `gorm:"type:varchar(255);default:null" json:"file_name"`   // 文件名称
	KeyURL     string  `gorm:"type:varchar(255);default:null" json:"key_url"`     // 地址
	BucketName string  `gorm:"type:varchar(255);default:null" json:"bucket_name"` // oss桶名
	Endpoint   string  `gorm:"type:varchar(255);default:null" json:"endpoint"`    // oss端点
	KeyPrefix  string  `gorm:"type:varchar(255);default:null" json:"key_prefix"`  // 文件地址目录
	Size       float64 `gorm:"type:double(20,2);default:null" json:"size"`        // 大小
	FileSize   string  `gorm:"type:varchar(16);default:null" json:"file_size"`    // 文件大小
	BaseEntity
}
