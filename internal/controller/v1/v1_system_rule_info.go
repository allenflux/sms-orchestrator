package v1

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"sms_backend/api/v1/system"
)

func (c *ControllerSystem) RuleInfo(ctx context.Context, req *system.RuleInfoReq) (res *system.RuleInfoRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
