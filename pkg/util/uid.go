package util

import (
	"database/sql"
	"gchat/pkg/logger"
	"gchat/pkg/util/uid"
)

var (
	MessageBodyIdUid *uid.Uid
	DeviceIdUid      *uid.Uid
)

const (
	DeviceIdBusinessId = "device_id"
)

func InitUID(db *sql.DB) {
	var err error
	DeviceIdUid, err = uid.NewUid(db, DeviceIdBusinessId, 5)
	if err != nil {
		logger.Sugar.Error(err)
		panic(err)
	}
}
