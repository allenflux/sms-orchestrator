package v1

import (
	"context"
	"sms_backend/internal/service"

	"sms_backend/api/v1/sms"
)

func (c *ControllerSms) SubTaskRecord(ctx context.Context, req *sms.SubTaskRecordReq) (res *sms.SubTaskRecordRes, err error) {
	return service.SubControllerSmsManagement().GetTaskRecordList(ctx, req)
}
