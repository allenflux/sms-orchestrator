package v1

import (
	"context"
	"sms_backend/internal/service"

	"sms_backend/api/v1/sms"
)

func (c *ControllerSms) ProjectDelete(ctx context.Context, req *sms.ProjectDeleteReq) (res *sms.ProjectDeleteRes, err error) {
	return service.MainControllerProjectManagement().DeleteProject(ctx, req)
}
