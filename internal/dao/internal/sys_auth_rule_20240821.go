// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysAuthRule20240821Dao is the data access object for table sys_auth_rule_20240821.
type SysAuthRule20240821Dao struct {
	table   string                     // table is the underlying table name of the DAO.
	group   string                     // group is the database configuration group name of current DAO.
	columns SysAuthRule20240821Columns // columns contains all the column names of Table for convenient usage.
}

// SysAuthRule20240821Columns defines and stores column names for table sys_auth_rule_20240821.
type SysAuthRule20240821Columns struct {
	Id         string //
	Pid        string // 父ID
	Name       string // 规则名称
	Title      string // 规则名称
	Icon       string // 图标
	Condition  string // 条件
	Remark     string // 备注
	MenuType   string // 类型 0目录 1菜单 2按钮
	Weigh      string // 权重
	IsHide     string // 显示状态
	Path       string // 路由地址
	Component  string // 组件路径
	IsLink     string // 是否外链 1是 0否
	ModuleType string // 所属模块
	ModelId    string // 模型ID
	IsIframe   string // 是否内嵌iframe
	IsCached   string // 是否缓存
	Redirect   string // 路由重定向地址
	IsAffix    string // 是否固定
	LinkUrl    string // 链接地址
	CreatedAt  string // 创建日期
	UpdatedAt  string // 修改日期
}

// sysAuthRule20240821Columns holds the columns for table sys_auth_rule_20240821.
var sysAuthRule20240821Columns = SysAuthRule20240821Columns{
	Id:         "id",
	Pid:        "pid",
	Name:       "name",
	Title:      "title",
	Icon:       "icon",
	Condition:  "condition",
	Remark:     "remark",
	MenuType:   "menu_type",
	Weigh:      "weigh",
	IsHide:     "is_hide",
	Path:       "path",
	Component:  "component",
	IsLink:     "is_link",
	ModuleType: "module_type",
	ModelId:    "model_id",
	IsIframe:   "is_iframe",
	IsCached:   "is_cached",
	Redirect:   "redirect",
	IsAffix:    "is_affix",
	LinkUrl:    "link_url",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

// NewSysAuthRule20240821Dao creates and returns a new DAO object for table data access.
func NewSysAuthRule20240821Dao() *SysAuthRule20240821Dao {
	return &SysAuthRule20240821Dao{
		group:   "default",
		table:   "sys_auth_rule_20240821",
		columns: sysAuthRule20240821Columns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysAuthRule20240821Dao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysAuthRule20240821Dao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysAuthRule20240821Dao) Columns() SysAuthRule20240821Columns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysAuthRule20240821Dao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysAuthRule20240821Dao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysAuthRule20240821Dao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
