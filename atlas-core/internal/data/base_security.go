package data

import (
	"atlas-core/internal/biz"
	"atlas-core/internal/model"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type baseSecurityRepo struct {
	data *Data
	log  *log.Helper
}

func NewBaseSecurityRepo(data *Data, logger log.Logger) biz.BaseSecurityRepo {
	return &baseSecurityRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (s baseSecurityRepo) GetInfo(ctx context.Context, query *model.BaseSecurity) (*model.BaseSecurity, error) {
	var baseSecurity model.BaseSecurity
	db := s.data.Db.Model(&model.BaseSecurity{})
	if query.GrantType != "" {
		db = db.Where("grant_types = ?", query.GrantType)
	}
	if err := db.First(&baseSecurity).Error; err != nil {
		return nil, err
	}
	return &baseSecurity, nil
}
