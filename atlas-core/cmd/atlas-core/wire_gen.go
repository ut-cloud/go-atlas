// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"atlas-core/internal/biz"
	"atlas-core/internal/conf"
	"atlas-core/internal/data"
	"atlas-core/internal/model"
	"atlas-core/internal/server"
	"atlas-core/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	grpcServer := server.NewGRPCServer(confServer, logger)
	pool := model.NewRedisPool(confData)
	db := model.NewDB(confData)
	client := model.NewRedis(confData)
	dataData, cleanup, err := data.NewData(confData, logger, db, client, pool)
	if err != nil {
		return nil, nil, err
	}
	sysMenuRepo := data.NewSysMenuRepo(dataData, logger)
	sysUserRepo := data.NewSysUserRepo(dataData, logger)
	sysRoleRepo := data.NewSysRoleRepo(dataData, logger)
	sysUserUsecase := biz.NewSysUserUsecase(sysUserRepo, sysRoleRepo, logger)
	sysDeptRepo := data.NewSysDeptRepo(dataData, logger)
	sysRoleUsecase := biz.NewSysRoleUsecase(sysRoleRepo, sysMenuRepo, sysDeptRepo, sysUserRepo, logger)
	authUsecase := biz.NewAuthUsecase(sysMenuRepo, sysUserUsecase, sysRoleUsecase, logger)
	sysMenuUsecase := biz.NewSysMenuUsecase(sysMenuRepo, sysRoleRepo, logger)
	coreRepo := data.NewCoreRepo(dataData, logger)
	coreUsecase := biz.NewCoreUsecase(coreRepo, logger)
	authService := service.NewAuthService(authUsecase, sysUserUsecase, sysRoleUsecase, sysMenuUsecase, coreUsecase, logger)
	sysPostRepo := data.NewSysPostRepo(dataData, logger)
	sysPostUsecase := biz.NewSysPostUsecase(sysPostRepo, logger)
	sysUserService := service.NewSysUserService(sysUserUsecase, sysRoleUsecase, sysPostUsecase, logger)
	sysRoleService := service.NewSysRoleService(sysRoleUsecase, logger)
	sysMenuService := service.NewSysMenuService(sysMenuUsecase, logger)
	sysDeptUsecase := biz.NewSysDeptUsecase(sysDeptRepo, logger)
	sysDeptService := service.NewSysDeptService(sysDeptUsecase, logger)
	sysDictRepo := data.NewSysDictRepo(dataData, logger)
	sysDictUsecase := biz.NewSysDictUsecase(sysDictRepo, coreRepo, logger)
	sysDictService := service.NewSysDictService(sysDictUsecase, logger)
	sysConfigRepo := data.NewSysConfigRepo(dataData, logger)
	sysConfigUsecase := biz.NewSysConfigUsecase(sysConfigRepo, logger)
	sysConfigService := service.NewSysConfigService(sysConfigUsecase, logger)
	sysPostService := service.NewSysPostService(sysPostUsecase, logger)
	monitorService := service.NewMonitorService()
	httpServer := server.NewHTTPServer(confServer, pool, authUsecase, logger, authService, sysUserService, sysRoleService, sysMenuService, sysDeptService, sysDictService, sysConfigService, sysPostService, monitorService)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
