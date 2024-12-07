// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// NicknameImportFails is the golang structure for table nickname_import_fails.
type NicknameImportFails struct {
	Id            int         `json:"id"              orm:"id"              description:""`       //
	Nickname      string      `json:"nickname"        orm:"nickname"        description:"昵称"`     // 昵称
	TypeName      string      `json:"type_name"       orm:"type_name"       description:"昵称类型"`   // 昵称类型
	FailReason    string      `json:"fail_reason"     orm:"fail_reason"     description:"失败原因"`   // 失败原因
	DataImportsId int         `json:"data_imports_id" orm:"data_imports_id" description:"数据导入ID"` // 数据导入ID
	CreatedAt     *gtime.Time `json:"created_at"      orm:"created_at"      description:""`       //
	ClientId      int         `json:"client_id"       orm:"client_id"       description:"客户ID"`   // 客户ID
}
