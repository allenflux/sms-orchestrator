package v1

import (
	"context"
	"sms_backend/internal/service"

	"sms_backend/api/v1/sms"
)

func (c *ControllerSms) SubTaskCreate(ctx context.Context, req *sms.SubTaskCreateReq) (res *sms.SubTaskCreateRes, err error) {
	return service.SubControllerSmsManagement().TaskCreate(ctx, req)
}
