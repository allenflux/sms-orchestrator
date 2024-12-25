package utility

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"os"
)

func Debug(ctx context.Context, p string, filename string) {
	fmt.Println("检查参数是否时顺序排列 ", p)
	filePath := "/home/azureuser/sms/backend/" + filename
	// 打开文件（如果不存在则创建）
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close() // 确保在函数结束时关闭文件

	// 要写入的内容
	content := p + "追加写入的内容。\n"

	// 写入文件
	if _, err := file.WriteString(content); err != nil {
		g.Log().Fatal(ctx, "Error writing to file: %v\n", err)
	}
}
