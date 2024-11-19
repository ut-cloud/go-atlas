package model

import "gorm.io/gorm"

const TableNameBaseSecurity = "base_security"

// BaseSecurity 参数配置表
type BaseSecurity struct {
	Id           string         `gorm:"column:id;type:varchar(100);primaryKey;autoIncrement:true;comment:参数主键" json:"id"`
	ClientId     string         `gorm:"column:client_id;type:varchar(100);comment:客户端id" json:"client_id"`
	RegisterUrl  string         `gorm:"column:register_url;type:varchar(100);comment:注册地址" json:"register_url"`
	ClientSecret string         `gorm:"column:client_secret;type:varchar(500);comment:密钥" json:"client_secret"`
	GrantType    string         `gorm:"column:grant_types;type:varchar(500);comment:授权类型" json:"grant_type"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

// TableName BaseSecurity's table name
func (*BaseSecurity) TableName() string {
	return TableNameBaseSecurity
}
