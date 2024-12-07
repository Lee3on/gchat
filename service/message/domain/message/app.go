package message

import (
	"context"
	"gchat/pkg/protocol/pb"
	"gchat/service/message/domain/message/service"
)

type app struct{}

var App = new(app)

func (*app) SendToUser(ctx context.Context, fromDeviceID, toUserID int64, message *pb.Message, isPersist bool) (int64, error) {
	return service.MessageService.SendToUser(ctx, fromDeviceID, toUserID, message, isPersist)
}

func (*app) PushAll(ctx context.Context, req *pb.PushAllReq) error {
	return service.PushService.PushAll(ctx, req)
}

func (*app) Sync(ctx context.Context, userId, seq int64) (*pb.SyncResp, error) {
	return service.MessageService.Sync(ctx, userId, seq)
}

func (*app) MessageAck(ctx context.Context, userId, deviceId, ack int64) error {
	return service.DeviceAckService.Update(ctx, userId, deviceId, ack)
}
