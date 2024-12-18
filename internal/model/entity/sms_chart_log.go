// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SmsChartLog is the golang structure for table sms_chart_log.
type SmsChartLog struct {
	Id                  int         `json:"id"                    orm:"id"                    description:""`                      //
	ProjectName         string      `json:"project_name"          orm:"project_name"          description:"项目名称"`                  // 项目名称
	ProjectId           int         `json:"project_id"            orm:"project_id"            description:"项目id"`                  // 项目id
	TargetPhoneNumber   string      `json:"target_phone_number"   orm:"target_phone_number"   description:"目标手机号"`                 // 目标手机号
	DeviceNumber        string      `json:"device_number"         orm:"device_number"         description:"执行设备号"`                 // 执行设备号
	SmsTopic            string      `json:"sms_topic"             orm:"sms_topic"             description:"短信主题"`                  // 短信主题
	SmsContent          string      `json:"sms_content"           orm:"sms_content"           description:"短信内容"`                  // 短信内容
	SmsStatus           int         `json:"sms_status"            orm:"sms_status"            description:"短信发送状态，1-失败 2-成"`       // 短信发送状态，1-失败 2-成
	AssociatedAccount   string      `json:"associated_account"    orm:"associated_account"    description:"所属子账号"`                 // 所属子账号
	CreatedAt           *gtime.Time `json:"created_at"            orm:"created_at"            description:"创建时间"`                  // 创建时间
	UpdateAt            *gtime.Time `json:"update_at"             orm:"update_at"             description:"修改时间"`                  // 修改时间
	DeleteAt            *gtime.Time `json:"delete_at"             orm:"delete_at"             description:"删除时间"`                  // 删除时间
	SentOrReceive       int         `json:"sent_or_receive"       orm:"sent_or_receive"       description:"1表示此条短信是发送 2表示此条短信是接收"` // 1表示此条短信是发送 2表示此条短信是接收
	AssociatedAccountId int         `json:"associated_account_id" orm:"associated_account_id" description:"子账号id"`                 // 子账号id
	RowHash             string      `json:"row_hash"              orm:"row_hash"              description:"log hash 防止重复上报"`       // log hash 防止重复上报
	TaskId              int         `json:"task_id"               orm:"task_id"               description:"Mission Task的id 任务 id"` // Mission Task的id 任务 id
}
