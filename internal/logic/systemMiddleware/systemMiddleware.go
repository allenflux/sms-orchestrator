/*
* @desc:中间件
* @company:云南奇讯科技有限公司
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2022/9/23 15:05
 */

package systemMiddleware

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/pkg/errors"
	"io"
	"reflect"
	"sms_backend/internal/model"
	"sms_backend/internal/service"
	"strings"
)

func init() {
	service.RegisterSystemMiddleware(New())
}

func New() *sSystemMiddleware {
	return &sSystemMiddleware{}
}

type sSystemMiddleware struct{}

// Ctx 自定义上下文对象
func (s *sSystemMiddleware) Ctx(r *ghttp.Request) {
	ctx := r.GetCtx()
	// 初始化登录用户信息
	data, err := service.GfToken().ParseToken(r)
	if err != nil {
		// 执行下一步请求逻辑
		r.Middleware.Next()
	}
	if data != nil {
		context := new(model.Context)
		userData, ok := data.Data.(map[string]interface{})
		if !ok {
			g.Log().Errorf(ctx, "type is %v", reflect.TypeOf(data.Data))
			r.Middleware.Next()
		}
		err = gconv.Struct(userData["Data"], &context.User)
		if err != nil {
			g.Log().Error(ctx, err)
			// 执行下一步请求逻辑
			r.Middleware.Next()
		}
		service.Context().Init(r, context)
	}
	// 执行下一步请求逻辑
	r.Middleware.Next()
}

func (s *sSystemMiddleware) PrintParams(r *ghttp.Request) {
	if !strings.Contains(r.Request.URL.String(), "upload") {
		all, err := io.ReadAll(r.Request.Body)
		if err != nil {
			g.Log().Error(r.GetCtx(), err)
			return
		}
		g.Log().Debugf(r.GetCtx(), "request method:%v url:%v header:[%+v] body:%s", r.Method, r.URL, r.Header, string(all))
	}
	r.Middleware.Next()
}

func (s *sSystemMiddleware) PrintAndHideError(r *ghttp.Request) {
	r.Middleware.Next()
	err := r.GetError()
	if err != nil {
		if gerr, ok := err.(*gerror.Error); ok {
			if gerr.Code() == gcode.CodeInternalPanic {
				if strings.Contains(gerr.Error(), "exception recovered") { // 直接发生崩溃
					g.Log().Error(r.GetCtx(), errors.WithStack(err))
					r.Response.WriteOver(map[string]interface{}{"code": 500, "message": "服务内部崩溃"})
					return
				} else { // 被捕获的崩溃
					g.Log().Error(r.GetCtx(), errors.WithStack(err))
					r.Response.WriteOver(map[string]interface{}{"code": 500, "message": gerr.Error()})
					return
				}
			} else if gerr.Code() == gcode.CodeDbOperationError { // 数据库操作错误
				g.Log().Error(r.GetCtx(), errors.WithStack(err))
				r.Response.WriteOver(map[string]interface{}{"code": 500, "message": "服务内部错误"})
				return
			}
		} else {
			//g.Log().Error(r.GetCtx(), reflect.TypeOf(err), err)
			//r.SetError(gerror.New("服务内部错误"))
		}
	}
}
