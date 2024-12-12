package v1

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"sms_backend/api/v1/log"
)

func (c *ControllerLog) List(ctx context.Context, req *log.ListReq) (res *log.ListRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
