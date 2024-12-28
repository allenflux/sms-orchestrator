package v1

import (
	"context"
	"sms_backend/internal/service"

	"sms_backend/api/v1/sms"
)

func (c *ControllerSms) ConversationRecord(ctx context.Context, req *sms.ConversationRecordReq) (res *sms.ConversationRecordRes, err error) {
	return service.MainControllerSmsManagement().GetSubGetConversationRecord(ctx, req)
}
