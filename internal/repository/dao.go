package repository

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type DAO interface {
	NewTaskQuery() TaskQuery
}

type dao struct{}

var DB *sql.DB

func NewDAO(db *sql.DB) DAO {
	DB = db
	return &dao{}
}

func NewDB() (*sql.DB, error) {
	dsn := "host=localhost user=postgres password=riyan dbname=boilerplate3_development port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return DB, nil
}

func (d *dao) NewTaskQuery() TaskQuery {
	return &taskQuery{}
}
