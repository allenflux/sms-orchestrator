package v1

import (
	"context"
	"sms_backend/internal/service"

	"sms_backend/api/v1/role"
)

func (c *ControllerRole) List(ctx context.Context, req *role.ListReq) (res *role.ListRes, err error) {
	return service.Role().GetList(ctx, req)
}
