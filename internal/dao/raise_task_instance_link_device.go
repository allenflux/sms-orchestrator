// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"sms_backend/internal/dao/internal"
)

// internalRaiseTaskInstanceLinkDeviceDao is internal type for wrapping internal DAO implements.
type internalRaiseTaskInstanceLinkDeviceDao = *internal.RaiseTaskInstanceLinkDeviceDao

// raiseTaskInstanceLinkDeviceDao is the data access object for table raise_task_instance_link_device.
// You can define custom methods on it to extend its functionality as you wish.
type raiseTaskInstanceLinkDeviceDao struct {
	internalRaiseTaskInstanceLinkDeviceDao
}

var (
	// RaiseTaskInstanceLinkDevice is globally public accessible object for table raise_task_instance_link_device operations.
	RaiseTaskInstanceLinkDevice = raiseTaskInstanceLinkDeviceDao{
		internal.NewRaiseTaskInstanceLinkDeviceDao(),
	}
)

// Fill with you ideas below.
