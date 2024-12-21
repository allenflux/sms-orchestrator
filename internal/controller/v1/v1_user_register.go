package v1

import (
	"context"
	"sms_backend/internal/service"

	"sms_backend/api/v1/user"
)

func (c *ControllerUser) Register(ctx context.Context, req *user.RegisterReq) (res *user.RegisterRes, err error) {
	return service.User().CreatedUser(ctx, req)
}
