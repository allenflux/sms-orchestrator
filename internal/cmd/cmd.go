package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/util/gmode"
	v1 "sms_backend/internal/controller/v1"
	"sms_backend/library/libOss"
)

func MiddlewareCORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			gmode.SetProduct()
			InitFun(ctx)
			apiServer := g.Server()
			//s.Group("/", func(group *ghttp.RouterGroup) {
			//	group.Middleware(ghttp.MiddlewareHandlerResponse, MiddlewareCORS)
			//	group.Bind(
			//		v1.NewCommon(),
			//	)
			//})
			apiServer.Group("/api/v1/sms", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse, MiddlewareCORS)
				//service.SysUser().CheckLogin,
				//service.SystemMiddleware().PrintAndHideError, service.SystemMiddleware().Ctx, service.SystemMiddleware().Auth, service.SystemMiddleware().PrintParams)
				group.Bind(
					v1.NewSms(),
				)
			})

			apiServer.Run()
			return nil
		},
	}
)

func InitFun(ctx context.Context) {
	CompressAndClearLog(ctx)
	libOss.InitOssClient(ctx)
	//DBUpgrade(ctx)
}
