package repo

import (
	"errors"
	"gchat/pkg/db"
	"gchat/pkg/gerrors"
	"gchat/service/user/domain/group/entity"

	"github.com/jinzhu/gorm"
)

type groupDao struct{}

var GroupDao = new(groupDao)

// Get gets a group record by groupId
func (*groupDao) Get(groupId int64) (*entity.Group, error) {
	var group = entity.Group{Id: groupId}
	err := db.DB.First(&group).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, gerrors.WrapError(err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &group, nil
}

// Save saves a group record
func (*groupDao) Save(group *entity.Group) error {
	err := db.DB.Save(&group).Error
	if err != nil {
		return gerrors.WrapError(err)
	}
	return nil
}
