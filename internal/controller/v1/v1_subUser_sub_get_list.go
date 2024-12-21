package v1

import (
	"context"
	"sms_backend/internal/service"

	"sms_backend/api/v1/subUser"
)

func (c *ControllerSubUser) SubGetList(ctx context.Context, req *subUser.SubGetListReq) (res *subUser.SubGetListRes, err error) {
	return service.SubUser().GetList(ctx, req)
}
