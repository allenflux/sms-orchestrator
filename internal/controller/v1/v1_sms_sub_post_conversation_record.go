package v1

import (
	"context"
	"sms_backend/internal/service"

	"sms_backend/api/v1/sms"
)

func (c *ControllerSms) SubPostConversationRecord(ctx context.Context, req *sms.SubPostConversationRecordReq) (res *sms.SubPostConversationRecordRes, err error) {
	return service.SubControllerSmsManagement().PostConversationRecord(ctx, req)
}
