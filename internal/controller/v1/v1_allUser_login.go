package v1

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"sms_backend/api/v1/allUser"
)

func (c *ControllerAllUser) Login(ctx context.Context, req *allUser.LoginReq) (res *allUser.LoginRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
