// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AccountImportFails is the golang structure for table account_import_fails.
type AccountImportFails struct {
	Id            int         `json:"id"              orm:"id"              description:""` //
	LoginType     string      `json:"login_type"      orm:"login_type"      description:""` //
	AccountType   string      `json:"account_type"    orm:"account_type"    description:""` //
	Username      string      `json:"username"        orm:"username"        description:""` //
	Password      string      `json:"password"        orm:"password"        description:""` //
	VerifyUrl     string      `json:"verify_url"      orm:"verify_url"      description:""` //
	FailReason    string      `json:"fail_reason"     orm:"fail_reason"     description:""` //
	DataImportsId int         `json:"data_imports_id" orm:"data_imports_id" description:""` //
	CreatedAt     *gtime.Time `json:"created_at"      orm:"created_at"      description:""` //
	UpdatedAt     *gtime.Time `json:"updated_at"      orm:"updated_at"      description:""` //
	DeletedAt     *gtime.Time `json:"deleted_at"      orm:"deleted_at"      description:""` //
	Code          string      `json:"code"            orm:"code"            description:""` //
}
