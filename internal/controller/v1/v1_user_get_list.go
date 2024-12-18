package v1

import (
	"context"
	"sms_backend/internal/service"

	"sms_backend/api/v1/user"
)

func (c *ControllerUser) GetList(ctx context.Context, req *user.GetListReq) (res *user.GetListRes, err error) {
	return service.User().GetList(ctx, req)
}
