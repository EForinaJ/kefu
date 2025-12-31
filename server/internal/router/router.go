package router

import (
	"kefu-server/internal/controller/account"
	"kefu-server/internal/controller/auth"
	"kefu-server/internal/controller/product"
	"kefu-server/internal/controller/site"
	"kefu-server/internal/middleware"

	"github.com/gogf/gf/v2/net/ghttp"
)

func LoadRouter(s *ghttp.Server) {
	s.Group("/api/v1/kefu", func(group *ghttp.RouterGroup) {
		group.Middleware(ghttp.MiddlewareHandlerResponse)
		group.Bind(
			auth.NewV1(),
			site.NewV1(),
		).Middleware(middleware.Response)
		group.Bind(
			account.NewV1(),
			product.NewV1(),
		).Middleware(middleware.Auth).Middleware(middleware.Response)
	})
}
