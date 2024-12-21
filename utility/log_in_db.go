package utility

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"sms_backend/internal/model/do"
)

func CreatedLog(ctx context.Context, data do.Log) (err error) {
	_, err = g.DB().Model("log").Ctx(ctx).Data(data).Insert()
	return err
}
