package v1

import (
	"context"
	"sms_backend/internal/service"

	"sms_backend/api/v1/sms"
)

func (c *ControllerSms) ProjectCreate(ctx context.Context, req *sms.ProjectCreateReq) (res *sms.ProjectCreateRes, err error) {
	return service.MainControllerProjectManagement().CreateProject(ctx, req)
}
