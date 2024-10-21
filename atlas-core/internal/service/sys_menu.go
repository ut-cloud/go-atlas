package service

import (
	v1 "atlas-core/api/core/v1"
	"atlas-core/internal/biz"
	"atlas-core/internal/pkg"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type SysMenuService struct {
	v1.UnimplementedSysMenuServer
	uc  *biz.SysMenuUsecase
	log *log.Helper
}

func NewSysMenuService(uc *biz.SysMenuUsecase, logger log.Logger) *SysMenuService {
	return &SysMenuService{uc: uc, log: log.NewHelper(logger)}
}

func (s *SysMenuService) CreateSysMenu(ctx context.Context, req *v1.SysMenuRep) (*v1.EmptyReply, error) {
	req.MenuId = pkg.GetID()
	return s.uc.Save(ctx, req)
}

func (s *SysMenuService) UpdateSysMenu(ctx context.Context, req *v1.SysMenuRep) (*v1.EmptyReply, error) {
	return s.uc.Update(ctx, req)
}
func (s *SysMenuService) DeleteSysMenu(ctx context.Context, req *v1.DeleteSysMenuRep) (*v1.EmptyReply, error) {
	return s.uc.Delete(ctx, req.Id)
}
func (s *SysMenuService) GetSysMenu(ctx context.Context, req *v1.GetSysMenuRep) (*v1.GetSysMenuReply, error) {
	return s.uc.GetSysMenu(ctx, req)
}
func (s *SysMenuService) ListSysMenu(ctx context.Context, req *v1.ListSysMenuRep) (*v1.ListSysMenuReply, error) {
	return s.uc.GetMenusList(ctx, req)
}
func (s *SysMenuService) GetSysRoleMenu(ctx context.Context, req *v1.GetSysRoleMenuRep) (*v1.GetSysRoleMenuReply, error) {
	return s.uc.GetSysRoleMenu(ctx, req)
}
func (s *SysMenuService) GetTreeSelect(ctx context.Context, req *v1.GetTreeSelectRep) (*v1.GetTreeSelectReply, error) {
	return s.uc.GetTreeSelect(ctx, req)
}
