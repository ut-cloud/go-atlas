package data

import (
	"atlas-core/internal/conf"
	"atlas-core/internal/model"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	redis2 "github.com/gomodule/redigo/redis"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var ProviderSet = wire.NewSet(NewData,
	model.NewDB,
	model.NewRedisPool,
	model.NewRedis,
	NewCoreRepo,
	NewSysUserRepo,
	NewSysRoleRepo,
	NewSysMenuRepo,
	NewSysDeptRepo,
	NewSysDictRepo,
	NewSysConfigRepo,
	NewSysPostRepo,
	NewBaseSecurityRepo,
)

type Data struct {
	Db    *gorm.DB
	Rdb   *redis.Client
	RPool *redis2.Pool
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, db *gorm.DB, rdb *redis.Client, rPool *redis2.Pool) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{Db: db, Rdb: rdb, RPool: rPool}, cleanup, nil
}
