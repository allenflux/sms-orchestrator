package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	_ "sms_backend/internal/boot"
	"sms_backend/internal/dao"
	_ "sms_backend/internal/logic"
	_ "sms_backend/internal/packed"
	"sms_backend/utility"

	"github.com/gogf/gf/v2/os/gctx"

	"sms_backend/internal/cmd"
)

func init() {
	g.Log().SetFlags(glog.F_TIME_STD | glog.F_FILE_LONG)
	dao.Init()
}
func main() {
	ctx := gctx.GetInitCtx()
	go utility.ProcessReceivedSmsQueue(ctx, utility.ReceivedSmsPayloadChan)
	go utility.ConsumerReportTaskResult(ctx, utility.SmsTaskPayloadChan)
	go utility.ConsumerFetchTaskQueue(ctx, utility.SmsDeviceNumberPayloadChan)
	cmd.Main.Run(ctx)
}
