package biz

import (
	"atlas-core/internal/model"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type BaseSecurityRepo interface {
	GetInfo(ctx context.Context, query *model.BaseSecurity) (*model.BaseSecurity, error)
}

type BaseSecurityUsecase struct {
	repo BaseSecurityRepo
	log  *log.Helper
}

func NewBaseSecurityUsecase(repo BaseSecurityRepo, logger log.Logger) *BaseSecurityUsecase {
	return &BaseSecurityUsecase{repo: repo, log: log.NewHelper(logger)}
}
