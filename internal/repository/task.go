package repository

import (
	"database/sql"
	"fmt"

	"github.com/riyan-eng/boilerplate3/internal/serrepconnector"
	"github.com/riyan-eng/boilerplate3/pkg/util"
)

type TaskQuery interface {
	ListTask()
	CreateTask(serrepconnector.CreateTaskReq)
	DeleteTask()
	DetailTask()
	UpdateTask()
}

type taskQuery struct {
	db *sql.DB
}

func (u *taskQuery) ListTask() {
}

func (u *taskQuery) CreateTask(req serrepconnector.CreateTaskReq) {
	query := fmt.Sprintf(`
		insert into public.tasks (id, name, detail) values ('%v', '%v', '%v')
	`, req.ID, req.Name, req.Detail)
	_, err := u.db.Exec(query)
	util.PanicIfNeeded(err)
}

func (u *taskQuery) DeleteTask() {
}

func (u *taskQuery) DetailTask() {
}

func (u *taskQuery) UpdateTask() {
}
