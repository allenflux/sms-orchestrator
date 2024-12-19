package v1

import (
	"context"
	"sms_backend/internal/service"

	"sms_backend/api/v1/sms"
)

func (c *ControllerSms) SubGetConversationRecord(ctx context.Context, req *sms.SubGetConversationRecordReq) (res *sms.SubGetConversationRecordRes, err error) {
	return service.SubControllerSmsManagement().GetSubGetConversationRecord(ctx, req)
}
