package v1

import (
	"context"
	"sms_backend/internal/service"

	"sms_backend/api/v1/log"
)

func (c *ControllerLog) List(ctx context.Context, req *log.ListReq) (res *log.ListRes, err error) {
	return service.Log().GetLogList(ctx, req)
}
