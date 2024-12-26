package log

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"sms_backend/api/v1/log"
	"sms_backend/internal/dao"
	"sms_backend/internal/model/entity"
	"sms_backend/internal/service"
	"sms_backend/library/liberr"
	"time"
)

type sLog struct{}

func New() *sLog { return &sLog{} }

func init() { service.RegisterLog(New()) }

func (s *sLog) GetLogList(ctx context.Context, req *log.ListReq) (res *log.ListRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		orm := dao.Log.Ctx(ctx)
		if len(req.DateRange) == 2 {
			if req.DateRange[0].String() == req.DateRange[1].String() {
				req.DateRange[1] = req.DateRange[1].Add(24 * time.Hour)
			}
			orm = orm.WhereBetween(dao.Log.Columns().CreatedAt, req.DateRange[0], req.DateRange[1])
		}
		if service.Context().GetSystemId(ctx) != 1 {
			orm = orm.Where(dao.Log.Columns().SystemId, service.Context().GetSystemId(ctx))
		}
		if req.PageNum != 0 && req.PageSize != 0 {
			orm = orm.Page(req.PageNum, req.PageSize)
		}
		if req.OrderBy != "" {
			orm = orm.Order(req.OrderBy)
		}

		res = &log.ListRes{List: make([]*entity.Log, 0)}
		err = orm.ScanAndCount(&res.List, &res.Total, false)
		liberr.ErrIsNil(ctx, err)
		if req.PageNum > 0 {
			res.Current = req.PageNum
		} else if req.PageSize == 0 {
			res.Current = 1
		}
	})
	return
}
