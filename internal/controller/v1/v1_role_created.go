package v1

import (
	"context"
	"sms_backend/internal/service"

	"sms_backend/api/v1/role"
)

func (c *ControllerRole) Created(ctx context.Context, req *role.CreatedReq) (res *role.CreatedRes, err error) {
	return service.Role().CreatedRole(ctx, req)
}
