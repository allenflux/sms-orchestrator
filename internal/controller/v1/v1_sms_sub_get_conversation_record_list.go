package v1

import (
	"context"
	"sms_backend/internal/service"

	"sms_backend/api/v1/sms"
)

func (c *ControllerSms) SubGetConversationRecordList(ctx context.Context, req *sms.SubGetConversationRecordListReq) (res *sms.SubGetConversationRecordListRes, err error) {
	return service.SubControllerSmsManagement().SubGetConversationRecordList(ctx, req)
}
