// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// EntSysConfig is the golang structure for table ent_sys_config.
type EntSysConfig struct {
	ConfigId    int         `json:"config_id"    orm:"config_id"    description:"参数主键"`        // 参数主键
	ConfigName  string      `json:"config_name"  orm:"config_name"  description:"参数名称"`        // 参数名称
	ConfigKey   string      `json:"config_key"   orm:"config_key"   description:"参数键名"`        // 参数键名
	ConfigValue string      `json:"config_value" orm:"config_value" description:"参数键值"`        // 参数键值
	ConfigType  int         `json:"config_type"  orm:"config_type"  description:"系统内置（Y是 N否）"` // 系统内置（Y是 N否）
	CreateBy    int         `json:"create_by"    orm:"create_by"    description:"创建者"`         // 创建者
	UpdateBy    int         `json:"update_by"    orm:"update_by"    description:"更新者"`         // 更新者
	Remark      string      `json:"remark"       orm:"remark"       description:"备注"`          // 备注
	CreatedAt   *gtime.Time `json:"created_at"   orm:"created_at"   description:"创建时间"`        // 创建时间
	UpdatedAt   *gtime.Time `json:"updated_at"   orm:"updated_at"   description:"修改时间"`        // 修改时间
}
