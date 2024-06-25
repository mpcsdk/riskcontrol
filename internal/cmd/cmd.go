package cmd

import (
	"context"

	"riskcontral/internal/controller/nrpcserver"
	"riskcontral/internal/controller/riskctrl"
	"riskcontral/internal/controller/tfa"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/os/gcmd"
)

func MiddlewareErrorHandler(r *ghttp.Request) {
	r.Middleware.Next()
	if err := r.GetError(); err != nil {
		g.Log().Error(r.Context(), err)
		r.Response.ClearBuffer()

		code := gcode.CodeInternalError
		r.Response.WriteJson(ghttp.DefaultHandlerResponse{
			Code:    code.Code(),
			Message: code.Message(),
			Data:    gtrace.GetTraceID(r.GetCtx()),
		})
	}
}
func MiddlewareCORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
func ResponseHandler(r *ghttp.Request) {
	r.Middleware.Next()
	// There's custom buffer content, it then exits current handler.
	g.Log().Info(context.Background(), r.GetBodyString())
	if r.Response.BufferLength() > 0 {
		return
	}
	var (
		err = r.GetError()
		res = r.GetHandlerResponse()
		// code = mpccode.Code(err)
		code = gerror.Code(err)
	)
	r.SetError(nil)
	if code == gcode.CodeNil {
		if err != nil {
			code = gcode.CodeInternalError
		} else {
			code = gcode.CodeOK
		}
	}
	g.Log().Info(context.Background(), res)
	r.Response.WriteJson(ghttp.DefaultHandlerResponse{
		Code:    code.Code(),
		Message: code.Message(),
		Data: func() interface{} {
			if code.Code() != gcode.CodeOK.Code() {
				return code.Detail()
			} else {
				return res
			}
		}(),
	})
}

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {

			s := g.Server()
			s.Use(MiddlewareErrorHandler)
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Middleware(MiddlewareCORS)
				group.Middleware(ResponseHandler)
				group.Bind(
					riskctrl.NewV1(),
					tfa.NewV1(),
				)
				// group.Group("/tfa", func(group *ghttp.RouterGroup) {
				// 	group.Bind(
				// 	)
				// })
			})
			//
			nrpcserver.Init()
			///
			s.Run()
			return nil
		},
	}
)
