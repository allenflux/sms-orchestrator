package v1

import (
	"context"
	"sms_backend/internal/service"

	"sms_backend/api/v1/career"
)

func (c *ControllerCareer) ReportTaskResult(ctx context.Context, req *career.ReportTaskResultReq) (res *career.ReportTaskResultRes, err error) {
	return service.CareerDeviceManagement().ReportTaskResult(ctx, req)
}
