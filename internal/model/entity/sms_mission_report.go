// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SmsMissionReport is the golang structure for table sms_mission_report.
type SmsMissionReport struct {
	Id                  int         `json:"id"                    orm:"id"                    description:""`                             //
	ProjectId           int         `json:"project_id"            orm:"project_id"            description:"项目id"`                         // 项目id
	TaskName            string      `json:"task_name"             orm:"task_name"             description:"任务名称"`                         // 任务名称
	FileName            string      `json:"file_name"             orm:"file_name"             description:"文件名"`                          // 文件名
	DeviceQuota         string      `json:"device_quota"          orm:"device_quota"          description:"执行设备"`                         // 执行设备
	TaskStatus          int         `json:"task_status"           orm:"task_status"           description:"任务状态，1-待发送 2-发送中 3-已发送 4-已撤销"` // 任务状态，1-待发送 2-发送中 3-已发送 4-已撤销
	SmsQuantity         int         `json:"sms_quantity"          orm:"sms_quantity"          description:"短信总条数"`                        // 短信总条数
	SurplusQuantity     int         `json:"surplus_quantity"      orm:"surplus_quantity"      description:"剩余数量"`                         // 剩余数量
	QuantitySent        int         `json:"quantity_sent"         orm:"quantity_sent"         description:"已发送数量"`                        // 已发送数量
	AssociatedAccount   string      `json:"associated_account"    orm:"associated_account"    description:"所属子账号"`                        // 所属子账号
	IntervalTime        string      `json:"interval_time"         orm:"interval_time"         description:"间隔时间(秒)"`                      // 间隔时间(秒)
	StartTime           *gtime.Time `json:"start_time"            orm:"start_time"            description:"任务开始时间"`                       // 任务开始时间
	CreatedAt           *gtime.Time `json:"created_at"            orm:"created_at"            description:"创建时间"`                         // 创建时间
	UpdateAt            *gtime.Time `json:"update_at"             orm:"update_at"             description:"修改时间"`                         // 修改时间
	DeleteAt            *gtime.Time `json:"delete_at"             orm:"delete_at"             description:"删除时间"`                         // 删除时间
	ProjectName         string      `json:"project_name"          orm:"project_name"          description:"项目名称"`                         // 项目名称
	AssociatedAccountId int         `json:"associated_account_id" orm:"associated_account_id" description:"所属子账号id"`                      // 所属子账号id
	GroupId             int         `json:"group_id"              orm:"group_id"              description:"分组id"`                         // 分组id
}
