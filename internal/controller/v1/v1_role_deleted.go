package v1

import (
	"context"
	"sms_backend/internal/service"

	"sms_backend/api/v1/role"
)

func (c *ControllerRole) Deleted(ctx context.Context, req *role.DeletedReq) (res *role.DeletedRes, err error) {
	return service.Role().DeleteRole(ctx, req)
}
