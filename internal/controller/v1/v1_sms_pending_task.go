package v1

import (
	"context"
	"sms_backend/internal/service"

	"sms_backend/api/v1/sms"
)

func (c *ControllerSms) PendingTask(ctx context.Context, req *sms.PendingTaskReq) (res *sms.PendingTaskRes, err error) {
	return service.MainControllerSmsManagement().GetPendingTasks(ctx, req)
}
