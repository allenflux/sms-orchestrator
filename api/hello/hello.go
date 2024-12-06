// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package hello

import (
	"context"
	"sms_backend/api/v1"
)

type IHelloV1 interface {
	Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error)
}
