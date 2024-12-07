// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// MinerPlatform is the golang structure for table miner_platform.
type MinerPlatform struct {
	Id       uint   `json:"id"       orm:"id"       description:""`      //
	Platform string `json:"platform" orm:"platform" description:"支持的平台"` // 支持的平台
}
