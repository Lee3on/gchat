package device

import (
	"context"
	"gchat/pkg/gerrors"
	"gchat/pkg/protocol/pb"
)

type app struct{}

var App = new(app)

// Register register a device
func (*app) Register(ctx context.Context, in *pb.RegisterDeviceReq) (int64, error) {
	device := Device{
		Type:          in.Type,
		Brand:         in.Brand,
		Model:         in.Model,
		SystemVersion: in.SystemVersion,
		SDKVersion:    in.SdkVersion,
	}

	// check if the device is legal
	if !device.IsLegal() {
		return 0, gerrors.ErrBadRequest
	}

	err := Repo.Save(&device)
	if err != nil {
		return 0, err
	}

	return device.Id, nil
}

// SignIn sign in
func (*app) SignIn(ctx context.Context, userId, deviceId int64, token string, connAddr string, clientAddr string) error {
	return Service.SignIn(ctx, userId, deviceId, token, connAddr, clientAddr)
}

// Offline offline a device
func (*app) Offline(ctx context.Context, deviceId int64, clientAddr string) error {
	device, err := Repo.Get(deviceId)
	if err != nil {
		return err
	}
	if device == nil {
		return nil
	}

	if device.ClientAddr != clientAddr {
		return nil
	}
	device.Status = DeviceOffLine

	err = Repo.Save(device)
	if err != nil {
		return err
	}
	return nil
}

// ListOnlineByUserId list all online devices of a user
func (*app) ListOnlineByUserId(ctx context.Context, userId int64) ([]*pb.Device, error) {
	return Service.ListOnlineByUserId(ctx, userId)
}

// GetDevice get device information
func (*app) GetDevice(ctx context.Context, deviceId int64) (*pb.Device, error) {
	device, err := Repo.Get(deviceId)
	if err != nil {
		return nil, err
	}
	if device == nil {
		return nil, gerrors.ErrDeviceNotExist
	}

	return device.ToProto(), err
}

// ServerStop stop a server
func (*app) ServerStop(ctx context.Context, connAddr string) error {
	return Service.ServerStop(ctx, connAddr)
}
