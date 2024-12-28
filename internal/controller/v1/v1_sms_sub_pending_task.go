package v1

import (
	"context"
	"sms_backend/internal/service"

	"sms_backend/api/v1/sms"
)

func (c *ControllerSms) SubPendingTask(ctx context.Context, req *sms.SubPendingTaskReq) (res *sms.SubPendingTaskRes, err error) {
	return service.SubControllerSmsManagement().GetPendingTasks(ctx, req)
}
