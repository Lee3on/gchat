package repo

import (
	"errors"
	"gchat/business/user/domain/model"
	"gchat/pkg/db"
	"gchat/pkg/gerrors"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

const (
	UserKey    = "user:"
	UserExpire = 2 * time.Hour
)

type userCache struct{}

var UserCache = new(userCache)

func (c *userCache) Get(userId int64) (*model.User, error) {
	var user model.User
	err := db.RedisUtil.Get(UserKey+strconv.FormatInt(userId, 10), &user)
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, gerrors.WrapError(err)
	}
	if errors.Is(err, redis.Nil) {
		return nil, nil
	}
	return &user, nil
}

func (c *userCache) Set(user model.User) error {
	err := db.RedisUtil.Set(UserKey+strconv.FormatInt(user.Id, 10), user, UserExpire)
	if err != nil {
		return gerrors.WrapError(err)
	}
	return nil
}

func (c *userCache) Del(userId int64) error {
	_, err := db.RedisCli.Del(UserKey + strconv.FormatInt(userId, 10)).Result()
	if err != nil {
		return gerrors.WrapError(err)
	}
	return nil
}
