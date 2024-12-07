package v1

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"sms_backend/api/v1/sms"
)

func (c *ControllerSms) SubTaskDelete(ctx context.Context, req *sms.SubTaskDeleteReq) (res *sms.SubTaskDeleteRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
