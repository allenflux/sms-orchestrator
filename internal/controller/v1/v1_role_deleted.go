package v1

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"sms_backend/api/v1/role"
)

func (c *ControllerRole) Deleted(ctx context.Context, req *role.DeletedReq) (res *role.DeletedRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
