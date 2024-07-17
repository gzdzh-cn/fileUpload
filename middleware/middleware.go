package middleware

import (
	"dzhgo/internal/config"
	"dzhgo/internal/middleware"
	"github.com/gogf/gf/v2/frame/g"
)

func init() {
	if config.Config.Middleware.Authority.Enable {
		g.Server().BindMiddleware("/app/fileUpload/*", BaseAuthorityMiddleware)
		g.Server().BindMiddleware("/admin/fileUpload/steam/getProcessById", BaseAuthorityMiddleware)
		g.Server().BindMiddleware("/admin/fileUpload/steam/getProcessById", middleware.BaseAuthorityMiddlewareOpen)
	}

}
