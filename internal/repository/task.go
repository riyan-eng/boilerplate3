package repository

import (
	"database/sql"
	"fmt"

	"github.com/riyan-eng/boilerplate3/internal/serrepconnector"
	"github.com/riyan-eng/boilerplate3/pkg/util"
)

type TaskQuery interface {
	ListTask(serrepconnector.ListTaskReq) *sql.Rows
	CreateTask(serrepconnector.CreateTaskReq)
	DeleteTask()
	DetailTask()
	UpdateTask()
}

type taskQuery struct {
	db *sql.DB
}

func (u *taskQuery) ListTask(r serrepconnector.ListTaskReq) *sql.Rows {
	query := fmt.Sprintf(`select id, name, detail from public.tasks 
		where lower(name) like '%%%v%%' order by created_at %v limit %v offset %v`,
		r.Search, r.Order, r.Limit, r.Offset)
	rows, err := u.db.Query(query)
	util.PanicIfNeeded(err)
	return rows
}

func (u *taskQuery) CreateTask(r serrepconnector.CreateTaskReq) {
	query := fmt.Sprintf(`
		insert into public.tasks (id, name, detail) values ('%v', '%v', '%v')
	`, r.ID, r.Name, r.Detail)
	_, err := u.db.Exec(query)
	util.PanicIfNeeded(err)
}

func (u *taskQuery) DeleteTask() {
}

func (u *taskQuery) DetailTask() {
}

func (u *taskQuery) UpdateTask() {
}
