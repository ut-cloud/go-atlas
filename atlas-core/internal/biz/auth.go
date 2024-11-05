package biz

import (
	v1 "atlas-core/api/core/v1"
	"atlas-core/internal/constants"
	"atlas-core/internal/model"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/ut-cloud/atlas-toolkit/stream"
	"github.com/ut-cloud/atlas-toolkit/utils"
)

type AuthUsecase struct {
	menuRepo SysMenuRepo
	uc       *SysUserUsecase
	rc       *SysRoleUsecase
	log      *log.Helper
}

func NewAuthUsecase(menuRepo SysMenuRepo, uc *SysUserUsecase, rc *SysRoleUsecase, logger log.Logger) *AuthUsecase {
	return &AuthUsecase{menuRepo: menuRepo, uc: uc, rc: rc, log: log.NewHelper(logger)}
}

func (u AuthUsecase) UserInfo(ctx context.Context, req *v1.UserInfoReq) (*v1.UserInfoReply, error) {
	userId := utils.GetLoginUserId(ctx)
	user, _ := u.uc.GetUserInfoById(ctx, userId)
	roles, _ := u.rc.GetRoleByUserId(ctx, userId)
	user.Roles = roles
	roleCodes := make([]string, 0)
	permissions := make([]string, 0)
	isAdmin := false
	for i := range roles {
		roleCodes = append(roleCodes, roles[i].RoleKey)
		if roles[i].RoleKey == constants.SuperAdmin {
			isAdmin = true
		}
	}
	if isAdmin {
		permissions = append(permissions, constants.AllPermission)
	} else {
		rolePerms, _ := u.menuRepo.GetRolePerms(ctx)
		stream.NewStream(rolePerms).
			Filter(func(perm *model.RoleMenuPerm) bool {
				return contains(roleCodes, perm.RoleKey)
			}).
			Filter(func(perm *model.RoleMenuPerm) bool {
				return perm.Perms != ""
			}).
			ForEach(ctx, func(perm *model.RoleMenuPerm) {
				permissions = append(permissions, perm.Perms)
			})
		permissions = distinct(permissions)
	}
	return &v1.UserInfoReply{
		Permissions: permissions,
		Roles:       roleCodes,
		User:        user,
	}, nil
}

func (u AuthUsecase) GetAllPerms(ctx context.Context) ([]*model.RoleMenuPerm, []string, error) {
	perms, err := u.menuRepo.GetRolePerms(ctx)
	roleKeys := make([]string, 0)
	roles, err := u.rc.GetListRoles(ctx)
	for i := range roles {
		roleKeys = append(roleKeys, roles[i].RoleKey)
	}
	return perms, roleKeys, err
}

// unique 函数去除切片中的重复元素
func distinct[T comparable](slice []T) []T {
	seen := make(map[T]struct{}) // 用于跟踪已经见过的元素
	var result []T               // 存放去重后的结果
	for _, v := range slice {
		if _, ok := seen[v]; !ok { // 检查元素是否已见过
			seen[v] = struct{}{}       // 标记元素为已见
			result = append(result, v) // 添加到结果切片
		}
	}
	return result
}

// 定义一个通用的 contains 函数
func contains[T comparable](slice []T, element T) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}
