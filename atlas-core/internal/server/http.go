package server

import (
	v1 "atlas-core/api/core/v1"
	"atlas-core/internal/conf"
	"atlas-core/internal/service"
	"context"
	"github.com/casbin/casbin/v2/model"
	fileAdapter "github.com/casbin/casbin/v2/persist/file-adapter"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/gorilla/handlers"
	casbinM "github.com/ut-cloud/atlas-toolkit/casbin"
	middleware2 "github.com/ut-cloud/atlas-toolkit/middleware"
	"github.com/ut-cloud/atlas-toolkit/response"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, logger log.Logger,
	auth *service.AuthService,
	sysUser *service.SysUserService,
	sysRole *service.SysRoleService,
	sysMenu *service.SysMenuService,
	sysDept *service.SysDeptService,
	sysDict *service.SysDictService,
	sysConfig *service.SysConfigService,
	sysPost *service.SysPostService,
	monitor *service.MonitorService) *http.Server {
	var opts = []http.ServerOption{
		NewMiddleware(logger),
		http.Filter(handlers.CORS(
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}),
		)),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	opts = append(opts, http.ResponseEncoder(response.EncoderResponse()))
	opts = append(opts, http.ErrorEncoder(response.EncoderError()))
	srv := http.NewServer(opts...)
	v1.RegisterAuthHTTPServer(srv, auth)
	v1.RegisterSysUserHTTPServer(srv, sysUser)
	v1.RegisterSysRoleHTTPServer(srv, sysRole)
	v1.RegisterSysMenuHTTPServer(srv, sysMenu)
	v1.RegisterSysDeptHTTPServer(srv, sysDept)
	v1.RegisterSysDictHTTPServer(srv, sysDict)
	v1.RegisterSysConfigHTTPServer(srv, sysConfig)
	v1.RegisterSysPostHTTPServer(srv, sysPost)
	v1.RegisterMonitorHTTPServer(srv, monitor)
	return srv
}

// NewMiddleware 创建中间件
func NewMiddleware(logger log.Logger) http.ServerOption {
	m, _ := model.NewModelFromFile("configs/authz/authz_model.conf")
	a := fileAdapter.NewAdapter("configs/authz/authz_policy.csv")
	return http.Middleware(
		validate.Validator(),
		recovery.Recovery(),
		selector.Server(
			middleware2.Auth(),
			casbinM.Server(
				casbinM.WithCasbinModel(m),
				casbinM.WithCasbinPolicy(a),
				casbinM.WithSecurityUserCreator(middleware2.NewSecurityUser),
			),
		).Match(NewWhiteListMatcher()).Build(),
	)
}

// NewWhiteListMatcher 设置白名单，不需要 token 验证的接口
func NewWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]struct{})
	whiteList["/api.system.v1.Auth/Login"] = struct{}{}
	whiteList["/api.system.v1.Auth/Captcha"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}
