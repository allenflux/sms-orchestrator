/*
* @desc:错误处理
* @company:云南奇讯科技有限公司
* @Author: yixiaohu
* @Date:   2022/3/2 14:53
 */

package liberr

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"sms_backend/library/libResponse"
)

func ErrIsNil(ctx context.Context, err error, msg ...string) {
	if !g.IsNil(err) {
		if len(msg) > 0 {
			g.Log().Error(ctx, err.Error())
			panic(msg[0])
		} else {
			panic(err.Error())
		}
	}
}

func ValueIsNil(value interface{}, msg string) {
	if g.IsNil(value) {
		panic(msg)
	}
}

// 后端服务内部错误
func ErrIsNilServerFail(ctx context.Context, err error) {
	if !g.IsNil(err) {
		g.Log().Error(ctx, err.Error())
		panic("后端服务内部错误")
	}
}

func WarningExit(ctx context.Context, msg string) {
	req := g.RequestFromCtx(ctx)
	req.Response.WriteOverExit(WarpMsg(msg))
}

func Warning(ctx context.Context, msg string) {
	req := g.RequestFromCtx(ctx)
	req.Response.WriteOver(WarpMsg(msg))
}

func WarpError(err error) libResponse.Response {
	return libResponse.Response{
		Code: 500,
		Msg:  err.Error(),
	}
}

func WarpMsg(msg string) libResponse.Response {
	return libResponse.Response{
		Code: 500,
		Msg:  msg,
	}
}
