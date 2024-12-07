package device

import (
	"context"
	"gchat/pkg/logger"
	"gchat/pkg/protocol/pb"
	"gchat/pkg/rpc"
	"time"

	"go.uber.org/zap"
)

type service struct{}

var Service = new(service)

func (*service) Register(ctx context.Context, device *Device) error {
	err := Repo.Save(device)
	if err != nil {
		return err
	}

	return nil
}

func (*service) SignIn(ctx context.Context, userId, deviceId int64, token string, connAddr string, clientAddr string) error {
	_, err := rpc.GetUserIntClient().Auth(ctx, &pb.AuthReq{UserId: userId, DeviceId: deviceId, Token: token})
	if err != nil {
		return err
	}

	device, err := Repo.Get(deviceId)
	if err != nil {
		return err
	}
	if device == nil {
		return nil
	}

	device.Online(userId, connAddr, clientAddr)

	err = Repo.Save(device)
	if err != nil {
		return err
	}
	return nil
}

func (*service) Auth(ctx context.Context, userId, deviceId int64, token string) error {
	_, err := rpc.GetUserIntClient().Auth(ctx, &pb.AuthReq{UserId: userId, DeviceId: deviceId, Token: token})
	if err != nil {
		return err
	}
	return nil
}

func (*service) ListOnlineByUserId(ctx context.Context, userId int64) ([]*pb.Device, error) {
	devices, err := Repo.ListOnlineByUserId(userId)
	if err != nil {
		return nil, err
	}
	pbDevices := make([]*pb.Device, len(devices))
	for i := range devices {
		pbDevices[i] = devices[i].ToProto()
	}
	return pbDevices, nil
}

func (*service) ServerStop(ctx context.Context, connAddr string) error {
	devices, err := Repo.ListOnlineByConnAddr(connAddr)
	if err != nil {
		return err
	}

	for i := range devices {
		// Since the device status is modified asynchronously, avoid device reconnection, resulting in inconsistent status
		err = Repo.UpdateStatusOffline(devices[i])
		if err != nil {
			logger.Logger.Error("DeviceRepo.Save error", zap.Any("device", devices[i]), zap.Error(err))
		}
		time.Sleep(2 * time.Millisecond)
	}
	return nil
}
