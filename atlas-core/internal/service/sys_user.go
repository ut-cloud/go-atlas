package service

import (
	v1 "atlas-core/api/core/v1"
	"atlas-core/internal/biz"
	"atlas-core/internal/model"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/ut-cloud/atlas-toolkit/utils"
)

type SysUserService struct {
	v1.UnimplementedSysUserServer
	uc  *biz.SysUserUsecase
	rc  *biz.SysRoleUsecase
	pc  *biz.SysPostUsecase
	log *log.Helper
}

func NewSysUserService(uc *biz.SysUserUsecase, rc *biz.SysRoleUsecase, pc *biz.SysPostUsecase, logger log.Logger) *SysUserService {
	return &SysUserService{uc: uc, rc: rc, pc: pc, log: log.NewHelper(logger)}
}

func (s *SysUserService) CreateSysUser(ctx context.Context, req *v1.CreateSysUserRep) (*v1.CreateSysUserReply, error) {
	return &v1.CreateSysUserReply{}, nil
}
func (s *SysUserService) UpdateSysUser(ctx context.Context, req *v1.UpdateSysUserRep) (*v1.UpdateSysUserReply, error) {
	user := &model.BizSysUser{
		UserID:      req.GetUserId(),
		DeptID:      req.GetDeptId(),
		UserName:    req.GetUserName(),
		NickName:    req.GetNickName(),
		Password:    req.GetPassword(),
		PhoneNumber: req.GetPhoneNumber(),
		Email:       req.GetEmail(),
		Status:      req.GetStatus(),
		Remark:      req.GetRemark(),
		Sex:         req.GetSex(),
	}
	err := s.uc.Update(ctx, user)
	if err != nil {
		return nil, err
	}
	if len(req.GetPostIds()) > 0 {
		err = s.rc.ModifyRoleForUser(ctx, req.GetUserId(), req.GetRoleIds())
		if err != nil {
			return nil, err
		}

	}
	if len(req.RoleIds) > 0 {
		err = s.pc.ModifyPostForUser(ctx, req.GetUserId(), req.GetPostIds())
		if err != nil {
			return nil, err
		}
	}
	return &v1.UpdateSysUserReply{}, nil
}
func (s *SysUserService) ResetPwd(ctx context.Context, req *v1.ResetPwdRep) (*v1.ResetPwdReply, error) {
	user := &model.BizSysUser{
		UserID:   req.GetUserId(),
		Password: req.GetPassword(),
	}
	err := s.uc.Update(ctx, user)
	if err != nil {
		return nil, err
	}
	return &v1.ResetPwdReply{}, err
}
func (s *SysUserService) DeleteSysUser(ctx context.Context, req *v1.DeleteSysUserRep) (*v1.DeleteSysUserReply, error) {
	err := s.uc.Delete(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &v1.DeleteSysUserReply{}, nil
}
func (s *SysUserService) SaveSysUser(ctx context.Context, req *v1.SaveSysUserRep) (*v1.EmptyReply, error) {
	id := utils.GetID()
	user := &model.BizSysUser{
		UserID:      id,
		DeptID:      req.GetDeptId(),
		UserName:    req.GetUserName(),
		NickName:    req.GetNickName(),
		Password:    req.GetPassword(),
		PhoneNumber: req.GetPhoneNumber(),
		Email:       req.GetEmail(),
		Status:      req.GetStatus(),
		Remark:      req.GetRemark(),
		Sex:         req.GetSex(),
	}
	err := s.uc.Save(ctx, user)
	if err != nil {
		return nil, err
	}
	if len(req.GetRoleIds()) > 0 {
		err = s.rc.ModifyRoleForUser(ctx, id, req.GetRoleIds())
		if err != nil {
			return nil, err
		}

	}
	if len(req.GetPostIds()) > 0 {
		err = s.pc.ModifyPostForUser(ctx, id, req.GetPostIds())
		if err != nil {
			return nil, err
		}
	}
	return &v1.EmptyReply{}, nil
}
func (s *SysUserService) GetSysUser(ctx context.Context, req *v1.GetSysUserRep) (*v1.GetSysUserReply, error) {
	user, err := s.uc.GetUserInfoById(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	myRoles, err := s.rc.GetRoleByUserId(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	var roleIds []string
	for _, role := range myRoles {
		roleIds = append(roleIds, role.RoleId)
	}
	roles, err := s.rc.GetListRoles(ctx)
	if err != nil {
		return nil, err
	}
	posts, err := s.pc.GetPostList(ctx, &model.SysPostQuery{})
	if err != nil {
		return nil, err
	}
	myPosts, err := s.pc.GetPostList(ctx, &model.SysPostQuery{UserId: req.GetId()})
	if err != nil {
		return nil, err
	}
	var postIds = make([]string, 0)
	for _, post := range myPosts {
		postIds = append(postIds, post.PostId)
	}
	return &v1.GetSysUserReply{
		User:    user,
		Roles:   roles,
		RoleIds: roleIds,
		Posts:   posts,
		PostIds: postIds,
	}, nil
}

func (s *SysUserService) GetOtherInfo(ctx context.Context, req *v1.GetOtherInfoRep) (*v1.GetOtherInfoReply, error) {
	roles, err := s.rc.GetListRoles(ctx)
	if err != nil {
		return nil, err
	}
	posts, err := s.pc.GetPostList(ctx, &model.SysPostQuery{})
	if err != nil {
		return nil, err
	}
	return &v1.GetOtherInfoReply{
		Roles: roles,
		Posts: posts,
	}, nil
}

func (s *SysUserService) ListSysUser(ctx context.Context, req *v1.ListSysUserRep) (*v1.ListSysUserReply, error) {
	return s.uc.GetPageSet(ctx, req, 0)
}

func (s *SysUserService) Profile(ctx context.Context, req *v1.ProfileRep) (*v1.ProfileReply, error) {
	userId := utils.GetLoginUserId(ctx)
	userInfo, err := s.uc.GetUserInfoById(ctx, userId)
	if err != nil {
		return nil, err
	}
	return &v1.ProfileReply{User: userInfo}, nil
}
func (s *SysUserService) GetAuthRoleSysUser(ctx context.Context, req *v1.GetAuthRoleSysUserRep) (*v1.GetAuthRoleSysUserReply, error) {
	return s.uc.GetAuthRoleSysUser(ctx, req)
}
func (s *SysUserService) AuthRoleSysUser(ctx context.Context, req *v1.AuthRoleSysUserRep) (*v1.EmptyReply, error) {
	return s.uc.AuthRoleSysUser(ctx, req)
}
func (s *SysUserService) UpdateProfile(ctx context.Context, req *v1.UpdateProfileRep) (*v1.EmptyReply, error) {
	return s.uc.UpdateProfile(ctx, req)
}
func (s *SysUserService) UpdatePassword(ctx context.Context, req *v1.UpdatePasswordRep) (*v1.EmptyReply, error) {
	return s.uc.UpdatePassword(ctx, req)
}
