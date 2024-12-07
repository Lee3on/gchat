package device

import (
	"errors"
	"gchat/pkg/db"
	"gchat/pkg/gerrors"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

const (
	UserDeviceKey    = "user_device:"
	UserDeviceExpire = 2 * time.Hour
)

type userDeviceCache struct{}

var UserDeviceCache = new(userDeviceCache)

// Get get all online devices of a user
func (c *userDeviceCache) Get(userId int64) ([]Device, error) {
	var devices []Device
	err := db.RedisUtil.Get(UserDeviceKey+strconv.FormatInt(userId, 10), &devices)
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, gerrors.WrapError(err)
	}

	if errors.Is(err, redis.Nil) {
		return nil, nil
	}
	return devices, nil
}

// Set save all online devices of a user to cache
func (c *userDeviceCache) Set(userId int64, devices []Device) error {
	err := db.RedisUtil.Set(UserDeviceKey+strconv.FormatInt(userId, 10), devices, UserDeviceExpire)
	return gerrors.WrapError(err)
}

// Del delete the online device list of a user
func (c *userDeviceCache) Del(userId int64) error {
	key := UserDeviceKey + strconv.FormatInt(userId, 10)
	_, err := db.RedisCli.Del(key).Result()
	return gerrors.WrapError(err)
}
