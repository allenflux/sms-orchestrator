package v1

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"sms_backend/api/v1/career"
)

func (c *ControllerCareer) ReportInfo(ctx context.Context, req *career.ReportInfoReq) (res *career.ReportInfoRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
