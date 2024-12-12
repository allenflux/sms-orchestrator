package v1

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"sms_backend/api/v1/subUser"
)

func (c *ControllerSubUser) SubUpdate(ctx context.Context, req *subUser.SubUpdateReq) (res *subUser.SubUpdateRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
