package v1

import (
	"context"
	"sms_backend/internal/service"

	"sms_backend/api/v1/user"
)

func (c *ControllerUser) Update(ctx context.Context, req *user.UpdateReq) (res *user.UpdateRes, err error) {
	return service.User().UpdateUser(ctx, req)
}
