package v1

import (
	"context"
	"sms_backend/internal/service"

	"sms_backend/api/v1/role"
)

func (c *ControllerRole) GetPermission(ctx context.Context, req *role.GetPermissionReq) (res *role.GetPermissionRes, err error) {
	return service.Role().GetPermissionList(ctx, req)
}
