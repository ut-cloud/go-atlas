package model

import (
	"atlas-core/internal/conf"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	redis2 "github.com/gomodule/redigo/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	slog "log"
	"os"
	"time"
)

func NewDB(c *conf.Data) *gorm.DB {
	newLogger := logger.New(
		slog.New(os.Stdout, "\r\n", slog.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //慢查询sql阀值
			Colorful:      true,        //禁用彩色打印
			LogLevel:      logger.Info,
		},
	)
	db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{
		Logger:                                   newLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy:                           schema.NamingStrategy{
			//SingularTable: true, //表名是否加s
		},
	})
	if err != nil {
		log.Errorf("failed opening connection to sqlite: %v", err)
		panic("failed to connect database")
	}
	//err2 := db.AutoMigrate(&User{})
	//if err2 != nil {
	//	panic("AutoMigrate is failed")
	//}
	return db
}

func NewRedis(c *conf.Data) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		DB:           int(c.Redis.Db),
		Password:     c.Redis.Password,
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
	})
	// ping Redis客户端，判断连接是否存在
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Error("Redis database connection failure, err : %v", err)
	}
	log.Info("Cache enabled successfully!")
	return rdb
}

// NewRedisPool 使用redisgo库创建 Redis 连接池 用于casbin使用
func NewRedisPool(c *conf.Data) *redis2.Pool {
	return &redis2.Pool{
		MaxIdle:     10,
		MaxActive:   20, // 最大连接数，0 为无限制
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis2.Conn, error) {
			conn, err := redis2.Dial("tcp", c.Redis.Addr,
				redis2.DialPassword(c.Redis.Password),
				redis2.DialDatabase(int(c.Redis.Db))) // 指定数据库
			if err != nil {
				return nil, err
			}
			return conn, nil
		},
		TestOnBorrow: func(c redis2.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
