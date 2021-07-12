package store

import (
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/dayu-go/comment/app/comment/service/internal/config"
	"github.com/jmoiron/sqlx"
)

var (
	DateFormat     = "2006-01-02"
	DateTimeFormat = "2006-01-02 15:04:05"
)

type Store struct {
	Db *sqlx.DB
}

func NewStore() *Store {
	return &Store{}
}

func (s *Store) NewDB(c config.DBConfig) *Store {
	db, err := sqlx.Connect(c.Driver, c.DSN)
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
		return nil
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("failed ping connection to mysql: %v", err)
		return nil
	}
	s.Db = db
	return s
}

func BeginTX(db *sqlx.DB, f func(t *sqlx.Tx) error) error {
	t, err := db.Beginx()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			_ = t.Rollback()
		}
	}()
	if err = f(t); err != nil {
		_ = t.Rollback()
		return err
	}
	return t.Commit()
}
