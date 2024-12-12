package v1

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"sms_backend/api/v1/role"
)

func (c *ControllerRole) Created(ctx context.Context, req *role.CreatedReq) (res *role.CreatedRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
