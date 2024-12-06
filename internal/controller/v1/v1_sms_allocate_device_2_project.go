package v1

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"sms_backend/api/v1/sms"
)

func (c *ControllerSms) AllocateDevice2Project(ctx context.Context, req *sms.AllocateDevice2ProjectReq) (res *sms.AllocateDevice2ProjectRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
