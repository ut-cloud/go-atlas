package service

import (
	pb "atlas-core/api/core/v1"
	"atlas-core/internal/biz"
	"atlas-core/internal/constants"
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/mojocn/base64Captcha"
	"github.com/ut-cloud/atlas-toolkit/utils"
	"strings"
)

type AuthService struct {
	pb.UnimplementedAuthServer
	auth *biz.AuthUsecase
	uc   *biz.SysUserUsecase
	rc   *biz.SysRoleUsecase
	mc   *biz.SysMenuUsecase
	ac   *biz.CoreUsecase
	log  *log.Helper
}

var Store = base64Captcha.DefaultMemStore

func NewAuthService(auth *biz.AuthUsecase, uc *biz.SysUserUsecase, rc *biz.SysRoleUsecase, mc *biz.SysMenuUsecase, ac *biz.CoreUsecase, logger log.Logger) *AuthService {
	return &AuthService{auth: auth, uc: uc, rc: rc, mc: mc, ac: ac, log: log.NewHelper(logger)}
}
func (s *AuthService) Logout(ctx context.Context, req *pb.LogoutReq) (*pb.LogoutReply, error) {
	return &pb.LogoutReply{}, nil
}
func (s *AuthService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginReply, error) {
	//判断验证码是否正确
	data, err := s.ac.RedisGetData(ctx, constants.CacheCaptcha+req.Uuid)
	if err != nil {
		return nil, errors.New("验证码失效")
	}
	if data.(string) != req.Code {
		//验证码错误则重新获取验证码
		_ = s.ac.RedisRemoveData(ctx, constants.CacheCaptcha+req.Uuid)
		return nil, errors.New("验证码验证失败")
	}
	//根据用户名查询用户信息
	user, err := s.uc.GetUserByUserName(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	//判断密码是否匹配
	if !utils.Verify(user.Password, req.Password) {
		return nil, errors.New("密码验证失败")
	}
	//获取当前用户的code
	roles, err := s.rc.GetRoleByUserId(ctx, user.UserID)
	if err != nil {
		return nil, err
	}
	var roleKeys []string
	for i := range roles {
		roleKeys = append(roleKeys, roles[i].RoleKey)
	}
	//返回token
	token, err := utils.GenerateToken(strings.Join(roleKeys, ","), user.UserID, user.UserName)
	if err != nil {
		return nil, err
	}
	return &pb.LoginReply{Token: token}, nil
}
func (s *AuthService) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterReply, error) {
	//注册用户
	//角色权限分配
	return &pb.RegisterReply{}, nil
}
func (s *AuthService) Captcha(ctx context.Context, req *pb.CaptchaReq) (*pb.CaptchaReply, error) {
	driver := base64Captcha.NewDriverDigit(80, 250, 4, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, Store)
	id, b64s, answer, err := cp.Generate()
	if err != nil {
		return nil, err
	}
	//将验证码添加到redis中
	err = s.ac.RedisSetData(ctx, constants.CacheCaptcha+id, answer)
	if err != nil {
		return nil, err
	}

	return &pb.CaptchaReply{
		Img:           b64s,
		Uuid:          id,
		CaptchaEnable: true,
	}, nil
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
