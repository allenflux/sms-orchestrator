package v1

import (
	"context"
	"sms_backend/internal/service"

	"sms_backend/api/v1/career"
)

func (c *ControllerCareer) ReportReceiveContent(ctx context.Context, req *career.ReportReceiveContentReq) (res *career.ReportReceiveContentRes, err error) {
	return service.CareerDeviceManagement().ReportReceiveContent(ctx, req)
}
