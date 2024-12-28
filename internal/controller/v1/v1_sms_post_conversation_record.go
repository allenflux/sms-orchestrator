package v1

import (
	"context"
	"sms_backend/internal/service"

	"sms_backend/api/v1/sms"
)

func (c *ControllerSms) PostConversationRecord(ctx context.Context, req *sms.PostConversationRecordReq) (res *sms.PostConversationRecordRes, err error) {
	return service.MainControllerSmsManagement().PostConversationRecord(ctx, req)
}
