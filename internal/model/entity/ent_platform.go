// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// EntPlatform is the golang structure for table ent_platform.
type EntPlatform struct {
	Id           uint   `json:"id"            orm:"id"            description:""`     //
	ParentId     int    `json:"parent_id"     orm:"parent_id"     description:""`     //
	CatalogLevel int    `json:"catalog_level" orm:"catalog_level" description:"目录级别"` // 目录级别
	Name         string `json:"name"          orm:"name"          description:"名称"`   // 名称
	Nickname     string `json:"nickname"      orm:"nickname"      description:"别名"`   // 别名
	ComponentUrl string `json:"component_url" orm:"component_url" description:"组件路径"` // 组件路径
}
