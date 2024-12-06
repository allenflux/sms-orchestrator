package common

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type FileUploadReq struct {
	g.Meta `path:"/file/upload" mime:"multipart/form-data"  summary:"文件上传" tags:"文件管理" method:"post"`
	File   *ghttp.UploadFile `json:"file"  description:"文件"` // 文件
}

type FileUploadRes struct {
	Key string `json:"key"`
}
