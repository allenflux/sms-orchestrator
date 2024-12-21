package v1

import (
	"context"
	"sms_backend/internal/service"

	"sms_backend/api/v1/user"
)

func (c *ControllerUser) Delete(ctx context.Context, req *user.DeleteReq) (res *user.DeleteRes, err error) {
	return service.User().DeleteUser(ctx, req)
}
