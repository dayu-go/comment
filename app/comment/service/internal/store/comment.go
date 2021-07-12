package store

import (
	"context"
	"database/sql"
	"time"

	"github.com/dayu-go/comment/app/comment/service/internal/config"
	"github.com/jmoiron/sqlx"
)

const (
	_inCommentSubjectSQL = "INSERT INTO comment_subject_1(obj_id, obj_type, user_id, create_time, update_time) VALUES (?,?,?,?,?);"
	_inCommentIndexSQL   = "INSERT INTO comment_content_1(comment_id, ip, platform, device, message, meta, create_time, update_time) VALUES (?,?,?,?,?,?,?,?);"
	_inCommentContentSQL = "INSERT INTO comment_content_1(comment_id, ip, platform, device, message, meta, create_time, update_time) VALUES (?,?,?,?,?,?,?,?);"
	_selCommentSubSQL    = "SELECT id, obj_id, obj_type, member_id, count, root_count, all_count, state, attrs, create_time, update_time FROM comment_subject_1 WHERE id = ?"
	_existsCommentSubSQL = "SELECT id FROM comment_index_1 WHERE id = ?"
)

// CommentSubject 评论主题表
type CommentSubject struct {
	Id         int64  `db:"id"`
	ObjId      int64  `db:"obj_id"`
	ObjType    int    `db:"obj_type"`
	MemberId   int64  `db:"member_id"`
	Count      int    `db:"count"`
	RootCount  int    `db:"root_count"`
	AllCount   int    `db:"all_count"`
	State      int    `db:"state"` // 状态（0-正常，1-隐藏）
	Attrs      int    `db:"attrs"` // 属性（bit 0-运营置顶，1-up置顶，2-大数据过滤）
	CreateTime string `db:"create_time"`
	UpdateTime string `db:"update_time"`
}

// 索引表
type CommentIndex struct {
	Id         int64  `db:"id"`
	ObjId      int64  `db:"obj_id"`
	ObjType    int    `db:"obj_type"`
	UserId     int64  `db:"user_id"`
	Root       int64  `db:"root"`
	Parent     int64  `db:"parent"`
	Floor      int    `db:"floor"`
	Count      int    `db:"count"`      // 评论总数(回复总数)
	RootCount  int    `db:"root_count"` // 根评论总数
	Like       int    `db:"like"`
	Hate       int    `db:"hate"`
	State      int    `db:"state"` // 状态（0-正常，1-隐藏）
	Attrs      int    `db:"attrs"`
	CreateTime string `db:"create_time"`
	UpdateTime string `db:"update_time"`
}

// 内容表
type CommentContent struct {
	CommentId  int64  `db:"comment_id"`
	IP         string `db:"ip"`
	Platform   string `db:"platform"`
	Device     string `db:"device"`
	Message    string `db:"message"`
	Meta       string `db:"meta"`
	CreateTime string `db:"create_time"`
	UpdateTime string `db:"update_time"`
}

type Comment struct {
	Id         int64  `db:"id"`
	ObjId      int64  `db:"obj_id"`
	ObjType    int    `db:"obj_type"`
	UserId     int64  `db:"user_id"`
	Root       int64  `db:"root"`
	Floor      int    `db:"floor"`
	Count      int    `db:"count"`
	RootCount  int    `db:"root_count"`
	Like       int    `db:"like"`
	Hate       int    `db:"hate"`
	State      int    `db:"state"` // 状态（0-正常，1-隐藏）
	Attrs      int    `db:"attrs"`
	Message    string `db:"message"`
	Meta       string `db:"meta"`
	CreateTime string `db:"create_time"`
	UpdateTime string `db:"update_time"`
}

type CreateCommentRequest struct {
	ObjId      int64
	ObjType    int
	ObjUserId  int64
	UserId     int64
	Count      int   // 评论总数
	Root       int64 // 根评论id，不为0是回复评论
	Floor      int   // 评论楼层
	ReplyCount int   // 评论回复总数
	RootCount  int   // 根评论总数
	IP         string
	Platform   int
	Device     string
	Meta       string
	Message    string
	CreateTime string
	UpdateTime string
}

