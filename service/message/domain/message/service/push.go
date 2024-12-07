package service

import (
	"context"
	"gchat/pkg/gerrors"
	"gchat/pkg/mq"
	"gchat/pkg/protocol/pb"
	"gchat/pkg/util"
	"time"

	"google.golang.org/protobuf/proto"
)

type pushService struct{}

var PushService = new(pushService)

func (s *pushService) PushAll(ctx context.Context, req *pb.PushAllReq) error {
	msg := pb.PushAllMsg{
		Message: &pb.Message{
			Code:     req.Code,
			Content:  req.Content,
			SendTime: util.UnixMilliTime(time.Now()),
		},
	}
	bytes, err := proto.Marshal(&msg)
	if err != nil {
		return gerrors.WrapError(err)
	}
	err = mq.Publish(mq.PushAllTopic, bytes)
	if err != nil {
		return err
	}
	return nil
}
