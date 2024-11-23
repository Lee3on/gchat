package model

import (
	"gchat/pkg/protocol/pb"
	"gchat/pkg/util"
	"time"
)

// Message represents a message
type Message struct {
	Id        int64     // Auto-increment primary key
	UserId    int64     // Associated user or group ID
	RequestId int64     // Request ID
	Code      int32     // Push code
	Content   []byte    // Push content
	Seq       int64     // Message synchronization sequence
	SendTime  time.Time // Message sending time
	Status    int32     // Message status
}

// MessageToPB converts the Message struct to its protobuf representation
func (m *Message) MessageToPB() *pb.Message {
	return &pb.Message{
		Code:     m.Code,
		Content:  m.Content,
		Seq:      m.Seq,
		SendTime: util.UnixMilliTime(m.SendTime),
		Status:   pb.MessageStatus(m.Status),
	}
}

// MessagesToPB converts a slice of Message structs to a slice of protobuf Message structs
func MessagesToPB(messages []Message) []*pb.Message {
	pbMessages := make([]*pb.Message, 0, len(messages))
	for i := range messages {
		pbMessage := messages[i].MessageToPB()
		if pbMessages != nil {
			pbMessages = append(pbMessages, pbMessage)
		}
	}
	return pbMessages
}
