package utility

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/qingwave/mossdb"
)

// create db instance
var mossPool = map[string]*mossdb.DB{}
var mossDir = "data/mossdb"

func GetMossDB(ctx context.Context) (*mossdb.DB, error) {

	_, ok := mossPool["mos"]
	if ok {
		g.Log().Info(ctx, "moss Pool中存在DB 获取已存在的 mos db")
	} else {
		g.Log().Info(ctx, "moss Pool中不存在DB 重新生成mos db ")
		db, err := mossdb.New(&mossdb.Config{Path: mossDir})
		if err != nil {
			g.Log().Errorf(ctx, "Get Moss DB Error %s", err)
			return nil, err
		}
		mossPool["mos"] = db
	}

	return mossPool["mos"], nil
}
