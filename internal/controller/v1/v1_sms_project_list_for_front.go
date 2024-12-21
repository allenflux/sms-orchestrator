package v1

import (
	"context"
	"sms_backend/internal/service"

	"sms_backend/api/v1/sms"
)

func (c *ControllerSms) ProjectListForFront(ctx context.Context, req *sms.ProjectListForFrontReq) (res *sms.ProjectListForFrontRes, err error) {
	return service.MainControllerProjectManagement().ProjectListListForFront(ctx, req)
}
