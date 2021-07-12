package store

import (
	"context"
	"database/sql"
	"log"
	"testing"
	"time"

	"github.com/dayu-go/comment/app/comment/service/internal/config"
	"github.com/jmoiron/sqlx"
)

// go test -v *.go -test.run=TestGetComment
func TestCreateComment(t *testing.T) {
	s := NewStore().NewDB(config.DBConfig{
		Driver: "mysql",
		DSN:    "root:123456@tcp(127.0.0.1:3306)/dayu?parseTime=true&loc=Local",
	})
	now := time.Now().Format("2006-01-02 15:04:05")
	c, err := s.CreateComment(context.Background(), CreateCommentRequest{
		ObjId:      1,
		ObjType:    1,
		ObjUserId:  100001,
		UserId:     100002,
		ReplyCount: 0,
		RootCount:  0,
		IP:         "192.168.0.1",
		Platform:   1,
		Device:     "devicexxx",
		Meta:       "color:red",
		Message:    "first comment",
		CreateTime: now,
		UpdateTime: now,
	})
	t.Logf("c:%+v, err:%v", c, err)
}

// go test -v *.go -test.run=TestGetComment
func TestGetComment(t *testing.T) {
	s := NewStore().NewDB(config.DBConfig{
		Driver: "mysql",
		DSN:    "root:123456@tcp(127.0.0.1:3306)/dayu?parseTime=true&loc=Local",
	})
	ok, c, err := s.GetComment(context.Background(), 1)
	t.Logf("ok:%v, c:%+v, err:%v", ok, c, err)
}

// go test -v *.go -test.run=TestConcurrentAddComment
func TestConcurrentAddComment(t *testing.T) {
	s := NewStore().NewDB(config.DBConfig{
		Driver: "mysql",
		DSN:    "root:123456@tcp(127.0.0.1:3306)/dayu?parseTime=true&loc=Local",
	})
	for i := 0; i < 300; i++ {
		go func(i int) {
			err := s.addCommentSubject(i, i*2, 2)
			t.Logf("i:%d, err:%v", i, err)
		}(i)
	}
	time.Sleep(20 * time.Second)
}

func (s *Store) addCommentSubject(flag, objId, objType int) error {
	return BeginTX(s.Db, func(t *sqlx.Tx) error {
		subjectId := 0
		query := `SELECT id FROM comment_subject_1 WHERE obj_id = ? AND obj_type = ?`
		if err := t.Get(&subjectId, query, objId, objType); err != nil {
			if err != sql.ErrNoRows {
				panic(err)
			}
		}
		if subjectId > 0 {
			log.Printf("comment subject has existed, id:%d, obj_id:%d, obj_type:%d", subjectId, objId, objType)
			return nil
		}
		sql := `INSERT INTO comment_subject_1(obj_id, obj_type) VALUES (?,?);`
		res, err := t.Exec(sql, objId, objType)
		if err != nil {
			return err
		}
		newId, err := res.LastInsertId()
		log.Printf("new id: %d, err:%v", newId, err)
		return nil
	})
}
