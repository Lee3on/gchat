package room

import (
	"fmt"
	"gchat/pkg/db"
	"gchat/pkg/gerrors"
	"gchat/pkg/protocol/pb"
	"gchat/pkg/util"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"google.golang.org/protobuf/proto"
)

const MessageKey = "room_message:%d"

const MessageExpireTime = 2 * time.Minute

type messageRepo struct{}

var MessageRepo = new(messageRepo)

// Add adds the message to the queue
func (*messageRepo) Add(roomId int64, msg *pb.Message) error {
	key := fmt.Sprintf(MessageKey, roomId)
	buf, err := proto.Marshal(msg)
	if err != nil {
		return gerrors.WrapError(err)
	}
	_, err = db.RedisCli.ZAdd(key, redis.Z{
		Score:  float64(msg.Seq),
		Member: buf,
	}).Result()

	db.RedisCli.Expire(key, MessageExpireTime)
	if err != nil {
		return gerrors.WrapError(err)
	}
	return nil
}

// List retrieves messages in the specified room with a sequence number greater than seq
func (*messageRepo) List(roomId int64, seq int64) ([]*pb.Message, error) {
	key := fmt.Sprintf(MessageKey, roomId)
	result, err := db.RedisCli.ZRangeByScore(key, redis.ZRangeBy{
		Min: strconv.FormatInt(seq, 10),
		Max: "+inf",
	}).Result()
	if err != nil {
		return nil, gerrors.WrapError(err)
	}

	var msgs []*pb.Message
	for i := range result {
		buf := util.Str2bytes(result[i])
		var msg pb.Message
		err = proto.Unmarshal(buf, &msg)
		if err != nil {
			return nil, gerrors.WrapError(err)
		}
		msgs = append(msgs, &msg)
	}
	return msgs, nil
}

// ListByIndex retrieves messages in the specified room by index
func (*messageRepo) ListByIndex(roomId int64, start, stop int64) ([]*pb.Message, error) {
	key := fmt.Sprintf(MessageKey, roomId)
	result, err := db.RedisCli.ZRange(key, start, stop).Result()
	if err != nil {
		return nil, gerrors.WrapError(err)
	}

	var msgs []*pb.Message
	for i := range result {
		buf := util.Str2bytes(result[i])
		var msg pb.Message
		err = proto.Unmarshal(buf, &msg)
		if err != nil {
			return nil, gerrors.WrapError(err)
		}
		msgs = append(msgs, &msg)
	}
	return msgs, nil
}

func (*messageRepo) DelBySeq(roomId int64, min, max int64) error {
	if min == 0 && max == 0 {
		return nil
	}
	key := fmt.Sprintf(MessageKey, roomId)
	_, err := db.RedisCli.ZRemRangeByScore(key, strconv.FormatInt(min, 10), strconv.FormatInt(max, 10)).Result()
	return gerrors.WrapError(err)
}
