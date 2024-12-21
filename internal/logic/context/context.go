/*
* @desc:context-service
* @company:云南奇讯科技有限公司
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2022/9/23 14:51
 */

package context

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/spf13/cast"
	"reflect"
	"sms_backend/internal/consts"
	"sms_backend/internal/model"
	"sms_backend/internal/service"
)

func init() {
	service.RegisterContext(New())
}

type sContext struct{}

func New() *sContext {
	return &sContext{}
}

// Init 初始化上下文对象指针到上下文对象中，以便后续的请求流程中可以修改。
func (s *sContext) Init(r *ghttp.Request, customCtx *model.Context) {
	r.SetCtxVar(consts.CtxKey, customCtx)
}

// Get 获得上下文变量，如果没有设置，那么返回nil
func (s *sContext) Get(ctx context.Context) *model.Context {
	value := ctx.Value(consts.CtxKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*model.Context); ok {
		return localCtx
	}
	return nil
}

// SetUser 将上下文信息设置到上下文请求中，注意是完整覆盖
func (s *sContext) SetUser(ctx context.Context, ctxUser *model.ContextUser) {
	s.Get(ctx).User = ctxUser
}

// GetLoginUser 获取当前登陆用户信息
func (s *sContext) GetLoginUser(ctx context.Context) *model.ContextUser {
	context := s.Get(ctx)
	if context == nil {
		return nil
	}
	return context.User
}

func (s *sContext) GetUserType(ctx context.Context) consts.EnumUserType {
	claim, err := service.GfToken().ParseToken(g.RequestFromCtx(ctx))
	if err != nil {
		g.Log().Error(ctx, err)
		return 0
	}
	m, ok := claim.Data.(map[string]interface{})
	if !ok {
		err := fmt.Errorf("type is %v", reflect.TypeOf(claim.Data))
		g.Log().Error(ctx, err)
		panic(err)
	}
	g.Log().Debug(ctx, m)
	return consts.EnumUserType(cast.ToInt(m["Type"]))
}

func (s *sContext) GetUserClaimMap(ctx context.Context) map[string]interface{} {
	claim, err := service.GfToken().ParseToken(g.RequestFromCtx(ctx))
	if err != nil {
		g.Log().Error(ctx, err)
	}
	m, ok := claim.Data.(map[string]interface{})
	if !ok {
		err := fmt.Errorf("type is %v", reflect.TypeOf(claim.Data))
		g.Log().Error(ctx, err)
		panic(err)
	}
	g.Log().Debug(ctx, m["Data"])
	stringMapE, err := cast.ToStringMapE(m["Data"])
	if err != nil {
		g.Log().Error(ctx, err)
		panic(err)
	}
	return stringMapE
}

func (s *sContext) GetEntUserRoleId(ctx context.Context) int64 {
	claimMap := s.GetUserClaimMap(ctx)
	return cast.ToInt64(claimMap["role_id"])
}

func (s *sContext) GetNoticeConnMapKey(ctx context.Context) string {
	return fmt.Sprintf("%d-%d", service.Context().GetUserType(ctx), service.Context().GetUserId(ctx))
}

func (s *sContext) GetUserId(ctx context.Context) int {
	claimMap := s.GetUserClaimMap(ctx)
	return cast.ToInt(claimMap["user_id"])
}

func (s *sContext) GetUsername(ctx context.Context) string {
	claimMap := s.GetUserClaimMap(ctx)
	return cast.ToString(claimMap["user_name"])
}

func (s *sContext) GetSystemId(ctx context.Context) int {
	claimMap := s.GetUserClaimMap(ctx)
	return cast.ToInt(claimMap["system_id"])
}

func (s *sContext) GetPidList(ctx context.Context) []int {
	claimMap := s.GetUserClaimMap(ctx)
	return cast.ToIntSlice(claimMap["permission_list"])
}
