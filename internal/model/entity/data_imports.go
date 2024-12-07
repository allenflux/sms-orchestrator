// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// DataImports is the golang structure for table data_imports.
type DataImports struct {
	Id           int         `json:"id"            orm:"id"            description:""`                                          //
	Type         int         `json:"type"          orm:"type"          description:"1-账户数据导入 2-昵称导入 3-授权导入 4-设备标签导入 5-app标签导入"` // 1-账户数据导入 2-昵称导入 3-授权导入 4-设备标签导入 5-app标签导入
	ClientId     int         `json:"client_id"     orm:"client_id"     description:"客户ID"`                                      // 客户ID
	CreatedAt    *gtime.Time `json:"created_at"    orm:"created_at"    description:""`                                          //
	UpdatedAt    *gtime.Time `json:"updated_at"    orm:"updated_at"    description:""`                                          //
	DeletedAt    *gtime.Time `json:"deleted_at"    orm:"deleted_at"    description:""`                                          //
	SuccessNum   int         `json:"success_num"   orm:"success_num"   description:"成功数量"`                                      // 成功数量
	FailNum      int         `json:"fail_num"      orm:"fail_num"      description:"失败数量"`                                      // 失败数量
	FileName     string      `json:"file_name"     orm:"file_name"     description:"文件名"`                                       // 文件名
	TotalNum     int         `json:"total_num"     orm:"total_num"     description:"总数量"`                                       // 总数量
	PlatformType int         `json:"platform_type" orm:"platform_type" description:"平台 1-facebook"`                             // 平台 1-facebook
	UserType     int         `json:"user_type"     orm:"user_type"     description:"1-群控 2-企业"`                                 // 1-群控 2-企业
}
