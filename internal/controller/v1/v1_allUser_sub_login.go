package v1

import (
	"context"
	"sms_backend/internal/service"

	"sms_backend/api/v1/allUser"
)

func (c *ControllerAllUser) SubLogin(ctx context.Context, req *allUser.SubLoginReq) (res *allUser.SubLoginRes, err error) {
	return service.Login().HandleSubUserLogin(ctx, req)
}
