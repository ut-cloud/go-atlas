package biz

import (
	v1 "atlas-core/api/core/v1"
	"atlas-core/internal/constants"
	"atlas-core/internal/model"
	"context"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/mojocn/base64Captcha"
	casbinM "github.com/ut-cloud/atlas-toolkit/casbin"
	"github.com/ut-cloud/atlas-toolkit/stream"
	"github.com/ut-cloud/atlas-toolkit/utils"
	"google.golang.org/api/idtoken"
	"strings"
)

type AuthUsecase struct {
	menuRepo SysMenuRepo
	uc       *SysUserUsecase
	rc       *SysRoleUsecase
	ac       *CoreUsecase
	bc       *BaseSecurityUsecase
	log      *log.Helper
}

func NewAuthUsecase(menuRepo SysMenuRepo, uc *SysUserUsecase, rc *SysRoleUsecase, ac *CoreUsecase, bc *BaseSecurityUsecase, logger log.Logger) *AuthUsecase {
	return &AuthUsecase{menuRepo: menuRepo, uc: uc, rc: rc, ac: ac, bc: bc, log: log.NewHelper(logger)}
}

var Store = base64Captcha.DefaultMemStore

// UserInfo 用来存储从 Google 返回的用户信息
type UserInfo struct {
	ID        string `json:"sub"`
	Email     string `json:"email"`
	FirstName string `json:"given_name"`
	LastName  string `json:"family_name"`
	FullName  string `json:"name"`
	Picture   string `json:"picture"`
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

func (u AuthUsecase) GetAllPerms(ctx context.Context) ([]*casbinM.RoleMenuPerm, []string, error) {
	perms, err := u.menuRepo.GetRolePerms(ctx)
	rolePerms := make([]*casbinM.RoleMenuPerm, len(perms))
	for i := range perms {
		rolePerms[i] = &casbinM.RoleMenuPerm{
			RoleKey: perms[i].RoleKey,
			Perms:   perms[i].Perms,
		}
	}
	roleKeys := make([]string, 0)
	roles, err := u.rc.GetListRoles(ctx)
	for i := range roles {
		roleKeys = append(roleKeys, roles[i].RoleKey)
	}
	return rolePerms, roleKeys, err
}

func (u AuthUsecase) GoogleLogin(ctx context.Context, req *v1.LoginReq) (*v1.LoginReply, error) {
	baseSecurity, err := u.bc.repo.GetInfo(ctx, &model.BaseSecurity{GrantType: req.GetGrantType()})
	if err != nil {
		return nil, err
	}
	userInfo, err := validateIDToken(req.GetCredential(), baseSecurity.ClientId)
	if err != nil {
		return nil, err
	}
	// 检查返回的用户信息
	if userInfo == nil {
		return nil, errors.New("invalid user info")
	}
	//判断用户是否存在
	_, err = u.uc.GetUserInfoById(ctx, userInfo.ID)
	if err != nil {
		user := &model.BizSysUser{
			UserID:   userInfo.ID,
			UserName: userInfo.Email,
			NickName: userInfo.FullName,
			Password: "atlas@go-atlas.cn",
			Email:    userInfo.Email,
			Avatar:   userInfo.Picture,
		}
		err := u.uc.Save(ctx, user)
		if err != nil {
			return nil, err
		}
		err = u.rc.ModifyRoleForUser(ctx, userInfo.ID, []string{"2"})
		if err != nil {
			return nil, err
		}
	}
	//获取当前用户的code
	roles, err := u.rc.GetRoleByUserId(ctx, userInfo.ID)
	if err != nil {
		return nil, err
	}
	var roleKeys []string
	for i := range roles {
		roleKeys = append(roleKeys, roles[i].RoleKey)
	}
	//返回token
	token, err := utils.GenerateToken(strings.Join(roleKeys, ","), userInfo.ID, userInfo.FullName)
	if err != nil {
		return nil, err
	}
	return &v1.LoginReply{Token: token}, nil
}

func (u AuthUsecase) PwdLogin(ctx context.Context, req *v1.LoginReq) (*v1.LoginReply, error) {
	//判断验证码是否正确
	data, err := u.ac.RedisGetData(ctx, constants.CacheCaptcha+req.Uuid)
	if err != nil {
		return nil, errors.New("验证码失效")
	}
	if data.(string) != req.Code {
		//验证码错误则重新获取验证码
		_ = u.ac.RedisRemoveData(ctx, constants.CacheCaptcha+req.Uuid)
		return nil, errors.New("验证码验证失败")
	}
	//根据用户名查询用户信息
	user, err := u.uc.GetUserByUserName(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	//判断密码是否匹配
	if !utils.Verify(user.Password, req.Password) {
		return nil, errors.New("密码验证失败")
	}
	//获取当前用户的code
	roles, err := u.rc.GetRoleByUserId(ctx, user.UserID)
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
	return &v1.LoginReply{Token: token}, nil
}

func (u AuthUsecase) Captcha(ctx context.Context, req *v1.CaptchaReq) (*v1.CaptchaReply, error) {
	driver := base64Captcha.NewDriverDigit(80, 250, 4, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, Store)
	id, b64s, answer, err := cp.Generate()
	if err != nil {
		return nil, err
	}
	//将验证码添加到redis中
	err = u.ac.RedisSetData(ctx, constants.CacheCaptcha+id, answer)
	if err != nil {
		return nil, err
	}

	return &v1.CaptchaReply{
		Img:           b64s,
		Uuid:          id,
		CaptchaEnable: true,
	}, nil
}

// validateIDToken validates the id_token and returns user information
func validateIDToken(idToken, googleClientID string) (*UserInfo, error) {
	// Context for the token validation
	ctx := context.Background()

	// Validate the ID token and retrieve the payload
	payload, err := idtoken.Validate(ctx, idToken, googleClientID)
	if err != nil {
		return nil, fmt.Errorf("failed to validate ID token: %v", err)
	}
	fmt.Println(payload)
	// Extract the claims map
	claims := payload.Claims

	// Extract relevant user info from the claims map
	userInfo := &UserInfo{
		ID:        claims["sub"].(string),         // The "sub" claim is the unique ID for the user
		Email:     claims["email"].(string),       // The "email" claim contains the user's email
		FirstName: claims["given_name"].(string),  // First name
		LastName:  claims["family_name"].(string), // Last name
		FullName:  claims["name"].(string),        // Full name
		Picture:   claims["picture"].(string),     // Profile picture URL
	}

	return userInfo, nil
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
