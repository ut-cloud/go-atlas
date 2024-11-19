package service

import (
	pb "atlas-core/api/core/v1"
	"atlas-core/internal/biz"
	"atlas-core/internal/model"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/ut-cloud/atlas-toolkit/utils"
)

type AuthService struct {
	pb.UnimplementedAuthServer
	auth *biz.AuthUsecase
	uc   *biz.SysUserUsecase
	rc   *biz.SysRoleUsecase
	mc   *biz.SysMenuUsecase
	log  *log.Helper
}

func NewAuthService(auth *biz.AuthUsecase, uc *biz.SysUserUsecase, rc *biz.SysRoleUsecase, mc *biz.SysMenuUsecase, logger log.Logger) *AuthService {
	return &AuthService{auth: auth, uc: uc, rc: rc, mc: mc, log: log.NewHelper(logger)}
}
func (s *AuthService) Logout(ctx context.Context, req *pb.LogoutReq) (*pb.LogoutReply, error) {
	return &pb.LogoutReply{}, nil
}
func (s *AuthService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginReply, error) {

	if req.GrantType == "password" {
		return s.auth.PwdLogin(ctx, req)
	}
	if req.GrantType == "google" {
		return s.auth.GoogleLogin(ctx, req)
	}
	return nil, nil

}
func (s *AuthService) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterReply, error) {
	id := utils.GetID()
	user := &model.BizSysUser{
		UserID:      id,
		UserName:    req.GetUsername(),
		NickName:    req.GetNickname(),
		Password:    req.GetPassword(),
		PhoneNumber: req.GetMobile(),
		Email:       req.GetEmail(),
	}
	err := s.uc.Save(ctx, user)
	if err != nil {
		return nil, err
	}
	return &pb.RegisterReply{}, nil
}
func (s *AuthService) Captcha(ctx context.Context, req *pb.CaptchaReq) (*pb.CaptchaReply, error) {
	return s.auth.Captcha(ctx, req)
}
func (s *AuthService) UserInfo(ctx context.Context, req *pb.UserInfoReq) (*pb.UserInfoReply, error) {
	return s.auth.UserInfo(ctx, req)
}
func (s *AuthService) Routers(ctx context.Context, req *pb.RoutersReq) (*pb.RoutersReply, error) {
	menus, err := s.mc.GetMenuByUserId(ctx, utils.GetLoginUserId(ctx))
	if err != nil {
		return nil, err
	}
	return &pb.RoutersReply{Routers: menus}, nil
}
