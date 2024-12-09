package v1

import (
	"context"
	"sms_backend/internal/service"

	"sms_backend/api/v1/sms"
)

func (c *ControllerSms) ProjectList(ctx context.Context, req *sms.ProjectListReq) (res *sms.ProjectListRes, err error) {
	return service.MainControllerProjectManagement().ProjectList(ctx, req)
}
