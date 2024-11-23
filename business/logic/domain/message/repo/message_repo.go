package repo

import (
	"fmt"
	"gchat/business/logic/domain/message/model"
	"gchat/pkg/db"
	"gchat/pkg/gerrors"
)

type messageRepo struct{}

var MessageRepo = new(messageRepo)

func (*messageRepo) tableName(userId int64) string {
	return fmt.Sprintf("message")
}

// Save saves a message record
func (d *messageRepo) Save(message model.Message) error {
	err := db.DB.Table(d.tableName(message.UserId)).Create(&message).Error
	if err != nil {
		return gerrors.WrapError(err)
	}
	return nil
}

// ListBySeq list messages with seq number greater than seq
func (d *messageRepo) ListBySeq(userId, seq, limit int64) ([]model.Message, bool, error) {
	DB := db.DB.Table(d.tableName(userId)).
		Where("user_id = ? and seq > ?", userId, seq)

	var count int64
	err := DB.Count(&count).Error
	if err != nil {
		return nil, false, gerrors.WrapError(err)
	}
	if count == 0 {
		return nil, false, nil
	}

	var messages []model.Message
	err = DB.Limit(limit).Find(&messages).Error
	if err != nil {
		return nil, false, gerrors.WrapError(err)
	}
	return messages, count > limit, nil
}
