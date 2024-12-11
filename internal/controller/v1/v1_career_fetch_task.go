package v1

import (
	"context"
	"sms_backend/internal/service"

	"sms_backend/api/v1/career"
)

func (c *ControllerCareer) FetchTask(ctx context.Context, req *career.FetchTaskReq) (res *career.FetchTaskRes, err error) {
	return service.CareerDeviceManagement().FetchTasks(ctx, req)
}
