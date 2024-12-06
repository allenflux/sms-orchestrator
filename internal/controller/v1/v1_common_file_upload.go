package v1

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"sms_backend/api/v1/common"
)

func (c *ControllerCommon) FileUpload(ctx context.Context, req *common.FileUploadReq) (res *common.FileUploadRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
