package v1

import (
	"context"
	"sms_backend/internal/service"

	"sms_backend/api/v1/sms"
)

func (c *ControllerSms) TaskRecord(ctx context.Context, req *sms.TaskRecordReq) (res *sms.TaskRecordRes, err error) {
	return service.MainControllerSmsManagement().GetTaskRecordList(ctx, req)
}
