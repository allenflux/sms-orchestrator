// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// MinerCommunity is the golang structure for table miner_community.
type MinerCommunity struct {
	Id          uint   `json:"id"            orm:"id"            description:""`     //
	DeviceDevId string `json:"device_dev_id" orm:"device_dev_id" description:"设备id"` // 设备id
	Community   string `json:"community"     orm:"community"     description:"社区"`   // 社区
}
