// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"sms_backend/internal/dao/internal"
)

// internalMinerImportMobileTaskLinkDeviceDao is internal type for wrapping internal DAO implements.
type internalMinerImportMobileTaskLinkDeviceDao = *internal.MinerImportMobileTaskLinkDeviceDao

// minerImportMobileTaskLinkDeviceDao is the data access object for table miner_import_mobile_task_link_device.
// You can define custom methods on it to extend its functionality as you wish.
type minerImportMobileTaskLinkDeviceDao struct {
	internalMinerImportMobileTaskLinkDeviceDao
}

var (
	// MinerImportMobileTaskLinkDevice is globally public accessible object for table miner_import_mobile_task_link_device operations.
	MinerImportMobileTaskLinkDevice = minerImportMobileTaskLinkDeviceDao{
		internal.NewMinerImportMobileTaskLinkDeviceDao(),
	}
)

// Fill with you ideas below.
