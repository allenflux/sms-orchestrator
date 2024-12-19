// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SmsMissionRecord is the golang structure of table sms_mission_record for DAO operations like Where/Data.
type SmsMissionRecord struct {
	g.Meta              `orm:"table:sms_mission_record, do:true"`
	Id                  interface{} //
	TaskName            interface{} // 任务名称
	SubTaskId           interface{} // 子任务id
	TargetPhoneNumber   interface{} // 目标手机号
	DeviceNumber        interface{} // 执行设备号
	SmsTopic            interface{} // 短信主题
	SmsContent          interface{} // 短信内容
	SmsStatus           interface{} // 短信发送状态，2-失败 1-成功
	AssociatedAccount   interface{} // 所属子账号
	ProjectName         interface{} // 所属项目名称
	ProjectId           interface{} // 所属项目id
	StartTime           *gtime.Time // 开始时间
	CreatedAt           *gtime.Time // 创建时间
	UpdateAt            *gtime.Time // 修改时间
	DeleteAt            *gtime.Time // 删除时间
	AssociatedAccountId interface{} // 子账户id
	RowHash             interface{} // 每行内容的hash串 为了防止同样的记录重复提交
}
