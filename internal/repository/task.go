package repository

type TaskQuery interface {
	ListTask()
	CreateTask()
	DeleteTask()
	DetailTask()
	UpdateTask()
}

type taskQuery struct{}

func (u *taskQuery) ListTask() {
}

func (u *taskQuery) CreateTask() {
}

func (u *taskQuery) DeleteTask() {
}

func (u *taskQuery) DetailTask() {
}

func (u *taskQuery) UpdateTask() {
}
