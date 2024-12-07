package v1

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"sms_backend/api/v1/system"
)

func (c *ControllerSystem) RuleGetParams(ctx context.Context, req *system.RuleGetParamsReq) (res *system.RuleGetParamsRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
