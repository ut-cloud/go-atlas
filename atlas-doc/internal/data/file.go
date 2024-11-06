package data

import (
	"atlas-doc/internal/biz"
	"atlas-doc/internal/model"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type fileRepo struct {
	data *Data
	log  *log.Helper
}

func (f fileRepo) GetInfo(ctx context.Context, id string) (*model.BizFile, error) {
	var fileInfo model.File
	if err := f.data.Db.Model(&fileInfo).Where("uuid = ?", id).First(&fileInfo).Error; err != nil {
		return nil, err
	}
	return fileInfo.EntityToBiz(), nil
}

func (f fileRepo) Save(ctx context.Context, file *model.File) error {
	if err := f.data.Db.Save(file).Error; err != nil {
		return err
	}
	return nil
}

func (f fileRepo) Update(ctx context.Context, file *model.File) error {
	//TODO implement me
	panic("implement me")
}

func (f fileRepo) Delete(ctx context.Context, fileId string) error {
	//TODO implement me
	panic("implement me")
}

func (f fileRepo) GetPageSet(ctx context.Context, pageNum int64, pageSize int64, query model.File) ([]*model.BizFile, int64, error) {
	//TODO implement me
	panic("implement me")
}

func NewFileRepo(data *Data, logger log.Logger) biz.FileRepo {
	return &fileRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
