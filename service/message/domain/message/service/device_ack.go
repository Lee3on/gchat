package service

import (
	"context"
	"gchat/service/message/domain/message/repo"
)

type deviceAckService struct{}

var DeviceAckService = new(deviceAckService)

// Update updates the ack of the device
func (*deviceAckService) Update(ctx context.Context, userId, deviceId, ack int64) error {
	if ack <= 0 {
		return nil
	}
	return repo.DeviceACKRepo.Set(userId, deviceId, ack)
}

// GetMaxByUserId Get the maximum ack based on the user id
func (*deviceAckService) GetMaxByUserId(ctx context.Context, userId int64) (int64, error) {
	acks, err := repo.DeviceACKRepo.Get(userId)
	if err != nil {
		return 0, err
	}

	var max int64 = 0
	for i := range acks {
		if acks[i] > max {
			max = acks[i]
		}
	}
	return max, nil
}
