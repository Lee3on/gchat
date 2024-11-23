package message

import (
	"context"
	"gchat/business/logic/domain/message/service"
	"gchat/pkg/protocol/pb"
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
