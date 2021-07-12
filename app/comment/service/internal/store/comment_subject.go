package store

import (
	"time"

	"github.com/jmoiron/sqlx"
)

const (
	_incrSubCntSQL = "UPDATE reply_subject_1 SET count=count+1,root_count=root_count+1,all_count=all_acount+1,update_time=? WHERE obj_id=? AND obj_type=?"
)

// TxIncrSubjectCount incr subject count and root count by transaction.
func (s *Store) TxIncrSubjectCount(tx *sqlx.Tx, oid int64, tp int, now time.Time) (rows int64, err error) {
	res, err := tx.Exec(_incrSubCntSQL, now, oid, tp)
	if err != nil {
		return
	}
	return res.RowsAffected()
}

// TxIncrSubjectRootCount incr subject count and root count by transaction.
func (s *Store) TxIncrSubjectRootCount(tx *sqlx.Tx, oid int64, tp int, now time.Time) (rows int64, err error) {
	res, err := tx.Exec(_incrSubRootCntSQL, now, oid, tp)
	if err != nil {
		return
	}
	return res.RowsAffected()
}
