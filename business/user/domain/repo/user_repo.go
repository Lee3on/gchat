package repo

import (
	"gchat/business/user/domain/model"
)

type userRepo struct{}

var UserRepo = new(userRepo)

func (*userRepo) Get(userId int64) (*model.User, error) {
	user, err := UserCache.Get(userId)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return user, nil
	}

	user, err = UserDao.Get(userId)
	if err != nil {
		return nil, err
	}

	if user != nil {
		err = UserCache.Set(*user)
		if err != nil {
			return nil, err
		}
	}
	return user, err
}

func (*userRepo) GetByPhoneNumber(phoneNumber string) (*model.User, error) {
	return UserDao.GetByPhoneNumber(phoneNumber)
}

func (*userRepo) GetByIds(userIds []int64) ([]model.User, error) {
	return UserDao.GetByIds(userIds)
}

func (*userRepo) Save(user *model.User) error {
	userId := user.Id
	err := UserDao.Save(user)
	if err != nil {
		return err
	}

	if userId != 0 {
		err = UserCache.Del(user.Id)
		if err != nil {
			return err
		}
	}
	return nil
}
