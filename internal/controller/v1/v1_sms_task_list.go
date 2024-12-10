package v1

import (
	"context"
	"sms_backend/internal/service"

	"sms_backend/api/v1/sms"
)

func (c *ControllerSms) TaskList(ctx context.Context, req *sms.TaskListReq) (res *sms.TaskListRes, err error) {
	return service.MainControllerSmsManagement().GetTaskList(ctx, req)
}
