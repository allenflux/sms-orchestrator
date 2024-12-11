// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SmsChartLog is the golang structure of table sms_chart_log for DAO operations like Where/Data.
type SmsChartLog struct {
	g.Meta            `orm:"table:sms_chart_log, do:true"`
	Id                interface{} //
	ProjectName       interface{} // 项目名称
	ProjectId         interface{} // 项目id
	TargetPhoneNumber interface{} // 目标手机号
	DeviceNumber      interface{} // 执行设备号
	SmsTopic          interface{} // 短信主题
	SmsContent        interface{} // 短信内容
	SmsStatus         interface{} // 短信发送状态，1-失败 2-成
	AssociatedAccount interface{} // 所属子账号
	CreatedAt         *gtime.Time // 创建时间
	UpdateAt          *gtime.Time // 修改时间
	DeleteAt          *gtime.Time // 删除时间
}
