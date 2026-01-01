package router

import (
	"server/internal/controller/account"
	"server/internal/controller/auth"
	"server/internal/controller/order"
	"server/internal/controller/product"
	"server/internal/controller/site"
	"server/internal/controller/witkey"
	"server/internal/middleware"

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
			order.NewV1(),
			witkey.NewV1(),
		).Middleware(middleware.Auth).Middleware(middleware.Response)
	})
}
