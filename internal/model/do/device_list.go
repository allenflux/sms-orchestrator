// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// DeviceList is the golang structure of table device_list for DAO operations like Where/Data.
type DeviceList struct {
	g.Meta          `orm:"table:device_list, do:true"`
	Id              interface{} //
	DeviceId        interface{} // 设备id
	Number          interface{} // 设备号码
	ActiveTime      *gtime.Time // 激活时间
	OwnerAccount    interface{} // 子账号
	AssignedItems   interface{} // 所属项目
	SentStatus      interface{} // 发送状态，1-异常 2-占用 3-空闲
	QuantitySent    interface{} // 发送数量
	QuantityAll     interface{} // 全部数量
	DeviceStatus    interface{} // 设备状态，1-异常 2-正常
	GroupName       interface{} // 分组名称
	GroupId         interface{} // 分组id
	CreatedAt       *gtime.Time // 创建时间
	UpdateAt        *gtime.Time // 修改时间
	DeleteAt        *gtime.Time // 删除时间
	DeviceNumber    interface{} // 设备序列号
	StartAt         *gtime.Time // 发送时间
	AssignedItemsId interface{} // 项目id
	OwnerAccountId  interface{} // 子账号id
}
