package v1

import (
	"context"
	"sms_backend/internal/service"

	"sms_backend/api/v1/sms"
)

func (c *ControllerSms) SubTaskList(ctx context.Context, req *sms.SubTaskListReq) (res *sms.SubTaskListRes, err error) {
	return service.SubControllerSmsManagement().GetTaskList(ctx, req)
}
