package biz

import "github.com/google/wire"

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewCoreUsecase,
	NewAuthUsecase,
	NewSysUserUsecase,
	NewSysRoleUsecase,
	NewSysMenuUsecase,
	NewSysDeptUsecase,
	NewSysDictUsecase,
	NewSysConfigUsecase,
	NewSysPostUsecase,
)
