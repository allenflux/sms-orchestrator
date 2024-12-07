// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Device is the golang structure for table device.
type Device struct {
	Id             int         `json:"id"               orm:"id"               description:""`                       //
	Name           string      `json:"name"             orm:"name"             description:""`                       //
	Model          string      `json:"model"            orm:"model"            description:""`                       //
	State          int         `json:"state"            orm:"state"            description:"1234"`                   // 1234
	GroupId        int         `json:"group_id"         orm:"group_id"         description:"ID"`                     // ID
	LeasedAt       *gtime.Time `json:"leased_at"        orm:"leased_at"        description:""`                       //
	LeaseDay       int         `json:"lease_day"        orm:"lease_day"        description:""`                       //
	LeaseEndAt     *gtime.Time `json:"lease_end_at"     orm:"lease_end_at"     description:""`                       //
	IdleState      int         `json:"idle_state"       orm:"idle_state"       description:"12"`                     // 12
	ClientId       int         `json:"client_id"        orm:"client_id"        description:"id"`                     // id
	ConnId         string      `json:"conn_id"          orm:"conn_id"          description:"websocketid"`            // websocketid
	FacebookConnId string      `json:"facebook_conn_id" orm:"facebook_conn_id" description:"facebook的websocket连接id"` // facebook的websocket连接id
	MinerConnId    string      `json:"miner_conn_id"    orm:"miner_conn_id"    description:"矿机的websocket连接id"`       // 矿机的websocket连接id
	DevId          string      `json:"dev_id"           orm:"dev_id"           description:"id"`                     // id
	LeaseOrderId   int         `json:"lease_order_id"   orm:"lease_order_id"   description:"id"`                     // id
	Ip             string      `json:"ip"               orm:"ip"               description:"ip"`                     // ip
	Remark         string      `json:"remark"           orm:"remark"           description:""`                       //
	CreatedAt      *gtime.Time `json:"created_at"       orm:"created_at"       description:""`                       //
	UpdatedAt      *gtime.Time `json:"updated_at"       orm:"updated_at"       description:""`                       //
	DeletedAt      *gtime.Time `json:"deleted_at"       orm:"deleted_at"       description:""`                       //
	IsNew          int         `json:"is_new"           orm:"is_new"           description:"12"`                     // 12
	JdUserNickname string      `json:"jd_user_nickname" orm:"jd_user_nickname" description:"金盾用户昵称"`                 // 金盾用户昵称
}
