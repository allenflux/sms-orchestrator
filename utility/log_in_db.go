package utility

import (
	"context"
	"sms_backend/internal/dao"
	"sms_backend/internal/model/do"
)

func CreatedLog(ctx context.Context, data do.Log) (err error) {
	_, err = dao.Log.Ctx(ctx).Data(data).Insert()
	return err
}
