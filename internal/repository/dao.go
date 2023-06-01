package repository

import "database/sql"

type DAO interface {
	NewTaskQuery() TaskQuery
}

type dao struct{}

var DB *sql.DB

func NewDAO() DAO {
	return &dao{}
}

func NewDB() (*sql.DB, error) {
	dsn := ""
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
