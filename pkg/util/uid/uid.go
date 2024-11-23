package uid

import (
	"database/sql"
	"errors"
	"log"
	"time"
)

type logger interface {
	Error(error)
}

// Logger Log interface. If a Logger is set, it will be used for logging;
// otherwise, the built-in log library will be used.
var Logger logger

// ErrTimeOut Error returned when fetching UID times out
var ErrTimeOut = errors.New("get uid timeout")

type Uid struct {
	db         *sql.DB    // Database connection
	businessId string     // Business ID
	ch         chan int64 // UID buffer pool
	min, max   int64      // Minimum and maximum values of the ID range
}

// NewUid Creates a Uid instance; len: size of the buffer pool
// db: Database connection
// businessId: Business ID
// len: Size of the buffer pool (length determines when to reload IDs from the DB)
func NewUid(db *sql.DB, businessId string, len int) (*Uid, error) {
	lid := Uid{
		db:         db,
		businessId: businessId,
		ch:         make(chan int64, len),
	}
	go lid.productId()
	return &lid, nil
}

// Get Retrieves an auto-incrementing ID. If a timeout occurs, it returns an error
// to prevent mass blocking and server crashes.
func (u *Uid) Get() (int64, error) {
	select {
	case <-time.After(1 * time.Second):
		return 0, ErrTimeOut
	case uid := <-u.ch:
		return uid, nil
	}
}

// productId Produces IDs. This method will block when the channel reaches its maximum capacity,
// until IDs in the channel are consumed.
func (u *Uid) productId() {
	_ = u.reLoad()

	for {
		if u.min >= u.max {
			_ = u.reLoad()
		}

		u.min++
		u.ch <- u.min
	}
}

// reLoad Fetches a range of IDs from the database. If it fails, it will retry every second.
func (u *Uid) reLoad() error {
	var err error
	for {
		err = u.getFromDB()
		if err == nil {
			return nil
		}

		// If a database error occurs, wait for one second and try again.
		if Logger != nil {
			Logger.Error(err)
		} else {
			log.Println(err)
		}
		time.Sleep(time.Second)
	}
}

// getFromDB Retrieves a range of IDs from the database
func (u *Uid) getFromDB() error {
	var (
		maxId int64
		step  int64
	)

	tx, err := u.db.Begin()
	if err != nil {
		return err
	}
	defer func() { _ = tx.Rollback() }()

	row := tx.QueryRow("SELECT max_id,step FROM uid WHERE business_id = ? FOR UPDATE", u.businessId)
	err = row.Scan(&maxId, &step)
	if err != nil {
		return err
	}

	_, err = tx.Exec("UPDATE uid SET max_id = ? WHERE business_id = ?", maxId+step, u.businessId)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}

	u.min = maxId
	u.max = maxId + step
	return nil
}
