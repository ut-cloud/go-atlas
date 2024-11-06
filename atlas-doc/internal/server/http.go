package server

import (
	v1 "atlas-doc/api/doc/v1"
	"atlas-doc/internal/conf"
	"atlas-doc/internal/service"
	"context"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/gorilla/handlers"
	middleware2 "github.com/ut-cloud/atlas-toolkit/middleware"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server,
	file *service.FileService,
	logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		NewMiddleware(),
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
	srv := http.NewServer(opts...)
	v1.RegisterFileHTTPServer(srv, file)
	route := srv.Route("/")
	route.POST("/doc/upload", file.Upload())
	return srv
}

// NewMiddleware 创建中间件
func NewMiddleware() http.ServerOption {
	return http.Middleware(
		validate.Validator(),
		recovery.Recovery(),
		selector.Server(
			middleware2.Auth(),
		).Match(NewWhiteListMatcher()).Build(),
	)
}

// NewWhiteListMatcher 设置白名单，不需要 token 验证的接口
func NewWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]struct{})
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}
