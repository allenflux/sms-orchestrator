package model

import (
	"sms_backend/internal/model/entity"
)

type DeviceGroup struct {
	entity.DeviceGroup
	Devices []*entity.Device `orm:"with:group_id=id" json:"devices"`
}
