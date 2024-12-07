package device

import (
	"gchat/pkg/protocol/pb"
	"time"
)

const (
	DeviceOnLine  = 1 // Device is online
	DeviceOffLine = 0 // Device is offline
)

type Device struct {
	Id            int64     // Device ID
	UserId        int64     // User ID
	Type          int32     // Device type: 1=Android; 2=IOS; 3=Windows; 4=MacOS; 5=Web
	Brand         string    // Manufacturer
	Model         string    // Model
	SystemVersion string    // System version
	SDKVersion    string    // SDK version
	Status        int32     // Online status: 0=Offline; 1=Online
	ConnAddr      string    // Connection layer service address
	ClientAddr    string    // Client address
	CreateTime    time.Time // Creation time
	UpdateTime    time.Time // Update time
}

func (d *Device) ToProto() *pb.Device {
	return &pb.Device{
		DeviceId:      d.Id,
		UserId:        d.UserId,
		Type:          d.Type,
		Brand:         d.Brand,
		Model:         d.Model,
		SystemVersion: d.SystemVersion,
		SdkVersion:    d.SDKVersion,
		Status:        d.Status,
		ConnAddr:      d.ConnAddr,
		ClientAddr:    d.ClientAddr,
		CreateTime:    d.CreateTime.Unix(),
		UpdateTime:    d.UpdateTime.Unix(),
	}
}

func (d *Device) IsLegal() bool {
	if d.Type == 0 || d.Brand == "" || d.Model == "" ||
		d.SystemVersion == "" || d.SDKVersion == "" {
		return false
	}
	return true
}

func (d *Device) Online(userId int64, connAddr string, clientAddr string) {
	d.UserId = userId
	d.ConnAddr = connAddr
	d.ClientAddr = clientAddr
	d.Status = DeviceOnLine
}

func (d *Device) Offline(userId int64, connAddr string, clientAddr string) {
	d.UserId = userId
	d.ConnAddr = connAddr
	d.ClientAddr = clientAddr
	d.Status = DeviceOffLine
}
