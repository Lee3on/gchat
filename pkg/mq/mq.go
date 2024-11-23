package mq

import (
	"gchat/pkg/db"
	"gchat/pkg/gerrors"
)

const (
	PushRoomTopic         = "push_room_topic"          // room message queue
	PushRoomPriorityTopic = "push_room_priority_topic" // room priority message queue
	PushAllTopic          = "push_all_topic"           // all rooms message queue
)

func Publish(topic string, bytes []byte) error {
	_, err := db.RedisCli.Publish(topic, bytes).Result()
	if err != nil {
		return gerrors.WrapError(err)
	}
	return nil
}
