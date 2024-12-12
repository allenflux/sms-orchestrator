package v1

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"sms_backend/api/v1/subUser"
)

func (c *ControllerSubUser) SubRegister(ctx context.Context, req *subUser.SubRegisterReq) (res *subUser.SubRegisterRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
