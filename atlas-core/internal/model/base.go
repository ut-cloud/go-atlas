package model

import (
	"context"
	"github.com/go-kratos/kratos/v2/metadata"
	"gorm.io/gorm"
	"time"
)

type BaseEntity struct {
	CreatedAt time.Time
	CreateBy  string `gorm:"type:varchar(100);not null;comment:创建人" json:"create_by"`
	UpdatedAt time.Time
	UpdateBy  string         `gorm:"type:varchar(100);not null;comment:更新人" json:"update_by"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (m *BaseEntity) insertEntity(ctx context.Context) {
	var uuid string
	if md, ok := metadata.FromServerContext(ctx); ok {
		uuid = md.Get("identity")
	}
	m.CreateBy = uuid
}

func (m *BaseEntity) updateEntity(ctx context.Context) {
	var uuid string
	if md, ok := metadata.FromServerContext(ctx); ok {
		uuid = md.Get("identity")
	}
	m.UpdateBy = uuid
}
