package main

import (
	_ "sms_backend/internal/packed"

	_ "sms_backend/internal/logic"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"github.com/gogf/gf/v2/os/gctx"

	"sms_backend/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
