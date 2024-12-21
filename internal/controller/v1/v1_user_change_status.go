package v1

import (
	"context"
	"sms_backend/internal/service"

	"sms_backend/api/v1/user"
)

func (c *ControllerUser) ChangeStatus(ctx context.Context, req *user.ChangeStatusReq) (res *user.ChangeStatusRes, err error) {
	return service.User().ChangeStatus(ctx, req)
}
