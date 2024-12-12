package v1

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"sms_backend/api/v1/user"
)

func (c *ControllerUser) ChangeStatus(ctx context.Context, req *user.ChangeStatusReq) (res *user.ChangeStatusRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
