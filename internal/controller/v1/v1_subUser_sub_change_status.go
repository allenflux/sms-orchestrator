package v1

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"sms_backend/api/v1/subUser"
)

func (c *ControllerSubUser) SubChangeStatus(ctx context.Context, req *subUser.SubChangeStatusReq) (res *subUser.SubChangeStatusRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
