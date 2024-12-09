package v1

import (
	"context"
	"sms_backend/internal/service"

	"sms_backend/api/v1/sms"
)

func (c *ControllerSms) ProjectUpdate(ctx context.Context, req *sms.ProjectUpdateReq) (res *sms.ProjectUpdateRes, err error) {
	return service.MainControllerProjectManagement().UpdateProject(ctx, req)
}
