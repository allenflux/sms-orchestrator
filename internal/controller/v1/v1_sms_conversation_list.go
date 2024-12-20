package v1

import (
	"context"
	"sms_backend/internal/service"

	"sms_backend/api/v1/sms"
)

func (c *ControllerSms) ConversationList(ctx context.Context, req *sms.ConversationListReq) (res *sms.ConversationListRes, err error) {
	return service.MainControllerSmsManagement().GetConversationRecordList(ctx, req)
}
