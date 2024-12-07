// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// EntSysRoleDept is the golang structure for table ent_sys_role_dept.
type EntSysRoleDept struct {
	RoleId int64 `json:"role_id" orm:"role_id" description:"角色ID"` // 角色ID
	DeptId int64 `json:"dept_id" orm:"dept_id" description:"部门ID"` // 部门ID
}
