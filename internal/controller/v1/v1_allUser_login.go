package v1

import (
	"context"
	"sms_backend/internal/service"

	"sms_backend/api/v1/allUser"
)

func (c *ControllerAllUser) Login(ctx context.Context, req *allUser.LoginReq) (res *allUser.LoginRes, err error) {
	return service.Login().Login(ctx, req)
}
