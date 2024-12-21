package v1

import (
	"context"
	"sms_backend/internal/service"

	"sms_backend/api/v1/role"
)

func (c *ControllerRole) Updated(ctx context.Context, req *role.UpdatedReq) (res *role.UpdatedRes, err error) {
	return service.Role().UpdatedRole(ctx, req)
}
