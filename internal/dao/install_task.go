// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"sms_backend/internal/dao/internal"
)

// internalInstallTaskDao is internal type for wrapping internal DAO implements.
type internalInstallTaskDao = *internal.InstallTaskDao

// installTaskDao is the data access object for table install_task.
// You can define custom methods on it to extend its functionality as you wish.
type installTaskDao struct {
	internalInstallTaskDao
}

var (
	// InstallTask is globally public accessible object for table install_task operations.
	InstallTask = installTaskDao{
		internal.NewInstallTaskDao(),
	}
)

// Fill with you ideas below.
