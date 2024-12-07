// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// EntSysOperLog is the golang structure for table ent_sys_oper_log.
type EntSysOperLog struct {
	OperId        int64       `json:"oper_id"        orm:"oper_id"        description:"日志主键"`                   // 日志主键
	Title         string      `json:"title"          orm:"title"          description:"模块标题"`                   // 模块标题
	BusinessType  int         `json:"business_type"  orm:"business_type"  description:"业务类型（0其它 1新增 2修改 3删除）"`  // 业务类型（0其它 1新增 2修改 3删除）
	Method        string      `json:"method"         orm:"method"         description:"方法名称"`                   // 方法名称
	RequestMethod string      `json:"request_method" orm:"request_method" description:"请求方式"`                   // 请求方式
	OperatorType  int         `json:"operator_type"  orm:"operator_type"  description:"操作类别（0其它 1后台用户 2手机端用户）"` // 操作类别（0其它 1后台用户 2手机端用户）
	OperName      string      `json:"oper_name"      orm:"oper_name"      description:"操作人员"`                   // 操作人员
	DeptName      string      `json:"dept_name"      orm:"dept_name"      description:"部门名称"`                   // 部门名称
	OperUrl       string      `json:"oper_url"       orm:"oper_url"       description:"请求URL"`                  // 请求URL
	OperIp        string      `json:"oper_ip"        orm:"oper_ip"        description:"主机地址"`                   // 主机地址
	OperLocation  string      `json:"oper_location"  orm:"oper_location"  description:"操作地点"`                   // 操作地点
	OperParam     string      `json:"oper_param"     orm:"oper_param"     description:"请求参数"`                   // 请求参数
	ErrorMsg      string      `json:"error_msg"      orm:"error_msg"      description:"错误消息"`                   // 错误消息
	OperTime      *gtime.Time `json:"oper_time"      orm:"oper_time"      description:"操作时间"`                   // 操作时间
}
