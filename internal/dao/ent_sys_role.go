// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"sms_backend/internal/dao/internal"
)

// internalEntSysRoleDao is internal type for wrapping internal DAO implements.
type internalEntSysRoleDao = *internal.EntSysRoleDao

// entSysRoleDao is the data access object for table ent_sys_role.
// You can define custom methods on it to extend its functionality as you wish.
type entSysRoleDao struct {
	internalEntSysRoleDao
}

var (
	// EntSysRole is globally public accessible object for table ent_sys_role operations.
	EntSysRole = entSysRoleDao{
		internal.NewEntSysRoleDao(),
	}
)

// Fill with you ideas below.
