package repository

import (
	"database/sql"
	"fmt"

	"github.com/riyan-eng/boilerplate3/internal/serrepconnector"
	"github.com/riyan-eng/boilerplate3/pkg/util"
)

type TaskQuery interface {
	ListTask(serrepconnector.TaskListReq) *sql.Rows
	CreateTask(serrepconnector.TaskCreateReq)
	DeleteTask(serrepconnector.TaskDeleteReq)
	DetailTask(serrepconnector.TaskDetailReq) *sql.Rows
	UpdateTask(serrepconnector.TaskUpdateReq)
}

type taskQuery struct {
	db *sql.DB
}

func (u *taskQuery) ListTask(req serrepconnector.TaskListReq) *sql.Rows {
	query := fmt.Sprintf(`select id, name, detail, count(*) over() as total from public.tasks 
		where lower(name) like lower('%%%v%%') order by created_at %v limit %v offset %v`,
		req.Search, req.Order, req.Limit, req.Offset)
	rows, err := u.db.Query(query)
	util.PanicIfNeeded(err)
	return rows
}

func (u *taskQuery) CreateTask(req serrepconnector.TaskCreateReq) {
	query := fmt.Sprintf(`
		insert into public.tasks (id, name, detail) values ('%v', '%v', '%v')
	`, req.ID, req.Name, req.Detail)
	_, err := u.db.Exec(query)
	util.PanicIfNeeded(err)
}

func (u *taskQuery) DeleteTask(req serrepconnector.TaskDeleteReq) {
	query := fmt.Sprintf(`
		delete from public.tasks where id = '%v'
	`, req.ID)
	_, err := u.db.Exec(query)
	util.PanicIfNeeded(err)
}

func (u *taskQuery) DetailTask(req serrepconnector.TaskDetailReq) *sql.Rows {
	query := fmt.Sprintf(`
		select id, name, detail from public.tasks where id = '%v'
	`, req.ID)
	rows, err := u.db.Query(query)
	util.PanicIfNeeded(err)
	return rows
}

func (u *taskQuery) UpdateTask(req serrepconnector.TaskUpdateReq) {
	query := fmt.Sprintf(`
		update public.tasks set name = '%v', detail = '%v', updated_at = current_timestamp where id = '%v' 
	`, req.Name, req.Detail, req.ID)
	_, err := u.db.Exec(query)
	util.PanicIfNeeded(err)
}
