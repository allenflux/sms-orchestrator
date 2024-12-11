// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SmsMissionReport is the golang structure of table sms_mission_report for DAO operations like Where/Data.
type SmsMissionReport struct {
	g.Meta              `orm:"table:sms_mission_report, do:true"`
	Id                  interface{} //
	ProjectId           interface{} // 项目id
	TaskName            interface{} // 任务名称
	FileName            interface{} // 文件名
	DeviceQuota         interface{} // 执行设备
	TaskStatus          interface{} // 任务状态，1-待发送 2-发送中 3-已发送 4-已撤销
	SmsQuantity         interface{} // 短信总条数
	SurplusQuantity     interface{} // 剩余数量
	QuantitySent        interface{} // 已发送数量
	AssociatedAccount   interface{} // 所属子账号
	IntervalTime        interface{} // 间隔时间(秒)
	StartTime           *gtime.Time // 任务开始时间
	CreatedAt           *gtime.Time // 创建时间
	UpdateAt            *gtime.Time // 修改时间
	DeleteAt            *gtime.Time // 删除时间
	ProjectName         interface{} // 项目名称
	AssociatedAccountId interface{} // 所属子账号id
	GroupId             interface{} // 分组id
}
