// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package v1

import (
	"context"

	"sms_backend/api/v1/common"
	"sms_backend/api/v1/sms"
)

type IV1Common interface {
	Captcha(ctx context.Context, req *common.CaptchaReq) (res *common.CaptchaRes, err error)
	FileUpload(ctx context.Context, req *common.FileUploadReq) (res *common.FileUploadRes, err error)
	GetCountryList(ctx context.Context, req *common.GetCountryListReq) (res *common.GetCountryListRes, err error)
	ApplyCallback(ctx context.Context, req *common.ApplyCallbackReq) (res *common.ApplyCallbackRes, err error)
}

type IV1Sms interface {
	DeviceList(ctx context.Context, req *sms.DeviceListReq) (res *sms.DeviceListRes, err error)
	AllocateDevice2Project(ctx context.Context, req *sms.AllocateDevice2ProjectReq) (res *sms.AllocateDevice2ProjectRes, err error)
}
