package repo

import (
	"database/sql"
	"errors"
	"gchat/pkg/db"
	"gchat/pkg/gerrors"
)

const (
	SeqObjectTypeUser = 1 // user
	SeqObjectTypeRoom = 2 // room
)

type seqRepo struct{}

var SeqRepo = new(seqRepo)

// Incr increments seq and gets the value after increment
func (*seqRepo) Incr(objectType int, objectId int64) (int64, error) {
	tx := db.DB.Begin()
	defer tx.Rollback()

	var seq int64
	err := tx.Raw("select seq from seq where object_type = ? and object_id = ? for update", objectType, objectId).
		Row().Scan(&seq)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return 0, gerrors.WrapError(err)
	}
	if errors.Is(err, sql.ErrNoRows) {
		err = tx.Exec("insert into seq (object_type,object_id,seq) values (?,?,?)", objectType, objectId, seq+1).Error
		if err != nil {
			return 0, gerrors.WrapError(err)
		}
	} else {
		err = tx.Exec("update seq set seq = seq + 1 where object_type = ? and object_id = ?", objectType, objectId).Error
		if err != nil {
			return 0, gerrors.WrapError(err)
		}
	}

	tx.Commit()
	return seq + 1, nil
}