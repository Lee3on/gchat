package friend

import (
	"errors"
	"gchat/pkg/db"
	"gchat/pkg/gerrors"

	"github.com/jinzhu/gorm"
)

type repo struct{}

var Repo = new(repo)

// Get gets a friend record by userId and friendId
func (*repo) Get(userId, friendId int64) (*Friend, error) {
	friend := Friend{}
	err := db.DB.First(&friend, "user_id = ? and friend_id = ?", userId, friendId).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &friend, nil
}

// Save saves a new friend record
func (*repo) Save(friend *Friend) error {
	return gerrors.WrapError(db.DB.Save(&friend).Error)
}

// List gets friend list
func (*repo) List(userId int64, status int) ([]Friend, error) {
	var friends []Friend
	err := db.DB.Where("user_id = ? and status = ?", userId, status).Find(&friends).Error
	return friends, gerrors.WrapError(err)
}
