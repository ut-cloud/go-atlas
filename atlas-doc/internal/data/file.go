package data

import (
	"atlas-doc/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type fileRepo struct {
	data *Data
	log  *log.Helper
}

func NewFileRepo(data *Data, logger log.Logger) biz.FileRepo {
	return &fileRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
