package v1

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"sms_backend/api/v1/system"
)

func (c *ControllerSystem) GetDict(ctx context.Context, req *system.GetDictReq) (res *system.GetDictRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
