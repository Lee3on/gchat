package device

import (
	"errors"
	"gchat/business/user/domain/model"
	"gchat/pkg/db"
	"gchat/pkg/gerrors"
	"time"

	"github.com/jinzhu/gorm"
)

type dao struct{}

var Dao = new(dao)

// Save save a new device record
func (*dao) Save(device *Device) error {
	device.CreateTime = time.Now()
	device.UpdateTime = time.Now()
	err := db.DB.Save(&device).Error
	if err != nil {
		return gerrors.WrapError(err)
	}
	return nil
}

// Get get a device record by deviceId
func (*dao) Get(deviceId int64) (*Device, error) {
	var device = Device{Id: deviceId}
	err := db.DB.First(&device).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, gerrors.WrapError(err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &device, nil
}

// ListOnlineByUserId query all online devices of a user
func (*dao) ListOnlineByUserId(userId int64) ([]Device, error) {
	var devices []Device
	err := db.DB.Find(&devices, "user_id = ? and status = ?", userId, DeviceOnLine).Error
	if err != nil {
		return nil, gerrors.WrapError(err)
	}
	return devices, nil
}

// ListOnlineByConnAddr query all online devices of a connection address
func (*dao) ListOnlineByConnAddr(connAddr string) ([]Device, error) {
	var devices []Device
	err := db.DB.Find(&devices, "conn_addr = ? and status = ?", connAddr, DeviceOnLine).Error
	if err != nil {
		return nil, gerrors.WrapError(err)
	}
	return devices, nil
}

// UpdateStatus update the online status of a device
func (*dao) UpdateStatus(deviceId int64, connAddr string, status int) (int64, error) {
	db := db.DB.Model(&model.Device{}).Where("id = ? and conn_addr = ?", deviceId, connAddr).
		Update("status", status)
	if db.Error != nil {
		return 0, gerrors.WrapError(db.Error)
	}
	return db.RowsAffected, nil
}
