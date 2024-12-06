package cmd

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gcompress"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gfile"
	"os"
	"strings"
	"time"
)

func CompressAndClearLog(ctx context.Context) {
	go func() {
		compressAndClearLog(ctx)
		gcron.Add(ctx, "0 5 0 * * *", compressAndClearLog)
	}()
}
func compressAndClearLog(ctx context.Context) {
	logPath := g.Cfg().MustGet(ctx, "server.api.logPath").String()
	files, err := gfile.ScanDirFileFunc(logPath, `*.log`, false, func(path string) string {
		if strings.Contains(path, time.Now().Format("0102")) || strings.Contains(path, time.Now().Format("01-02")) {
			return ""
		}
		return path
	})
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	for _, file := range files {
		g.Log().Infof(ctx, "开始压缩文件：%s", file)
		err = gcompress.GzipFile(file, file+".gz", 9)
		if err != nil {
			g.Log().Error(ctx, err)
			return
		}
		g.Log().Infof(ctx, "压缩文件成功,删除原文件：%s", file)
		err = os.Remove(file)
		if err != nil {
			g.Log().Error(ctx, err)
			return
		}
	}
}
