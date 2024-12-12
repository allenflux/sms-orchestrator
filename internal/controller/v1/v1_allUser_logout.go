package v1

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"sms_backend/api/v1/allUser"
)

func (c *ControllerAllUser) Logout(ctx context.Context, req *allUser.LogoutReq) (res *allUser.LogoutRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