func (s *Store) CreateCommentSubject(ctx context.Context, req CreateCommentRequest) (resp CommentSubject, err error) {
	err = BeginTX(s.Db, func(t *sqlx.Tx) error {
		ok, subject, err := s.FindCommentSubject(ctx, req.ObjId, req.ObjType)
		if err != nil {
			return err
		}
		if ok {
			sql := `UPDATE comment_subject_1 SET state = 1, update_time = ? WHERE id = ?;`
			_, err := t.Exec(sql, req.UpdateTime, subject.Id)
			if err != nil {
				return err
			}
		}
		sql := `INSERT INTO comment_subject_1(obj_id, obj_type, user_id, create_time, update_time) VALUES (?,?,?,?,?);`
		_, err = t.Exec(sql, req.ObjId, req.ObjType, req.UserId, req.CreateTime, req.UpdateTime)
		return err
	})
	return
}

func (s *Store) FindCommentSubject(ctx context.Context, objId int64, objType int) (ok bool, sub CommentSubject, err error) {
	if err = s.Db.Get(&sub, _selCommentSubSQL, objId, objType); err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return
		}
		return
	}
	ok = true
	return
}

func (s *Store) ExistsCommentIndex(ctx context.Context, id int64) (ok bool, err error) {
	var rootID int64
	if err = s.Db.Get(&rootID, _existsCommentSubSQL, id); err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return
		}
		return
	}
	ok = rootID > 0
	return
}

func (s *Store) FindComment(ctx context.Context, id int64) (ok bool, cmt Comment, err error) {
	query := `SELECT
	a.id,
	a.obj_id,
	a.obj_type,
	a.user_id,
	a.root,
	a.floor,
	a.count,
	a.root_count,
	a.LIKE,
	a.hate,
	a.state,
	a.attrs,
	a.create_time,
	a.update_time,
	b.message,
	b.meta 
FROM
	comment_index_1 a,
	comment_content_1 b 
WHERE
	a.id = b.comment_id 
	AND a.id = ?;`
	if err = s.Db.Get(&cmt, query, id); err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return
		}
		return
	}
	ok = true
	return
}

func (s *Store) CreateComment(ctx context.Context, req CreateCommentRequest) (newId int64, err error) {
	BeginTX(s.Db, func(t *sqlx.Tx) error {
		// find comment subject
		ok, _, err := s.FindCommentSubject(ctx, req.ObjId, req.ObjType)
		if err != nil {
			return err
		}
		if !ok {
			_, err := t.Exec(_inCommentSubjectSQL, req.ObjId, req.ObjType, req.ObjUserId, req.CreateTime, req.UpdateTime)
			if err != nil {
				return err
			}
		}

		res, err := t.Exec(_inCommentIndexSQL, req.ObjId, req.ObjType, req.UserId, req.Root, req.Floor, req.RootCount, req.CreateTime, req.UpdateTime)
		if err != nil {
			return err
		}
		newId, err = res.LastInsertId()
		if err != nil {
			return err
		}
		_, err = t.Exec(_inCommentContentSQL, newId, req.IP, req.Platform, req.Device, req.Message, req.Meta, req.CreateTime, req.UpdateTime)
		return err
	})
	return
}

func (s *Store) GetComment(ctx context.Context, id int64) (ok bool, cmt Comment, err error) {
	sqlStr := `SELECT id, obj_id, obj_type FROM comment_index_1 WHERE id = ?;`
	err = s.Db.Get(&cmt, sqlStr, 1)
	if err == sql.ErrNoRows {
		return false, cmt, nil
	}
	return
}

func (s *Store) ListComment(c config.DBConfig) (com Comment, err error) {

	return
}

// TxIncrSubjectRootCount incr subject count and root count by transaction.
func (s *Store) TxIncrRootCount(tx *sqlx.Tx, oid int64, tp int, now time.Time) (rows int64, err error) {
	res, err := tx.Exec(_incrSubRootCntSQL, now, oid, tp)
	if err != nil {
		return
	}
	return res.RowsAffected()
}
